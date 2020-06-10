// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
)

var errTraceLimitReached = errors.New("the number of logs reached the specified limit")

// Storage represents a contract's storage.
type Storage map[common.Hash]common.Hash

// Copy duplicates the current storage.
func (s Storage) Copy() Storage {
	cpy := make(Storage)
	for key, value := range s {
		cpy[key] = value
	}
	return cpy
}

// LogConfig are the configuration options for structured logger the EVM
type LogConfig struct {
	DisableMemory  bool // disable memory capture
	DisableStack   bool // disable stack capture
	DisableStorage bool // disable storage capture

	// These options only work when `DisableStorage` is false
	EnableStorageRead    bool // enable capture storage read for all steps
	EnableStorageWritten bool // enable capture storage written for all steps
	EnableStorageView    bool // enable capture storage view(visited partial storage) for all steps

	Debug bool // print output during capture end
	Limit int  // maximum length of output, but zero means unlimited
}

//go:generate gencodec -type StructLog -field-override structLogMarshaling -out gen_structlog.go

// StructLog is emitted to the EVM each cycle and lists information about the current internal state
// prior to the execution of the statement.
type StructLog struct {
	Pc          uint64     `json:"pc"`
	Op          OpCode     `json:"op"`
	Gas         uint64     `json:"gas"`
	GasCost     uint64     `json:"gasCost"`
	Memory      []byte     `json:"memory"`
	MemorySize  int        `json:"memSize"`
	Stack       []*big.Int `json:"stack"`
	ReturnStack []uint64   `json:"returnStack"`

	// Deprecated but keep it here for backward compatibility.
	// Use `StorageRead`, `StorageWritten` or `StorageView` instead.
	Storage        map[common.Hash]common.Hash `json:"-"`
	StorageRead    map[common.Hash]common.Hash `json:"-"`
	StorageWritten map[common.Hash]common.Hash `json:"-"`
	StorageView    map[common.Hash]common.Hash `json:"-"`

	Depth         int    `json:"depth"`
	RefundCounter uint64 `json:"refund"`
	Err           error  `json:"-"`
}

// overrides for gencodec
type structLogMarshaling struct {
	Stack       []*math.HexOrDecimal256
	ReturnStack []math.HexOrDecimal64
	Gas         math.HexOrDecimal64
	GasCost     math.HexOrDecimal64
	Memory      hexutil.Bytes
	OpName      string `json:"opName"` // adds call to OpName() in MarshalJSON
	ErrorString string `json:"error"`  // adds call to ErrorString() in MarshalJSON
}

// OpName formats the operand name in a human-readable format.
func (s *StructLog) OpName() string {
	return s.Op.String()
}

// ErrorString formats the log's error as a string.
func (s *StructLog) ErrorString() string {
	if s.Err != nil {
		return s.Err.Error()
	}
	return ""
}

// Tracer is used to collect execution traces from an EVM transaction
// execution. CaptureState is called for each step of the VM with the
// current VM state.
// Note that reference types are actual VM data structures; make copies
// if you need to retain them beyond the current call.
type Tracer interface {
	CaptureStart(from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int) error
	CaptureState(env *EVM, pc uint64, op OpCode, gas, cost uint64, memory *Memory, stack *Stack, rStack *ReturnStack, contract *Contract, depth int, err error) error
	CaptureFault(env *EVM, pc uint64, op OpCode, gas, cost uint64, memory *Memory, stack *Stack, rStack *ReturnStack, contract *Contract, depth int, err error) error
	CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error
}

// StructLogger is an EVM state logger and implements Tracer.
//
// StructLogger can capture state based on the given Log configuration and also keeps
// a track record of modified storage which is used in reporting snapshots of the
// contract their storage.
type StructLogger struct {
	cfg LogConfig

	logs   []StructLog
	output []byte
	err    error

	// storageRead and storageWritten is all the storage read from/written to
	// so far during the current call. The value might be overwritten by the
	// following instructions.
	storageRead    map[common.Address]Storage
	storageWritten map[common.Address]Storage

	// storageView is the partial storage which has been visited during the
	// current call. It represents the current status of the storage.
	storageView map[common.Address]Storage
}

// NewStructLogger returns a new logger
func NewStructLogger(cfg *LogConfig) *StructLogger {
	logger := &StructLogger{
		storageRead:    make(map[common.Address]Storage),
		storageWritten: make(map[common.Address]Storage),
		storageView:    make(map[common.Address]Storage),
	}
	if cfg != nil {
		logger.cfg = *cfg
	}
	return logger
}

// CaptureStart implements the Tracer interface to initialize the tracing operation.
func (l *StructLogger) CaptureStart(from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int) error {
	return nil
}

// CaptureState logs a new structured log message and pushes it out to the environment
//
// CaptureState also tracks SLOAD/SSTORE ops to track storage change.
func (l *StructLogger) CaptureState(env *EVM, pc uint64, op OpCode, gas, cost uint64, memory *Memory, stack *Stack, rStack *ReturnStack, contract *Contract, depth int, err error) error {
	// check if already accumulated the specified number of logs
	if l.cfg.Limit != 0 && l.cfg.Limit <= len(l.logs) {
		return errTraceLimitReached
	}
	// initialise new changed values storage container for this contract
	// if not present.
	if l.storageRead[contract.Address()] == nil {
		l.storageRead[contract.Address()] = make(Storage)
	}
	if l.storageWritten[contract.Address()] == nil {
		l.storageWritten[contract.Address()] = make(Storage)
	}
	if l.storageView[contract.Address()] == nil {
		l.storageView[contract.Address()] = make(Storage)
	}
	// capture SLOAD opcodes and record the read entry in the local storage
	if op == SLOAD && stack.len() >= 1 {
		var (
			address = common.Hash(stack.data[stack.len()-1].Bytes32())
			value   = env.StateDB.GetState(contract.Address(), address)
		)
		l.storageRead[contract.Address()][address] = value
		l.storageView[contract.Address()][address] = value
	}
	// capture SSTORE opcodes and record the written entry in the local storage.
	if op == SSTORE && stack.len() >= 2 {
		var (
			value   = common.Hash(stack.data[stack.len()-2].Bytes32())
			address = common.Hash(stack.data[stack.len()-1].Bytes32())
		)
		l.storageWritten[contract.Address()][address] = value
		l.storageView[contract.Address()][address] = value
	}
	// Copy a snapshot of the current memory state to a new buffer
	var mem []byte
	if !l.cfg.DisableMemory {
		mem = make([]byte, len(memory.Data()))
		copy(mem, memory.Data())
	}
	// Copy a snapshot of the current stack state to a new buffer
	var stck []*big.Int
	if !l.cfg.DisableStack {
		stck = make([]*big.Int, len(stack.Data()))
		for i, item := range stack.Data() {
			stck[i] = new(big.Int).Set(item.ToBig())
		}
	}
	// Copy a snapshot of the current storage to a new container
	var (
		storage        Storage // legacy storage which contains all writes
		storageRead    Storage // storage container for all reads
		storageWritten Storage // storage container for all writes
		storageView    Storage // storage container for all reads and writes
	)
	if !l.cfg.DisableStorage {
		storage = l.storageWritten[contract.Address()].Copy()
		if l.cfg.EnableStorageRead {
			storageRead = l.storageRead[contract.Address()].Copy()
		}
		if l.cfg.EnableStorageWritten {
			storageWritten = l.storageWritten[contract.Address()].Copy()
		}
		if l.cfg.EnableStorageView {
			storageView = l.storageView[contract.Address()].Copy()
		}
	}
	var rstack []uint64
	if !l.cfg.DisableStack && rStack != nil {
		rstck := make([]uint64, len(rStack.data))
		copy(rstck, rStack.data)
	}
	// create a new snapshot of the EVM.
	log := StructLog{pc, op, gas, cost, mem, memory.Len(), stck, rstack, storage /* backward compatibility*/, storageRead, storageWritten, storageView, depth, env.StateDB.GetRefund(), err}
	l.logs = append(l.logs, log)
	return nil
}

// CaptureFault implements the Tracer interface to trace an execution fault
// while running an opcode.
func (l *StructLogger) CaptureFault(env *EVM, pc uint64, op OpCode, gas, cost uint64, memory *Memory, stack *Stack, rStack *ReturnStack, contract *Contract, depth int, err error) error {
	return nil
}

// CaptureEnd is called after the call finishes to finalize the tracing.
func (l *StructLogger) CaptureEnd(output []byte, gasUsed uint64, t time.Duration, err error) error {
	l.output = output
	l.err = err
	if l.cfg.Debug {
		fmt.Printf("0x%x\n", output)
		if err != nil {
			fmt.Printf(" error: %v\n", err)
		}
	}
	return nil
}

// StructLogs returns the captured log entries.
func (l *StructLogger) StructLogs() []StructLog { return l.logs }

// Error returns the VM error captured by the trace.
func (l *StructLogger) Error() error { return l.err }

// Output returns the VM return value captured by the trace.
func (l *StructLogger) Output() []byte { return l.output }

// WriteTrace writes a formatted trace to the given writer
func WriteTrace(writer io.Writer, logs []StructLog) {
	for _, log := range logs {
		fmt.Fprintf(writer, "%-16spc=%08d gas=%v cost=%v", log.Op, log.Pc, log.Gas, log.GasCost)
		if log.Err != nil {
			fmt.Fprintf(writer, " ERROR: %v", log.Err)
		}
		fmt.Fprintln(writer)

		if len(log.Stack) > 0 {
			fmt.Fprintln(writer, "Stack:")
			for i := len(log.Stack) - 1; i >= 0; i-- {
				fmt.Fprintf(writer, "%08d  %x\n", len(log.Stack)-i-1, math.PaddedBigBytes(log.Stack[i], 32))
			}
		}
		if len(log.ReturnStack) > 0 {
			fmt.Fprintln(writer, "ReturnStack:")
			for i := len(log.Stack) - 1; i >= 0; i-- {
				fmt.Fprintf(writer, "%08d  0x%x (%d)\n", len(log.Stack)-i-1, log.ReturnStack[i], log.ReturnStack[i])
			}
		}
		if len(log.Memory) > 0 {
			fmt.Fprintln(writer, "Memory:")
			fmt.Fprint(writer, hex.Dump(log.Memory))
		}
		// For backward compatibility, still print storage here.
		if len(log.Storage) > 0 {
			fmt.Fprintln(writer, "Storage:")
			for h, item := range log.Storage {
				fmt.Fprintf(writer, "%x: %x\n", h, item)
			}
		}
		if len(log.StorageRead) > 0 {
			fmt.Fprintln(writer, "StorageRead:")
			for h, item := range log.StorageRead {
				fmt.Fprintf(writer, "%x: %x\n", h, item)
			}
		}
		if len(log.StorageWritten) > 0 {
			fmt.Fprintln(writer, "StorageWritten:")
			for h, item := range log.StorageWritten {
				fmt.Fprintf(writer, "%x: %x\n", h, item)
			}
		}
		fmt.Fprintln(writer)
	}
}

// WriteLogs writes vm logs in a readable format to the given writer
func WriteLogs(writer io.Writer, logs []*types.Log) {
	for _, log := range logs {
		fmt.Fprintf(writer, "LOG%d: %x bn=%d txi=%x\n", len(log.Topics), log.Address, log.BlockNumber, log.TxIndex)

		for i, topic := range log.Topics {
			fmt.Fprintf(writer, "%08d  %x\n", i, topic)
		}

		fmt.Fprint(writer, hex.Dump(log.Data))
		fmt.Fprintln(writer)
	}
}

type mdLogger struct {
	out io.Writer
	cfg *LogConfig
}

// NewMarkdownLogger creates a logger which outputs information in a format adapted
// for human readability, and is also a valid markdown table
func NewMarkdownLogger(cfg *LogConfig, writer io.Writer) *mdLogger {
	l := &mdLogger{writer, cfg}
	if l.cfg == nil {
		l.cfg = &LogConfig{}
	}
	return l
}

func (t *mdLogger) CaptureStart(from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int) error {
	if !create {
		fmt.Fprintf(t.out, "From: `%v`\nTo: `%v`\nData: `0x%x`\nGas: `%d`\nValue `%v` wei\n",
			from.String(), to.String(),
			input, gas, value)
	} else {
		fmt.Fprintf(t.out, "From: `%v`\nCreate at: `%v`\nData: `0x%x`\nGas: `%d`\nValue `%v` wei\n",
			from.String(), to.String(),
			input, gas, value)
	}

	fmt.Fprintf(t.out, `
|  Pc   |      Op     | Cost |   Stack   |   RStack  |
|-------|-------------|------|-----------|-----------|
`)
	return nil
}

func (t *mdLogger) CaptureState(env *EVM, pc uint64, op OpCode, gas, cost uint64, memory *Memory, stack *Stack, rStack *ReturnStack, contract *Contract, depth int, err error) error {
	fmt.Fprintf(t.out, "| %4d  | %10v  |  %3d |", pc, op, cost)

	if !t.cfg.DisableStack { // format stack
		var a []string
		for _, elem := range stack.data {
			a = append(a, fmt.Sprintf("%d", elem))
		}
		b := fmt.Sprintf("[%v]", strings.Join(a, ","))
		fmt.Fprintf(t.out, "%10v |", b)
	}
	if !t.cfg.DisableStack { // format return stack
		var a []string
		for _, elem := range rStack.data {
			a = append(a, fmt.Sprintf("%2d", elem))
		}
		b := fmt.Sprintf("[%v]", strings.Join(a, ","))
		fmt.Fprintf(t.out, "%10v |", b)
	}
	fmt.Fprintln(t.out, "")
	if err != nil {
		fmt.Fprintf(t.out, "Error: %v\n", err)
	}
	return nil
}

func (t *mdLogger) CaptureFault(env *EVM, pc uint64, op OpCode, gas, cost uint64, memory *Memory, stack *Stack, rStack *ReturnStack, contract *Contract, depth int, err error) error {

	fmt.Fprintf(t.out, "\nError: at pc=%d, op=%v: %v\n", pc, op, err)

	return nil
}

func (t *mdLogger) CaptureEnd(output []byte, gasUsed uint64, tm time.Duration, err error) error {
	fmt.Fprintf(t.out, `
Output: 0x%x
Consumed gas: %d
Error: %v
`,
		output, gasUsed, err)
	return nil
}
