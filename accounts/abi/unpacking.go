package abi

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

// toGoSliceType parses the input and casts it to the proper slice defined by the ABI
// argument in T.
func toGoSlice(i int, t Argument, output []byte) (interface{}, error) {
	index := i * 32
	// The slice must, at very least be large enough for the index+32 which is exactly the size required
	// for the [offset in output, size of offset].
	if index+32 > len(output) {
		return nil, fmt.Errorf("abi: cannot marshal in to go slice: insufficient size output %d require %d", len(output), index+32)
	}
	elem := t.Type.Elem

	// first we need to create a slice of the type
	var refSlice reflect.Value
	switch elem.T {
	case IntTy, UintTy, BoolTy:
		// create a new reference slice matching the element type
		switch t.Type.Kind {
		case reflect.Bool:
			refSlice = reflect.ValueOf([]bool(nil))
		case reflect.Uint8:
			refSlice = reflect.ValueOf([]uint8(nil))
		case reflect.Uint16:
			refSlice = reflect.ValueOf([]uint16(nil))
		case reflect.Uint32:
			refSlice = reflect.ValueOf([]uint32(nil))
		case reflect.Uint64:
			refSlice = reflect.ValueOf([]uint64(nil))
		case reflect.Int8:
			refSlice = reflect.ValueOf([]int8(nil))
		case reflect.Int16:
			refSlice = reflect.ValueOf([]int16(nil))
		case reflect.Int32:
			refSlice = reflect.ValueOf([]int32(nil))
		case reflect.Int64:
			refSlice = reflect.ValueOf([]int64(nil))
		default:
			refSlice = reflect.ValueOf([]*big.Int(nil))
		}
	case AddressTy: // address must be of slice Address
		refSlice = reflect.ValueOf([]common.Address(nil))
	case HashTy: // hash must be of slice hash
		refSlice = reflect.ValueOf([]common.Hash(nil))
	case FixedBytesTy:
		refSlice = reflect.ValueOf([][]byte(nil))
	default: // no other types are supported
		return nil, fmt.Errorf("abi: unsupported slice type %v", elem.T)
	}

	var slice []byte
	var size int
	var offset int
	if t.Type.IsSlice {
		// get the offset which determines the start of this array ...
		offset = int(binary.BigEndian.Uint64(output[index+24 : index+32]))
		if offset+32 > len(output) {
			return nil, fmt.Errorf("abi: cannot marshal in to go slice: offset %d would go over slice boundary (len=%d)", len(output), offset+32)
		}

		slice = output[offset:]
		// ... starting with the size of the array in elements ...
		size = int(binary.BigEndian.Uint64(slice[24:32]))
		slice = slice[32:]
		// ... and make sure that we've at the very least the amount of bytes
		// available in the buffer.
		if size*32 > len(slice) {
			return nil, fmt.Errorf("abi: cannot marshal in to go slice: insufficient size output %d require %d", len(output), offset+32+size*32)
		}

		// reslice to match the required size
		slice = slice[:size*32]
	} else if t.Type.IsArray {
		//get the number of elements in the array
		size = t.Type.SliceSize

		//check to make sure array size matches up
		if index+32*size > len(output) {
			return nil, fmt.Errorf("abi: cannot marshal in to go array: offset %d would go over slice boundary (len=%d)", len(output), index+32*size)
		}
		//slice is there for a fixed amount of times
		slice = output[index : index+size*32]
	}

	for i := 0; i < size; i++ {
		var (
			inter        interface{}             // interface type
			returnOutput = slice[i*32 : i*32+32] // the return output
		)
		// set inter to the correct type (cast)
		switch elem.T {
		case IntTy, UintTy:
			inter = readInteger(t.Type.Kind, returnOutput)
		case BoolTy:
			inter = !allZero(returnOutput)
		case AddressTy:
			inter = common.BytesToAddress(returnOutput)
		case HashTy:
			inter = common.BytesToHash(returnOutput)
		case FixedBytesTy:
			inter = returnOutput
		}
		// append the item to our reflect slice
		refSlice = reflect.Append(refSlice, reflect.ValueOf(inter))
	}

	// return the interface
	return refSlice.Interface(), nil
}

func readInteger(kind reflect.Kind, b []byte) interface{} {
	switch kind {
	case reflect.Uint8:
		return uint8(b[len(b)-1])
	case reflect.Uint16:
		return binary.BigEndian.Uint16(b[len(b)-2:])
	case reflect.Uint32:
		return binary.BigEndian.Uint32(b[len(b)-4:])
	case reflect.Uint64:
		return binary.BigEndian.Uint64(b[len(b)-8:])
	case reflect.Int8:
		return int8(b[len(b)-1])
	case reflect.Int16:
		return int16(binary.BigEndian.Uint16(b[len(b)-2:]))
	case reflect.Int32:
		return int32(binary.BigEndian.Uint32(b[len(b)-4:]))
	case reflect.Int64:
		return int64(binary.BigEndian.Uint64(b[len(b)-8:]))
	default:
		return new(big.Int).SetBytes(b)
	}
}

// todo: this is inefficient for a bool, just look at the last cell, save yourself 32 iterations
func allZero(b []byte) bool {
	for _, byte := range b {
		if byte != 0 {
			return false
		}
	}
	return true
}

// toGoType parses the input and casts it to the proper type defined by the ABI
// argument in T.
func toGoType(i int, t Argument, output []byte) (interface{}, error) {
	// we need to treat slices differently
	if (t.Type.IsSlice || t.Type.IsArray) && t.Type.T != BytesTy && t.Type.T != StringTy && t.Type.T != FixedBytesTy && t.Type.T != FunctionTy {
		return toGoSlice(i, t, output)
	}

	index := i * 32
	if index+32 > len(output) {
		return nil, fmt.Errorf("abi: cannot marshal in to go type: length insufficient %d require %d", len(output), index+32)
	}

	// Parse the given index output and check whether we need to read
	// a different offset and length based on the type (i.e. string, bytes)
	var returnOutput []byte
	switch t.Type.T {
	case StringTy, BytesTy: // variable arrays are written at the end of the return bytes
		// parse offset from which we should start reading
		offset := int(binary.BigEndian.Uint64(output[index+24 : index+32]))
		if offset+32 > len(output) {
			return nil, fmt.Errorf("abi: cannot marshal in to go type: length insufficient %d require %d", len(output), offset+32)
		}
		// parse the size up until we should be reading
		size := int(binary.BigEndian.Uint64(output[offset+24 : offset+32]))
		if offset+32+size > len(output) {
			return nil, fmt.Errorf("abi: cannot marshal in to go type: length insufficient %d require %d", len(output), offset+32+size)
		}

		// get the bytes for this return value
		returnOutput = output[offset+32 : offset+32+size]
	default:
		returnOutput = output[index : index+32]
	}

	// convert the bytes to whatever is specified by the ABI.
	switch t.Type.T {
	case IntTy, UintTy:
		return readInteger(t.Type.Kind, returnOutput), nil
	case BoolTy:
		return !allZero(returnOutput), nil
	case AddressTy:
		return common.BytesToAddress(returnOutput), nil
	case HashTy:
		return common.BytesToHash(returnOutput), nil
	case BytesTy, FixedBytesTy, FunctionTy:
		return returnOutput, nil
	case StringTy:
		return string(returnOutput), nil
	}
	return nil, fmt.Errorf("abi: unknown type %v", t.Type.T)
}
