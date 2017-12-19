// Code generated by go-bindata.
// sources:
// 4byte_tracer.js
// call_tracer.js
// emvdis_tracer.js
// noop_tracer.js
// opcount_tracer.js
// prestate_tracer.js
// DO NOT EDIT!

package tracers

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __4byte_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\xdd\x6e\xdb\x38\x13\xbd\x96\x9e\x62\x3e\xdf\xc4\x4a\x64\xc9\xf2\x4f\xfc\x13\xa4\xf8\xb2\x6d\xd0\x16\x68\xd1\x05\x12\xec\xcd\x62\x2f\x68\x72\x64\x71\x2d\x93\x02\x39\x72\xec\xa4\x7e\xf7\x05\x29\xb9\xb6\x93\x60\x81\xd5\x85\x2d\x91\x33\x87\x47\x67\xe6\x8c\xd2\xcb\xcb\x30\x84\xd1\x62\x47\x68\xc1\x22\x33\xbc\x40\x0b\xb9\x36\xcd\x5a\x4f\x0a\x54\x24\x73\x89\xc6\xc6\xc0\x94\x00\xae\xcb\x12\x39\x59\xa0\x02\xd7\x3e\xb0\xd2\x96\x7a\x95\xd1\x1c\xad\x95\x6a\x99\x84\xf0\x95\xce\xc2\x60\x8d\x54\x68\x61\xe1\x04\x0c\x58\xa9\xd5\x12\x9e\x24\x15\x3e\xc4\xca\x67\x04\x9d\x37\xf7\x75\x55\x95\x12\x05\x08\x46\x2c\x06\xab\x81\x0a\x46\x21\x30\x30\xb8\x41\x63\x51\x80\x95\x4b\xc5\xa8\x36\x08\x9c\x29\x58\x20\xac\x19\xf1\x02\x05\xb0\x25\x93\xca\xd2\x1b\x4c\x07\x95\x84\xe1\xfd\x96\xad\xab\x12\xe7\x61\x08\x00\xf0\x01\x04\x2e\xea\x65\x42\x86\x71\x7c\x34\x4c\x59\xc6\x49\x6a\xd5\x85\x4e\x7f\x3b\xc8\x46\x38\x9e\x4d\x70\x38\x16\xac\x3f\x1d\x5e\xcf\x06\xf9\x78\x38\xbd\xce\x46\x19\x5e\xcf\xf2\xd1\x04\x67\x93\xe1\x62\xc0\xc7\xd7\x38\x61\xd3\xfe\x64\xb8\xc8\x90\xf5\xa7\xb9\x98\x8c\x27\x19\xce\x04\x76\x62\x78\xf1\xc0\x66\x0e\x9d\x46\xe0\xce\x3e\xf2\xe7\xbe\xf8\x5f\x80\xfe\x76\x30\x11\x7c\x30\x9b\x60\x2f\x1b\x4c\xe7\x90\xc5\xbf\x36\x86\x53\xce\x47\xd3\x61\xd6\xeb\xcf\x61\x70\x5c\x1e\x0f\x46\xf9\x70\x3a\x9d\xf5\x66\xd7\x67\xe1\x4c\xe4\xe3\x59\x3e\x9b\xf5\x06\xd3\x73\x1c\x3e\x98\x66\x22\x9b\xa1\xc3\xc9\xfc\xea\x3e\x0c\xff\xcf\x6a\x2a\xb4\x81\xef\x5f\xe0\xe1\x09\x95\xc0\x10\x2e\x2f\xd3\xf0\x25\x0c\xa4\xb0\x30\x87\x97\x7d\x1c\x86\x41\x9a\x02\x67\x65\xf9\xb8\xab\x10\x0c\x52\x6d\x94\x85\x8b\x9c\x95\x16\x2f\x7c\xdd\x95\x56\x3d\x17\x60\x63\xd0\xc6\x89\xec\x53\x2a\xc4\x55\x4f\x2a\x81\x5b\x1f\xe4\xb4\xcf\xa5\xb1\x04\x15\x33\x6c\x0d\x2c\x27\x34\x70\xb1\x61\x65\x8d\x17\x31\xc8\x04\x13\x58\xe3\xda\xd5\x8c\x19\xf2\x04\x7f\x1d\x3a\x87\xbc\x56\x4d\x41\x74\x65\xc9\x44\x07\xdd\x00\xec\x93\x24\x5e\xbc\x59\x6e\xd2\x2d\x42\xe7\xe3\xdd\xb7\x6f\x9d\x79\x18\x04\xc1\xf1\xf9\xe3\x8f\x4f\xf7\x9d\xf9\x59\x70\x90\xa6\xcb\x98\xc5\x9b\xb8\x21\xd1\xfe\xd9\x67\x7f\xa3\x6b\x3a\xfc\xdb\xe7\xf3\xb4\x46\x0f\x18\x42\x9a\x5a\x62\x7c\x05\x15\x19\x20\xdd\xa4\xbf\x43\xe7\xd3\xfd\xb7\xfb\xcf\x77\x8f\xf7\x2d\xad\xb7\x01\x0f\x8f\x77\x8f\x5f\x3f\xbe\xb3\xed\xae\x86\xe4\x7f\x62\xe8\xae\x96\xe4\xe0\xe6\x5f\x49\xee\xc3\x57\xf1\xbe\xc4\x37\x4d\xaf\xb4\x7d\x60\x49\x1b\x04\xcb\x36\xe8\x0b\xba\x94\x1b\x54\xe0\x8a\xdc\xda\xd9\x4f\x06\x67\x31\x67\xb9\x30\xf0\xe1\x27\xd5\x93\x22\xf6\x66\x8c\x5e\xc2\x20\xd8\x30\x03\x2b\xdc\xc1\x2d\x74\x3a\x57\x52\x5c\x75\x7a\x9d\x2b\xb7\x79\x13\x06\x01\x15\xd2\x26\x52\xd8\x3f\x57\xb8\xfb\x0b\x6e\xe1\xec\xf9\x2a\x83\x9f\x3f\x21\xbb\x09\x83\x23\x2d\xac\x40\x5a\x90\x6a\xa3\x57\x28\x7c\xc3\xb9\x09\xb1\x03\x5d\x71\x2d\xd0\xcf\x0d\xcf\xf8\x8f\xef\x80\x5b\xe4\x35\xa1\x4d\x1c\x3f\xac\x4e\xe8\x95\x7a\x19\x83\x58\x44\x70\x6c\x24\x47\x92\xd3\x81\xc1\xa1\x23\x5d\x64\xa2\xab\x84\xf4\x03\x19\xa9\x96\xdd\x28\xba\xf9\x95\x22\x73\xe8\xfe\x8f\x53\xd4\xca\x78\xdc\x48\x53\x78\x58\xc9\x0a\x98\xda\x41\x65\xb0\xc7\xf5\xba\x92\x25\x7a\xd6\x9c\x39\x06\x36\x06\x2a\xb4\x45\x60\x06\xe1\xef\xda\x12\xe4\x4c\xf1\xc3\x5b\xd8\xb3\x23\xa4\xfd\xdd\x60\x0b\x21\xba\xa4\xef\x84\x30\x68\xad\xa7\xe6\xab\x9c\x38\x07\x76\xb3\x28\xf9\xcd\x8d\x9c\x6e\x14\x45\xa7\x2f\x76\xac\xf3\x3b\x0d\x90\xa6\xf0\x99\x51\x81\x06\xa4\x22\x34\x8a\x95\xde\x8c\x20\x90\x98\x2c\xed\x99\x3a\x52\x3d\x3c\xc3\x2d\xbc\x3a\x96\x13\x5c\x41\x16\x25\x5f\x15\x5d\x8f\xba\xaf\xd4\xf1\x29\x1f\x6e\x61\xf4\x9a\x50\x83\xf7\x23\xcf\xdf\x03\x7c\x07\xcc\x5d\xbe\x30\xbe\xd1\xba\xa4\xbf\xe0\xd6\x0b\xe0\xac\x60\x76\x89\x2d\x25\xc7\xae\x47\x8c\x5b\xe0\x2b\x18\x45\x51\xec\x59\xf7\x46\x27\x50\xfb\x63\x37\x19\xb4\x75\x49\xa7\xfd\xf4\x54\xa0\x02\x27\x80\x6b\xa1\xb6\x18\x50\x38\x17\x2c\xd0\x19\x80\xd0\x30\x42\x01\x7a\xd3\x5a\xa0\x1d\x93\x1e\xae\x99\x7c\x4e\xc3\x16\xb8\xfd\x12\xb9\x6f\x82\xff\x4c\x06\xcd\xfa\x49\x23\x72\xda\x3a\x69\x82\x20\x4d\x1f\x0e\x5e\xd3\xb5\x1b\x98\xae\x0c\xce\x5f\xc0\x4a\xab\xc3\x20\x90\xb9\x0b\x4e\xa4\xaa\x6a\x4a\x4a\x54\x4b\x2a\xe0\x03\x8c\xbc\xc1\x82\x37\xd2\x1c\x43\x1b\x65\xfa\xb1\x13\xe3\x35\x40\x6f\x04\x51\x18\x04\xfb\x30\x38\x8c\xb7\x83\xfd\x1a\xcb\xed\xc3\x7f\x02\x00\x00\xff\xff\x1b\x64\x99\x95\x2d\x08\x00\x00")

func _4byte_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		__4byte_tracerJs,
		"4byte_tracer.js",
	)
}

func _4byte_tracerJs() (*asset, error) {
	bytes, err := _4byte_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4byte_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _call_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x51\x73\xdb\x38\xee\x7f\xb6\x3f\x05\xda\x87\x8d\x3d\x75\x6d\xb7\xdd\x7f\x1f\x9c\xbf\xf7\x26\x93\x7a\xf7\x3a\x93\xbb\xec\xb4\xd9\xdd\x87\x4e\x1e\x68\x09\xb2\xd9\x50\xa4\x96\xa4\xec\xf8\xba\xf9\xee\x37\x00\x29\x59\xb2\x95\xd4\xd7\x9b\xbb\xd9\x7b\xc9\x44\x22\x00\x81\xc0\x0f\x3f\x80\xf4\x64\x02\x89\x50\xea\xc6\x8a\x04\x2d\x48\x07\x02\xb2\x52\x29\x58\x2a\xb3\xd5\xe0\xad\xd0\x4e\x24\x5e\x1a\xfe\x9f\x44\xfc\x5a\x78\xc0\x7b\x7a\xf2\x0e\x84\x4e\xc1\x62\x61\x2c\xfd\xaf\x54\x7f\x32\x01\xbf\x46\x90\xda\xa3\xd5\x42\xb1\x6d\x07\xb9\x48\x11\x96\x3b\x10\x4d\x83\x23\x10\xca\xe8\x15\x6c\xa5\x5f\x83\xd0\x3b\x28\x1d\x66\xa5\x02\xa9\x33\x63\x73\x41\x22\xe3\xfe\x97\x7e\x2f\x7a\xe8\xbc\x48\xee\xc8\x41\xb2\x9f\x94\xd6\xa2\xf6\x60\x31\x29\xad\x93\x1b\x64\x11\x08\x32\x26\x63\x99\xc5\xaf\x7f\x03\xbc\xc7\xa4\x0c\x96\x7a\xb5\x91\x19\x7c\xfa\xf2\x70\x3b\xea\xb3\xe9\x14\x5d\x82\x3a\xc5\x94\xf7\x77\xe7\x60\xbb\x46\xbf\x46\x0b\x5b\x3c\xdb\x20\x7c\x2e\x9d\x6f\xc8\x64\xd6\xe4\x20\x34\x98\xd2\x53\x28\x1a\xd1\x91\xda\x1b\x36\x28\xe8\x7f\x8d\x96\x3d\x1a\xf7\x7b\xb5\xf2\x0c\x32\xa1\x1c\xc6\xef\x3a\x8f\x05\xed\x46\xea\x8d\xb9\x23\xcb\xc6\x02\x6e\xd0\xee\xc0\x14\x89\x49\x31\xc4\x99\xf6\x51\x6f\x03\xdd\xb8\xdf\x23\xbd\x19\x64\xa5\xe6\xcf\x0e\x94\x59\x8d\x20\x5d\x0e\xe1\x4b\xbf\x47\x66\x2f\x45\xe1\x4b\x8b\x1c\x4f\xb4\xd6\x58\x07\x32\xcf\x31\x95\xc2\xa3\xda\xf5\x7b\xbd\x8d\xb0\x61\x01\xe6\xa0\xcc\x6a\xbc\x42\xbf\xa0\xc7\xc1\xf0\xbc\xdf\xeb\xc9\x0c\x06\x61\xf5\xd9\x7c\x0e\xa5\x4e\x31\x93\x1a\xd3\x60\xbe\xe7\xd7\xd2\x8d\x33\x51\x2a\x5f\x7f\x97\x94\x7a\x16\x7d\x69\x35\xfd\xfb\x10\xbc\xf8\x0d\xc1\x68\xb5\x83\x44\x90\x2b\x4b\x53\x7a\x70\x3b\xe7\x31\x8f\x9b\x73\x23\xc8\x84\xa3\x10\xca\x0c\xb6\x08\x85\xc5\x97\xc9\x1a\x29\x77\x3a\xc1\xe8\xa5\xdb\x39\x4e\xea\x1c\xe8\x6b\x63\x53\x8c\xbd\xf9\x7b\x99\x2f\xd1\x0e\x86\xf0\x1d\x4c\xef\xb3\xe9\x10\xe6\x73\xfe\xa7\xf2\x3d\xea\x44\x7f\xc9\x8a\x29\xe2\x46\x59\xff\xa3\xb7\x52\xaf\xc2\x5e\xa3\xaf\xef\x33\x10\xa0\x71\x0b\x89\xd1\x0c\x6a\xca\xca\x12\xa5\x5e\x41\x62\x51\x78\x4c\x47\x20\xd2\x14\xbc\x09\xc8\xab\x71\xd6\xfe\x24\x7c\xf7\x1d\x7f\x6b\x0e\x67\x97\x1f\x16\x17\x37\x8b\xb3\x86\x13\x52\x5f\x67\x59\xf4\x83\x75\xc7\x05\xe2\xdd\xe0\xd5\x70\xbc\x11\xaa\xc4\xeb\x2c\x78\x14\x65\x17\x3a\x85\x79\xd4\x79\x71\xa8\xf3\xba\xa5\x43\x4a\x93\x09\x5c\x38\x87\xf9\x52\xe1\x71\xed\xc5\xe2\xe4\x3a\x75\xde\x58\x64\xa0\x25\x26\x2f\x14\x12\x80\xaa\xaf\xc6\x48\xb3\xc7\x3d\xbf\x2b\x70\x06\x00\x60\x8a\x11\xbf\x20\xd8\xf3\x0b\x6f\xfe\x8a\xf7\x9c\x8e\x2a\x5a\x04\xa0\x8b\x34\xb5\xe8\xdc\x60\x38\x0c\xe2\x52\x17\xa5\x9f\xb5\xc4\x73\xcc\x8d\xdd\x8d\x9d\x92\x09\x0e\x78\x6b\xa3\xb0\xd3\x4a\x67\x25\xdc\x7b\x4d\x3a\x11\x94\x3f\x09\x37\xd8\x2f\x5d\x1a\xe7\x67\xd5\x12\x3d\x54\x6b\x1c\x0b\x52\x3b\x9b\xde\x9f\x1d\x47\x6b\x3a\xdc\x27\xfd\xd5\xdb\x21\xa9\x3c\x9c\xd7\x50\xae\x19\x61\x5c\x94\x6e\x3d\x60\xe4\xec\x57\xf7\x55\x3f\x07\x6f\x4b\xec\x44\x3a\xa3\xe7\x18\x39\x0e\x55\x46\xb4\xe1\x6d\x99\x30\x82\x56\x82\x49\x85\x8b\x5a\x10\xc9\xba\x72\xc9\x31\xf7\xc6\x3c\x0a\xa4\x8f\x8b\xab\x1f\xdf\x2d\x3e\xde\x7c\xf8\xe5\xf2\xa6\x09\x27\x85\x99\x27\xa7\xda\x7b\x50\xa8\x57\x7e\xcd\xfe\x93\xb9\xf6\xea\x27\xd2\x79\xf9\xea\x36\xbc\x81\x79\x47\x75\xf7\x9e\xd6\x80\x4f\xb7\x6c\xfb\xe1\x38\x7c\x6d\xd1\x10\xcc\x2f\x01\x44\xa6\x78\x68\x72\x44\x47\xd9\xe5\xe8\xd7\x26\x65\x1e\x4c\x44\xa0\xd2\x2a\x8a\xa9\xd1\x78\x72\xf1\x0d\xaa\xea\xbb\xb8\xba\x3a\x83\x3f\xfe\x80\xc6\xf3\xe5\xf5\xbb\x45\xf3\xdd\xbb\xc5\xd5\xe2\xa7\x8b\x9b\xc5\xa1\xec\xc7\x9b\x8b\x9b\xf7\x97\xfc\x76\x18\xa3\x32\x99\xc0\xc7\x3b\x59\x30\xa1\x32\x4d\x99\xbc\x90\x0a\x1b\xfe\xba\x11\xf8\xb5\x71\x08\x44\x76\xdc\x2f\x32\xa1\x93\x8a\xc7\x5d\x95\x34\x6f\x28\x65\xa6\xaa\x95\x63\x2a\x68\x02\x75\x58\xa7\x51\xba\x9f\x2d\xc6\x8f\xa6\x03\x6f\x2a\xbf\xf6\x01\x0d\x19\x61\xae\x63\x92\x19\x9c\xbe\x49\xf8\x0b\x4c\x61\x06\xaf\x22\x93\x3c\x41\x55\xaf\xe1\x05\x99\xff\x06\xc2\x7a\xd3\xa1\xf9\xe7\xa4\x2d\x6f\x58\xb8\x12\xf7\xe6\xbf\x4f\x67\xa6\xf4\xd7\x59\x36\x83\xc3\x20\x7e\x7f\x14\xc4\x5a\xfe\x0a\xf5\xb1\xfc\xff\x1d\xc9\xef\xa9\x8f\x50\x65\x0a\x78\x76\x04\x91\x40\x3c\xcf\x0e\xea\x20\x06\x97\xa7\x19\xb6\x06\xf3\x47\xc8\xf6\x75\x1b\xc3\x8f\xb1\xc5\xbf\x45\xb6\x9d\x53\x19\xcd\x5e\xed\xb9\x6b\x04\x16\xbd\x95\xb8\x41\x90\xfe\xcc\xb1\x49\x9a\x4f\xcd\x56\xe8\x04\xc7\xf0\x1b\x06\x8b\x1a\x91\xc9\x25\xce\xb3\x34\x8e\xf0\x88\x47\x33\xa9\xd4\x7b\xce\x11\x3c\x76\x5a\x84\x5c\xec\x60\x89\x34\x7f\xdd\xed\x60\x25\x1c\xa4\x3b\x2d\x72\x99\xb8\x60\x8f\x67\x59\x8b\x2b\x61\xd9\xac\xc5\xdf\x4b\x74\x1e\x53\x06\xb2\x48\x7c\x29\x94\xda\xc1\x4a\x6e\x50\xb3\xf6\xe0\xf5\x9b\xe9\x14\x9c\x97\x05\xea\x74\x04\x6f\xdf\x4c\xde\x7e\x0f\xb6\x54\x38\x1c\xf7\x1b\x34\x5e\x6f\x35\x66\x83\x16\x22\x7a\xde\x61\xe1\xd7\x83\x21\xfc\xf0\x48\x3f\x78\x84\xdc\x3b\x65\xe1\x25\xbc\xba\x1d\x93\x5f\xf3\x16\x6e\x43\x26\x01\x95\xc3\x68\x6d\x32\x81\x9b\xeb\x77\xd7\x83\x3b\x61\x85\x12\x4b\x1c\xce\xe0\xa6\x8a\xd5\x56\xc4\x81\x9f\x92\x02\x85\x12\x52\x83\x48\x12\x53\x6a\x4f\x81\xaf\x66\x77\xb5\x23\x7e\x3f\xf3\x95\xbd\xb5\xd8\x20\xc9\xa1\x73\x15\xdd\x73\xd6\xc8\x1d\x91\x93\x36\x48\xed\x64\x8a\x8d\xac\x10\x3b\x18\xa6\xe6\x28\xb1\x95\x4a\x55\x06\x73\xe3\xe8\x23\x4b\x84\xad\xa5\x73\x86\x93\x3a\x21\x38\x40\x8a\x14\x6d\x07\x46\x83\x00\x65\x3c\x1d\x18\xb8\xc6\x41\xd8\x95\x1b\x07\xbe\xa7\xcf\x12\xe7\x68\xb3\x1d\xb7\x81\xdc\x84\x2a\x4f\xf4\x07\xe3\x80\x06\xbc\x97\xce\xf3\x00\x49\x5e\x4a\x07\x01\xc9\x52\xaf\x46\x50\x98\x82\x79\xfa\xc4\x59\xf2\xc3\xe2\xd7\xc5\x87\xba\xf9\x9f\x9e\xc4\x6a\xc4\x7f\x5e\x9f\x80\xc0\xd2\xf1\xc2\x63\xfa\xbc\x63\x66\xef\x00\xd4\xfc\x11\x40\x91\xfd\x7d\x6f\xfc\xb9\xb1\x1d\x25\x9c\xdf\x27\x66\x85\xe1\xf8\xd2\x74\xc0\x95\xca\xbb\x03\xee\x3e\x24\x07\x53\x54\x1d\x82\x9c\x62\xda\x21\x62\xef\x98\xac\x63\xc0\x7d\x13\x78\x02\x82\x4c\x83\x00\x78\xbd\x9a\xd0\x44\xe0\x7c\xf6\xd0\x94\x9e\x92\x4e\x5d\x7a\x4f\x71\x2b\xe1\x7e\x71\x9c\xdb\x48\x72\x4b\xb9\x7a\xaf\xfd\xa0\x5a\x7c\xaf\xe1\x25\x54\x0f\x44\xdd\xf0\xb2\x55\x2b\x1d\x1c\xd8\x4b\x51\xa1\x47\xd8\x9b\x38\x87\x83\x57\x64\x28\x6c\x9a\x43\x63\xd1\x1f\xb7\xe0\x69\xb4\x46\x61\x79\x66\xd1\x8f\xf1\xf7\x52\x28\x37\x98\xd6\x23\x41\xd8\x81\x37\xdc\xc4\xe6\x75\x1b\xab\xfa\x1c\xe9\xb4\x86\x8c\x68\x30\xa8\xc5\x68\x54\x6a\xe9\x32\xf4\xa6\x14\x9f\xb4\x10\x4d\x44\x72\xa8\x33\x16\xe1\xd7\x35\x65\xf6\x9a\x02\xf0\xbc\x6e\xfb\x99\x90\xaa\xb4\xf8\xfc\x1c\x3a\xc8\xc5\x95\x36\x13\x09\xe7\xd2\x21\xf0\x11\xd4\x81\x33\x39\xae\xcd\x36\x38\xd0\x45\x51\xc7\xe0\xa8\x71\x70\xd0\x24\x48\x8c\x2a\xbe\x74\x62\x85\x0d\x70\xd4\x01\xaf\x12\xd5\x79\x2e\xfe\x66\xe8\xbc\xa8\x1f\xbf\x82\xa2\xf0\x95\xaf\x42\xe3\x29\x6c\x74\x66\xf9\x68\x96\xa9\x84\x78\xa2\x69\x3c\x54\xae\x86\x81\xa3\x46\xce\xbf\x92\xf7\xff\x4c\xe2\x43\xe6\xe3\xdf\x53\x0b\xed\x50\x36\xec\xb1\x2d\x1c\x76\xba\x1f\x62\xbe\x8e\x82\x7a\xf5\x31\x00\x3c\x36\x1f\x11\x54\xf5\x67\x4c\xfc\x1e\xae\x3c\xd2\xd0\x53\x61\x71\x23\x4d\x49\xdd\x0a\xff\x97\xce\x7f\xf5\x7c\xf7\xd0\xef\x3d\xc4\x3b\x2f\x4e\x5f\xf3\xd2\x6b\xbb\xc6\x30\x64\x85\xd1\xa8\xd1\x2b\x0c\x37\xd2\x78\x15\x46\x08\x71\xe3\x7e\x8f\xf5\x9f\xb8\xfc\x8a\xf5\xee\x4d\x41\xbd\x3f\xb6\x22\x65\x51\xa4\xbb\xba\xfb\x8d\xc2\xd4\x01\x6b\xa1\xd3\x78\xf2\x10\x69\x2a\xc9\x1e\x63\x91\x3c\x14\x2b\x21\x75\xbf\x33\x8c\x5f\x6d\xb9\x5d\xc8\x38\x1a\x64\x9b\x5d\x33\x9e\x18\xe9\x78\xc7\x1e\xf7\x4f\xe8\x8e\x07\xb5\x74\x78\x8f\x17\xaf\x02\x8d\x76\x65\xce\x63\x2f\x88\x8d\x90\x4a\xd0\x51\x8b\xc7\x29\x9d\x42\xa2\x50\x68\x1e\x9d\x28\x79\x66\x83\xd6\xf5\x4f\x00\xf9\xb7\x60\xfc\x80\x1c\xab\xc7\x18\x8e\xd3\x6b\xf6\xd4\x8a\x0d\xdb\xff\x51\x09\xef\x23\xbc\x1a\xe1\x0d\x95\x25\xbd\x83\x42\xd0\x18\xda\x3f\xad\xa4\x78\x40\x22\x99\x1f\x60\xda\x18\xc2\xff\x2c\x45\x76\x0c\xb1\xab\x7a\x18\x8b\x9b\xf7\xc6\x8c\x40\xa1\xe0\x23\x11\xc4\xc3\x4d\x35\x7c\x3e\x75\x42\xab\xaa\x37\x8c\x6f\x47\xe5\xcb\x97\x58\x6b\xac\xae\x3b\xc2\x1c\xbf\x44\xd4\x20\x3d\x5a\x41\x87\x1f\x42\x57\xfc\xa5\x80\xbc\x74\x6c\x8e\xf3\x22\xa9\xe8\xa2\xe1\x78\x6d\x4f\xfd\x59\xea\xd5\xb8\xdf\x0b\xef\x1b\xf5\x9e\xf8\xfb\x7d\xbd\x87\x66\xc8\x9a\xf1\x02\xa0\x3e\xff\x27\xfe\x9e\x67\x46\x3e\x23\x1f\x5c\x02\xd0\x1a\xbd\x0a\x07\xe8\x83\x23\x3f\x2b\xc6\x63\xff\xe1\xcd\x22\xad\xf1\xbb\x16\xc0\x59\x74\x25\x5c\x30\x73\x50\x12\xfe\xfe\xb8\x22\x2a\x05\x2a\x86\x59\xb7\x02\x2d\x75\x28\x1d\x5c\x43\x90\x30\xbf\x0a\xab\xa1\xb1\xcf\x9a\xab\xe1\x55\xdc\xa8\xcc\x1b\xb1\x91\x39\xc7\xe6\xe1\xbc\x9b\xe4\xa6\x15\x1e\xbb\xc9\x8c\x62\x5e\x03\xf6\x11\xd5\xe6\xc1\xe2\x58\xe4\x29\xaa\x64\xeb\x15\xb3\x3d\xa2\xca\xd6\x1b\xa3\x87\xbf\x3f\xdd\x64\x2d\xdc\x74\xb1\x25\xd3\x65\x24\xf2\x4c\x94\x0b\x91\xad\x0c\x04\x54\x07\x5f\x19\xd1\xf2\x1f\x18\x2d\x36\xeb\xa7\x5a\x02\x8b\xe1\x87\x05\x1e\x48\xa9\x7c\xcc\x92\x9b\x7f\xe9\xe8\xcc\xb8\xaf\x8b\x14\x9d\xb4\x98\x42\x26\x51\xa5\x60\x52\xb4\x7c\x22\xfd\xec\x8c\x0e\x3f\x21\xa1\x95\x64\x31\xfc\x54\x06\x37\x6b\x2a\x4d\x32\xaa\x65\x82\x7e\x07\x19\x0a\xfe\x2d\xc8\x1b\x28\x84\x73\x90\xa3\xa0\x33\x68\x56\x2a\xb5\x03\x63\x53\x24\xe3\xf5\xa1\x8c\x4a\xd2\x40\xe9\xd0\x3a\xd8\xae\x4d\x6c\x93\x3c\xa5\x15\x34\x74\x4a\x3f\x8a\xf7\x2e\xd2\x15\x4a\xec\x40\x7a\x6a\xc9\x71\x53\xcd\x2a\xad\x7f\x80\xe1\x5f\x71\x0c\x75\xdd\xe3\x12\xad\xce\x75\xed\x1a\xe5\xd7\xf4\xd4\xae\xce\x78\xae\x69\xd7\xe5\xfe\x46\xaa\x5d\x84\x55\xdb\x68\x57\x5a\xb3\x09\xb5\xcb\x89\x57\xf8\xa9\x5d\x48\x8d\x79\x99\x17\x18\x1c\xb5\x02\x3f\x1d\x94\x16\x7b\x19\x6b\x2b\xfc\xdc\x58\x8b\xf3\xd3\x28\x02\x86\xb2\x38\xa0\xe0\xdc\xe1\x8e\x98\x38\xc4\xa8\xd1\x56\xc2\x8b\x4f\x77\xb8\xbb\xed\xee\x22\x11\x8e\x0d\xb9\xba\x6d\x54\x90\x0e\x6b\x4f\x14\x72\xed\x85\x9c\x4f\xcf\x41\xfe\x7f\x53\xa1\xea\x7c\x20\x5f\xbc\xa8\xbe\xd9\x5c\xff\x24\x6f\xab\xea\xac\x11\x7f\xb0\x3e\x6c\x79\x14\x6b\x24\xc8\x50\x51\xf4\x1f\xfa\xff\x0c\x00\x00\xff\xff\x87\x00\x70\xa9\x7b\x1e\x00\x00")

func call_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		_call_tracerJs,
		"call_tracer.js",
	)
}

func call_tracerJs() (*asset, error) {
	bytes, err := call_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "call_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _emvdis_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4b\x6f\xdb\x38\x10\xbe\xfb\x57\x0c\x7c\xb2\x12\x57\x15\xf5\x8a\x2d\xaf\x17\x5b\xa4\xc6\x22\x45\x5f\x68\xdc\x93\xe1\x83\xe2\xd0\xb1\x10\x5b\x14\x28\x2a\x6d\x36\xf0\x7f\x5f\x0c\x67\xf4\x70\xec\x4d\xd0\xc5\xee\xa1\x9c\x58\x9c\xf9\xe6\xf1\x7d\x24\xfb\xf6\xec\xac\x37\xdf\x64\x25\x18\x9d\xae\xa4\x06\x2d\x4d\xa5\xf3\x12\xca\x6a\xbd\xce\x56\x32\x37\x90\xe5\x6b\xa5\x77\xa9\xc9\x54\x0e\x6b\xad\x76\x90\x92\x2f\x18\xd5\x2b\xa4\xc6\x4d\x90\x0f\xbb\xdb\xac\x7c\x53\x9a\xc7\xad\x84\xdb\xac\x4c\xcb\x52\xee\x6e\xb6\x8f\xbd\xde\x1f\x69\x65\x36\x4a\xc3\xe7\x6c\x75\x0f\x1f\xd4\x26\x2f\x55\x3e\x84\x4f\xa9\x36\x59\x0e\xd7\x3f\x64\x7e\x2b\x7b\x67\x67\x6f\x7b\xbd\xa7\x1e\x40\xbf\x34\xe9\xea\xbe\x9f\xc0\xe2\xa9\xaf\x8a\x12\xff\x58\xee\x97\x43\xdc\xc9\x8b\xaa\xdc\x48\xfc\xf4\xe4\x25\xe0\x0d\x41\x24\x20\x86\xe0\xdb\x35\xb0\x6b\x68\xd7\xc8\xae\xb1\x5d\x2f\xec\x3a\xb2\xeb\xd8\xae\xc2\x23\x43\xd1\x82\xdc\x04\xf9\x09\x72\x14\xe4\xe9\x93\xa7\xcf\x79\x28\x91\x4f\x99\x7c\x4a\xe5\x53\x2e\x9f\x50\x02\x72\x09\x09\x25\x24\x94\x88\x50\x22\x42\x89\xc8\x25\x22\x94\x88\x0b\x8e\x6c\x3f\x11\xa1\x44\x17\xf4\x8b\x50\x22\x42\x89\xa9\xe5\x98\x02\x62\x6e\x91\x02\x62\x2a\x3e\xa6\x80\x98\x02\x46\x14\x30\xa2\xb4\x23\x9f\x7e\x05\x64\x08\x65\x44\x69\x47\x31\x19\x4a\x3b\x22\x94\x11\xa1\x8c\xa9\xf8\xb1\xb0\x7b\x63\xca\x37\xa6\x7c\x63\x9e\x6a\x3d\x56\x9e\xab\xc7\x83\xf5\x7c\xb6\x01\xdb\x90\x6d\xc4\x96\x27\xef\xf1\xe8\x3d\x9e\xbd\xc7\x78\x0d\x4f\x8c\x27\x18\x4f\x30\x9e\x60\x3c\xc1\x78\x35\x93\x35\x95\x35\x97\x4c\xa6\x60\x36\x05\xd3\x29\x98\x4f\xc1\x84\x0a\x66\x54\x30\xa5\x82\x39\x15\x3e\xe3\xf9\xa3\x04\x7c\xb4\xe3\x04\x82\x21\x88\xc0\x4b\x20\x44\x2b\x12\x88\xd0\xfa\x09\xc4\x68\x83\x04\x2e\xd0\x86\x09\x8c\xd0\x46\x09\x8c\xd1\x22\x1e\xaa\x36\x40\x40\x44\x0c\xb0\x42\x84\x0c\xb0\x44\xc4\x0c\xb1\x46\x04\x0d\xb1\x48\x44\x0d\xb1\x4a\x84\x0d\xb1\x4c\xc4\x0d\x43\xaa\x23\x8c\xa8\x8e\x30\xa6\x3a\xc2\x0b\xaa\x03\xd5\x67\x03\xc6\x54\x07\xea\x0f\xeb\x40\x01\x62\x1d\x56\x81\x58\x87\xd5\x20\xd6\x61\x55\x88\x90\xa8\x43\x5b\x87\x55\x22\x82\xa2\x16\x6d\x1d\x56\x8d\x08\x6b\xf5\x88\xb8\xac\x48\x11\x0b\xb6\x3e\xdb\x80\x6d\x68\xad\x1f\xf2\x29\x0a\xf9\x18\x85\x7c\x8e\xc2\x80\xf7\xd9\xcf\x1e\x82\xbd\x3d\xe9\x5a\x96\xd5\xd6\xf4\x13\x58\x57\xf9\x0a\x6f\x9d\x81\x03\x4f\x7c\x2f\x81\xd9\x64\xa5\x6b\x6f\x89\x85\xb7\x74\x55\x51\x4e\x80\xa2\x4a\x23\x8b\x6e\xcc\x56\xdd\x0d\xe1\xf6\xc6\x01\xbc\x57\x00\x1e\x52\x0d\x6b\x9d\xee\x24\x4c\xbb\x18\xed\x9f\xee\x56\xe6\x77\x66\x03\x6f\x40\x2c\x27\x36\x24\x5b\x23\x88\x2b\xb5\xae\x41\x80\x20\x16\x7d\xa9\xb5\xd2\xfd\x25\x4c\x81\x3d\x5c\xa3\xae\x8d\xce\xf2\xbb\x81\x43\xc1\x7b\x90\xdb\x52\xd6\x18\xb7\xb2\x30\x1b\x98\x76\x53\x73\xbe\x16\x5a\x15\x78\xd1\xc2\xb4\xf9\x00\xd0\x57\xd8\x12\x02\xa8\xc2\x35\xea\x73\xb5\xbb\x91\x7a\xe0\x0c\x5b\x07\x0b\xdc\x07\x72\xb2\x3f\x3a\x9b\xcd\x20\x17\xcb\xfa\xeb\x7e\xc2\x7f\x64\xeb\x81\xed\x05\x47\x58\xb7\xfe\x3b\x78\x4e\x27\x3b\x8e\xac\xd0\xf2\x41\x15\x30\x85\xc6\x79\x71\x14\xd6\x4e\xcc\x4e\x48\xe9\x01\x46\x66\x30\x05\x6f\x02\x19\xfc\x46\x4d\xf3\xfd\xbd\x20\x44\x57\x15\xcb\x09\x64\xe7\xe7\x4e\x13\x08\x9c\xcc\xa5\xb2\x5d\xf4\xb7\xc3\xa3\x71\x15\x52\xde\x0f\x32\xc7\x9d\xcb\x9f\x66\x20\x62\xc7\xa9\x53\xee\xd9\x96\x3f\x32\xb3\xa2\x08\x3b\xad\x9a\x8f\xb6\xa5\x55\x5a\x4a\xe8\x5f\xbe\xfb\xf8\xb1\x9f\x1c\x7d\xba\xfc\xf2\x7e\xd6\x7c\xa6\xe6\xb3\xbc\x34\xa9\x36\xcc\x72\xa7\x8c\xc0\x71\xaf\x72\x13\x87\x03\x67\xf2\x3c\x20\xfb\x4b\x1e\xfb\x87\x27\xfc\x89\xee\x45\xff\x2e\x2d\x1b\x21\x75\x42\xbc\x17\x42\x8c\x3a\x15\x21\xda\xd9\x1c\x87\x3c\xa4\xdb\x4a\x9e\x8a\xf2\x1d\xf7\x50\xb8\xdd\xa8\x2c\x2f\x2a\xd3\x44\xed\xe4\x4e\xe9\x47\xb7\xdc\x66\x2b\x39\xe0\xd9\x0c\x9b\x21\x9d\x73\xf7\x27\x60\xda\xd3\x92\x57\xdb\xed\xf1\x3e\x1d\xed\x17\x1c\xf0\xbf\x02\xb8\xbb\xe8\xe8\xac\x73\x92\xac\x52\xc8\xb7\x93\xfd\x46\xcb\xf4\x7e\x72\xc0\xf3\xfb\xd9\xc7\xd9\x9f\xef\xe6\xb3\x13\x12\xb8\x9e\xbf\x9b\x5f\x5d\x1e\x6c\xbc\x22\x02\xff\x17\x45\x70\x4a\x34\x6d\x83\xb6\x3f\x38\x12\xef\x3f\xeb\xe5\x5f\x08\xe6\x97\x14\xd3\x72\xff\x5f\x91\xff\x3a\xfb\xff\x33\xfd\xdf\x66\xf3\xef\xdf\x3e\x3f\xe3\x57\x55\x27\xb8\x3d\x35\x4c\x76\x3e\x4d\xae\x38\x11\x40\x17\x25\xbf\x5b\x27\x8e\x90\xaa\xcc\xd0\xa6\x3f\xaf\x71\x5f\x2e\xff\x7a\xfe\xe5\xeb\x73\xd5\x7e\xbf\xba\xbc\x3a\xb8\xb7\x5e\x4b\xea\x0d\xc1\x7b\x39\xcd\x87\xef\x9f\xbe\xbe\x9f\x5d\xcf\x3b\xa0\x35\x11\xc5\xaa\xb9\x0c\x8a\xd5\xf3\xfb\x97\x9f\x3a\x55\xb8\x59\xf9\x15\x29\x71\xba\x8f\x49\x0d\xb1\x95\x79\x83\x71\xf0\xa8\xc1\x1b\xf0\x7e\x46\xf2\x39\x6a\xfb\xda\x1c\xd3\xcc\x6f\x6c\x9d\xa4\x95\xc4\xc1\x1b\xdf\x36\xde\xbc\x91\x75\x7c\x0f\xff\xed\x7b\x7f\x07\x00\x00\xff\xff\x84\xd3\x90\x9f\x03\x0d\x00\x00")

func emvdis_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		_emvdis_tracerJs,
		"emvdis_tracer.js",
	)
}

func emvdis_tracerJs() (*asset, error) {
	bytes, err := emvdis_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "emvdis_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _noop_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xc1\x4a\x03\x31\x10\x86\xcf\xcd\x53\xfc\x47\x85\xd2\xbd\xfb\x08\x82\x27\xc5\xfb\x24\x3b\xdb\x4d\x8d\x99\x75\x32\x59\x2b\xa5\xef\x2e\x49\xad\xf4\x20\xde\xc2\xc0\xff\x7d\x1f\x19\x06\x64\x91\xe5\x45\x29\xb0\x22\x16\x1c\x6a\x31\xd8\xcc\xf0\xa4\xec\x25\x33\xbc\xc4\xc4\xba\x24\x32\x46\x90\x91\xa1\xfc\x51\xa3\xf2\x88\x49\xe5\x1d\x84\x47\x5a\xe9\x39\x68\x5c\xcc\x0d\x03\xc4\x1f\x38\x18\x4c\xe0\x19\xb5\x90\x4f\x0c\x2a\x20\x98\x52\x2e\x14\x2c\x4a\x6e\xef\xc0\xba\x73\x27\xb7\x19\x06\x14\xe3\xa5\xb9\x63\x5e\xe5\xad\x71\x45\xc1\x2b\xeb\x17\x64\xe9\x46\x9b\xe9\x12\xf5\xfa\x04\x3e\x72\xa8\xc6\x65\xe7\x36\x6d\xf7\x80\xa9\xe6\x0e\xbd\x4b\xb2\xdf\x62\xf4\xf7\x38\xe1\xbc\x75\x9d\x3c\x51\x4d\x76\x8b\xfe\x9c\x39\x77\x12\x05\xab\x94\x7e\x68\x2d\x49\x26\x50\xbe\x0a\x27\x8a\xa9\x19\xfa\xfe\x7f\x85\x72\xf9\xcb\x41\x29\x75\xcf\x05\x58\x30\xd3\xca\xf0\xcc\x19\xd1\x58\xc9\x78\x84\xac\xac\xa0\x3c\x42\xd9\xaa\xe6\xd2\x71\x6d\x33\xc5\x4c\xe9\x0a\x96\xa9\xdf\xda\x8f\xc5\xbc\xdf\xb9\xcd\xe5\x7e\x13\x15\xec\xf8\x1b\xe5\xce\xee\x3b\x00\x00\xff\xff\x01\x86\xa4\x01\xd3\x01\x00\x00")

func noop_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		_noop_tracerJs,
		"noop_tracer.js",
	)
}

func noop_tracerJs() (*asset, error) {
	bytes, err := noop_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "noop_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opcount_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xb1\xae\xe2\x30\x10\x45\xeb\xf8\x2b\x6e\xb9\x2b\x50\xb2\x35\xfd\x96\x74\x2b\xfa\x49\x32\x26\x5e\xcc\x18\xd9\xe3\x2c\x08\xf1\xef\x2b\x3b\xe4\x29\x7a\xa2\x8c\x35\xf7\x9c\xb9\x93\xae\x43\xb8\x0d\x21\x8b\xfe\x89\x34\x70\x84\x4b\x20\x24\xba\xde\x3c\x43\x97\x27\x9d\x48\xf1\x37\x27\x45\x1d\x4c\xd0\x89\x21\xf9\xda\x73\x44\xb0\x70\x92\x34\xe6\x41\x5d\x90\x64\xba\x0e\x7c\xe7\x21\x2b\x8f\xe8\x1f\x75\xf2\xf7\xe9\x88\x9e\x6d\x88\x5c\x3f\x35\x92\x24\xaa\xe3\x50\x8e\x57\x27\xa4\x3c\xb6\xe6\x69\x9a\xae\x5b\x0c\x55\x7c\xf9\xee\x29\x9c\xad\xeb\x4b\xd4\x9a\xa6\xc6\x0e\xf8\xb5\x37\x95\x92\x94\x6f\xa5\x89\x93\x39\x5c\x78\x84\x0d\x11\x3c\x73\x7c\xd4\xb2\x23\x2f\x95\x0a\xfe\x74\x5c\x31\xa9\x35\x4d\xc9\x1d\x60\xb3\x54\xc3\x0f\x1f\xce\x7b\x8c\xfd\x4f\x3c\xa1\x93\x4b\x6d\xb5\xec\x76\x78\xbd\x35\x96\xb2\xd7\xad\xe7\xdf\xc4\x52\xb1\x34\x68\x26\xff\x46\x97\xa6\xc1\x82\x64\xb5\x5b\x72\xbe\xe8\x6a\xfe\xb3\x6f\x55\x44\x4e\x9f\x1c\xe4\x7d\xf5\x2c\xc0\x84\x89\x66\x46\xcf\x2c\x70\xca\xb1\x1c\x14\x61\xe6\x08\x92\x11\x91\x35\x47\x49\x15\x57\x32\xd6\x09\xf9\x15\x1c\xec\xfa\x53\x06\x27\xe7\xd6\x34\xcb\xfb\x66\xa9\x41\xef\xeb\x52\x0b\x69\x73\x0b\xbc\xcc\xcb\xfc\x0f\x00\x00\xff\xff\x80\xec\x28\x5c\x43\x02\x00\x00")

func opcount_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		_opcount_tracerJs,
		"opcount_tracer.js",
	)
}

func opcount_tracerJs() (*asset, error) {
	bytes, err := opcount_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "opcount_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _prestate_tracerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x57\xcf\x8f\x1a\xb9\x12\x3e\x77\xff\x15\xf5\xe6\x02\x28\xa4\x99\xbc\x43\x0e\x8c\x38\x10\x06\xbd\x17\x69\x36\xb3\x0a\x68\xb3\xbb\x51\x0e\x6e\xbb\x1a\x1c\x8c\x8d\x6c\x37\x3f\x14\xf1\xbf\xaf\xca\xed\x6e\x68\x02\x93\xec\xee\x9c\xa6\xed\xf2\x57\xe5\xfa\xaa\xbe\x32\x83\x01\x6c\x2c\x3a\xcf\x3c\xce\x2d\xe3\x68\xc1\x94\x7e\x53\x7a\x07\xae\x2c\x0a\xc9\x25\x6a\x0f\x52\x17\xc6\xae\x99\x97\x46\x83\x37\xc0\x2d\x32\x8f\xc0\x40\x19\xce\x14\xe0\x1e\x79\x19\xf6\x4c\x91\x0e\x06\xe0\x97\x08\xde\x32\xed\x18\x0f\xab\x85\x35\x6b\x60\xc0\x4b\xe7\xe9\x1f\xe7\x70\x9d\x2b\x14\xb0\x40\x8d\x4e\x3a\xc8\x95\xe1\xab\x2c\xfd\x96\x26\x67\xc1\x80\x74\x01\xa8\x36\xf2\x4b\xe6\x61\x87\x1d\x8b\x90\x97\x52\x09\xa9\x17\x59\x9a\xd4\xd6\x43\xd0\xa5\x52\xfd\x34\x40\x28\x63\x56\xe5\x66\xcc\xb9\x29\x43\xec\x5f\x91\xfb\x0a\xcc\x6d\x90\xcb\x42\xa2\x00\xd6\xec\x7a\x13\xb6\x1a\xbf\x26\x27\xfb\x2c\x4d\x5a\x30\x43\x28\x4a\x1d\xae\xd3\x65\x42\xd8\x3e\x88\xbc\xf7\x2d\x4d\x92\x2d\xb3\x84\x05\x23\xf0\xe6\xff\xb8\x0f\x9b\xbd\x87\x34\x49\x64\x01\x5d\xbf\x94\x2e\xab\x81\x3f\x33\xce\xbf\xc0\x68\x34\x82\x52\x0b\x2c\xa4\x46\xd1\x03\x82\x48\xae\x99\x55\x3b\x49\xce\x14\xd3\x1c\x87\xd0\xb9\xdf\x77\xe0\x15\x88\x3c\x5b\xa0\x7f\x57\xad\x56\xce\x32\x6f\x66\xde\x4a\xbd\xe8\xbe\x79\xdb\xeb\x87\x53\xda\x84\x33\x10\xcd\x3f\x98\xc6\xb8\xda\xe7\x46\x84\xed\x18\x73\x65\x35\x31\x22\x1a\x45\x2b\xe7\x8d\x65\x0b\x1c\xc2\xb7\x23\x7d\x1f\xe9\x56\xc7\x34\x39\xb6\xb2\x3c\xab\x8c\x6e\x64\x39\x42\x00\x6a\x6f\x0f\x60\x8a\x8a\x51\xb9\x45\xdd\x22\x20\xe0\xbd\x44\xc2\xac\x0e\xe5\x82\x84\x15\x1e\x7e\xcc\x04\x6d\x48\xb1\x6f\x36\x56\x78\xe8\x3d\xa4\x37\x29\xca\x62\xd0\x9f\xa5\xd8\x5f\xe7\x8b\x00\xb7\x4c\x35\x80\x55\xfe\x66\x84\x70\x8a\xab\x17\x7c\x07\x1f\x64\xfb\x9f\x11\xdc\xdd\xef\xef\xff\xe5\xdf\x5d\x8c\xe0\x4a\xc9\x5c\x84\xfd\x13\xa1\x1d\xdb\x7c\x5a\x74\xa5\xf2\xd4\x76\x52\x6f\xcd\x0a\x05\xec\x96\xc4\x93\x52\x81\x1a\xb3\xa1\xaa\x71\xb0\x64\x5b\x84\x1c\x51\x83\xf4\x68\x99\x47\x01\x66\x8b\x16\x98\x16\x60\xd1\x97\x56\xbb\x86\xce\x42\x6a\xa6\x6a\xe0\xc8\xbe\xb7\x8c\x57\xbd\x5b\xad\x9f\x71\xca\xfd\x3e\xb0\x19\xee\x38\x18\xc0\xd8\x03\xdd\x13\x36\x46\x6a\xdf\x87\x1d\x82\x46\x14\x24\x40\x02\x45\xc9\x7d\xc0\xeb\x6c\x99\x2a\xb1\x53\x89\x8c\x5f\x62\x75\xd4\x94\x1e\xed\xb9\x08\xf5\x43\x80\x6b\xb3\x45\x90\x1e\x72\xc6\x57\x10\x1b\xdf\x58\xb9\x90\x3a\x8d\x39\x6d\x35\x3d\x45\x94\x11\x70\x08\x2b\xd4\x0c\x71\x4f\x2b\xef\x02\xff\xb9\x5c\xbc\xd7\xfe\xa2\x88\xaa\xcc\xd7\x47\x7b\x5f\xb2\xd8\xc4\x99\x53\x92\x63\xf7\xbf\xbd\x3e\xbc\x79\xdb\x54\xa6\x37\x04\x05\x3f\x06\xf3\xe6\x36\x54\x7a\x59\x11\xd7\x8f\x05\x37\xa4\x24\xaf\x82\xd7\xcc\x95\x39\xd1\x51\xdd\x33\xe4\xb1\xad\x26\x0f\x2f\xe0\xb6\xef\x56\xe3\xc6\xd4\x64\x4c\x88\xdb\xa0\x15\x45\x8f\xc8\x2d\xae\x69\xba\x10\x0b\x9c\x29\x85\xb6\xe3\x20\x68\x57\x3f\x96\x53\xe0\x0b\xd7\x1b\x7f\xa8\x67\x8e\x67\x76\x81\xde\xfd\x38\xb0\x80\xf3\xfa\x75\x2d\xc5\x21\x15\x87\x0d\xc2\x68\x04\x9d\xc9\xc7\xe9\x78\x3e\xed\xc4\x66\x1a\x0c\xe0\x13\x05\xa0\x21\x57\x32\x17\xea\x00\x02\x15\x92\x2f\x8a\xcb\xe8\x90\xa2\x46\x9a\xfa\xc0\x1c\x30\x7d\x00\xdc\x4b\xe7\xa5\x5e\x40\xa5\x58\x3b\x53\x2a\x11\xe1\x42\x8f\x70\x56\x3a\xaa\xd6\x8b\x61\xe8\x0d\xe4\x08\x16\x49\xdf\x68\x0e\x85\x76\x63\x4a\x0a\x90\x3a\x36\x8d\x75\x1e\x36\x8a\x71\xcc\x08\xaf\x09\xe6\x36\xbf\x51\x99\xc9\xf5\xc7\xd0\x82\x01\xe8\x34\x68\x99\xa2\x41\x4d\xee\x1d\x74\x6b\x8c\x5e\x9a\x24\xb6\xb6\x3e\xc3\x7e\x38\x49\x82\xf3\xb8\x39\x17\x84\xc2\x58\xc0\x2d\x92\x94\x07\x35\xa8\x86\x32\xf9\xfa\xed\x97\xf8\x0a\x40\x97\xa5\x09\x9d\x3b\xeb\x6b\x65\x16\xed\xbe\x16\x55\x5a\x78\x69\x2d\xf1\xdf\x8c\x82\x82\x7a\xfc\x6b\xe9\x3c\xe5\xd4\x52\x7a\xa2\x5a\x5c\x13\xeb\x20\xcd\x34\xf5\x7b\xdf\x0f\x51\x9a\x9f\x61\x5e\x91\xbb\x38\x2d\x61\x27\x95\x82\x8d\xf1\xa8\xbd\x64\x4a\x1d\x88\x87\x9d\x35\x7a\x01\x4b\xb4\xd8\x07\x27\xc9\x2a\x28\x4e\x30\x95\x9a\xab\x52\x54\x65\x10\xea\x38\xe2\xb9\x10\xb3\xa2\x83\x3b\xe9\x97\x61\x7f\x8d\xce\xb1\x05\x66\x54\x49\x85\xdc\x57\x89\x91\x1a\x3a\x95\xc8\x75\x7b\x9d\xac\x09\xb2\x2d\x31\xca\x2c\xb2\xba\xc8\x48\xab\xc7\x42\x58\x74\xae\xdb\x8b\x9a\xd3\x30\xfb\x69\x89\x9a\x92\x0f\x1a\x77\xd0\x3c\x91\x18\xe7\xe8\x1c\x8a\x3e\x30\x21\x48\xda\x2e\x9e\x33\x69\x92\xb8\x9d\xf4\x7c\x09\xc1\x93\xd9\x9c\x7a\xb1\x17\xeb\x9f\x33\x87\x70\x37\xfd\x7d\x3e\x79\x7e\x9c\x4e\x9e\x7f\xfd\xe3\x6e\x08\xad\xb5\xd9\xfb\x3f\xa7\xcd\xda\xbb\xf1\xd3\xf8\xc3\x64\x7a\x37\x3c\xcd\xa1\xf6\x85\xbc\xa9\xaf\x40\x0e\x9d\x67\x7c\x95\x6d\x10\x57\xdd\xfb\xb6\x0e\x9c\x2e\x98\x24\xb9\x45\xb6\x7a\x38\x05\x53\x35\x68\xf4\x51\x4b\x2e\x8c\xe0\x66\xb2\x1e\x6e\x47\x33\x89\xf6\xdd\x5a\xc8\x4f\x4f\xa2\x20\x15\x2f\xc6\x31\x7e\x7a\x6a\x6e\x4e\x1f\x94\x8e\x66\xe1\x71\xfa\x34\xfd\xdf\x78\x3e\x6d\x59\xcd\xe6\xe3\xf9\xfb\x49\xb5\xf4\xb7\x53\xf4\xe6\xa7\x53\xd4\x99\xcd\xe6\xcf\x1f\xa7\x9d\x61\xfc\x7a\x7a\x1e\x3f\x76\xbe\x73\x18\xdf\x4d\x2f\x15\x99\x37\x9f\x8c\x15\xff\x84\xab\xb3\xb7\x43\xc1\xae\x3d\x1d\x82\x08\x71\x5f\x5e\xfc\x44\x00\xa6\x6b\xfd\x28\x98\x54\x24\x18\xe1\xfc\x55\xc5\x38\xa6\xc7\xf4\xaf\x00\x00\x00\xff\xff\x36\x30\x7f\xdf\xa3\x0c\x00\x00")

func prestate_tracerJsBytes() ([]byte, error) {
	return bindataRead(
		_prestate_tracerJs,
		"prestate_tracer.js",
	)
}

func prestate_tracerJs() (*asset, error) {
	bytes, err := prestate_tracerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prestate_tracer.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"4byte_tracer.js":    _4byte_tracerJs,
	"call_tracer.js":     call_tracerJs,
	"emvdis_tracer.js":   emvdis_tracerJs,
	"noop_tracer.js":     noop_tracerJs,
	"opcount_tracer.js":  opcount_tracerJs,
	"prestate_tracer.js": prestate_tracerJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"4byte_tracer.js":    {_4byte_tracerJs, map[string]*bintree{}},
	"call_tracer.js":     {call_tracerJs, map[string]*bintree{}},
	"emvdis_tracer.js":   {emvdis_tracerJs, map[string]*bintree{}},
	"noop_tracer.js":     {noop_tracerJs, map[string]*bintree{}},
	"opcount_tracer.js":  {opcount_tracerJs, map[string]*bintree{}},
	"prestate_tracer.js": {prestate_tracerJs, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
