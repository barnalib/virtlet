// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// deploy/data/virtlet-ds.yaml
package tools

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

var _deployDataVirtletDsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\xdd\x6f\xdb\x38\x12\x7f\xcf\x5f\x31\x68\x80\x6b\x0b\x9c\xe2\xa4\xb8\xbd\x76\x8d\xbb\x87\x34\xf6\x66\x8d\x26\xb6\xe1\x7c\x74\xdf\x0c\x9a\x1a\xcb\x5c\x53\xa4\x96\xa4\x94\xf8\xfe\xfa\x03\x49\x49\xd1\x97\x1d\x27\x4d\x82\x6e\x5e\xea\x92\x9c\x1f\x67\x86\xf3\x49\x31\x08\x82\x03\x92\xb0\x5b\x54\x9a\x49\xd1\x07\x92\x24\xba\x97\x9d\x1c\xac\x99\x08\xfb\x30\x20\x18\x4b\x71\x85\xe6\x20\x46\x43\x42\x62\x48\xff\x00\x40\x90\x18\xfb\x90\x31\x65\x38\x9a\xfc\xff\x3a\x21\x14\xfb\xb0\x4e\x17\x18\xe8\x8d\x36\x18\x1f\xe8\x04\xa9\x5d\xae\x91\x23\x35\x52\xd9\xdf\x00\x31\x31\x74\x75\x41\x16\xc8\xb5\x1f\x00\x50\xa9\x30\xac\x0e\x69\x30\x4e\x38\x31\x98\xd3\x54\x36\xb7\x7f\x4d\x06\xec\x1f\xaf\x41\x76\x82\x02\x14\x2c\xd9\xbf\x95\xd4\x66\x8c\xe6\x4e\xaa\x75\x1f\x8c\x4a\x31\x1f\x0f\x85\x9e\x4a\xce\xe8\xa6\x0f\x67\x3c\xd5\x06\xd5\x6f\x4c\x69\xf3\x9d\x99\xd5\xef\x9e\x24\x5f\x78\xe8\x20\xa6\xa3\x01\x30\xed\x00\xc0\x48\xf8\x70\xf2\x11\x50\x90\x05\x47\xb8\xbd\xd4\x76\x44\xa7\x2a\x63\x19\x16\x7c\x00\x95\xc2\x10\x26\x50\x81\x42\x6d\x88\x7a\x80\xfb\x60\x24\x2c\x10\xe8\x0a\xe9\x1a\xc3\x8f\x40\x44\x08\x1f\x3e\x7d\xb4\x20\x39\xa4\x59\x21\xa4\x1a\x41\x2e\x41\x68\x14\x06\x15\x30\x01\x4c\xb0\x0a\x6c\x45\xbc\xe9\x68\x50\x13\xed\x10\x16\x52\x1a\x6d\x14\x49\x20\x51\x92\x62\x98\x2a\x04\x81\x18\x3a\x4e\xa9\x42\x62\x10\x88\xc5\x5a\xb2\x28\x26\x89\x45\xaf\x1c\xe9\xc3\x49\xe7\x80\x1a\x55\xc6\x28\x9e\x52\x2a\x53\x61\xc6\xb5\x63\x29\xf7\x94\x82\x6f\xec\x71\xc0\x6d\xae\x81\x44\x86\x1a\xa4\x70\xd2\x08\x19\xa2\x86\x3b\x66\x56\x80\xf7\x46\x91\x99\x3f\xb6\xff\x16\xda\x72\xc7\x9a\x43\x91\xe5\xd2\x8a\xba\x79\x38\x64\x4b\x7d\xda\x1a\x05\x50\xf8\x57\xca\x14\x86\x83\x54\x31\x11\x5d\xd1\x15\x86\x29\x67\x22\x1a\x45\x42\x96\xc3\xc3\x7b\xa4\xa9\xb1\x56\x5f\xa1\xf4\x98\x57\xb9\xc9\x5e\xa3\x8a\x75\x7d\x3a\xf0\x16\x3c\xbc\x4f\x14\x6a\xeb\x33\x8d\x79\xbb\x62\x8d\x9b\x7e\x4d\x9c\xc6\x0a\x00\x99\xa0\x22\xd6\x27\x60\x24\x5a\x93\x19\xe1\x29\xb6\x60\x2d\x70\x43\xb7\x56\xee\xb3\xe2\xdc\x4b\x82\x43\xb8\x5e\x61\xc3\x28\x80\xca\x84\xa1\x2e\x00\xde\x6b\x58\x72\xbc\xcf\x24\x4f\x63\x84\x50\xb1\xac\xb4\x9b\x43\x6b\x09\xf6\x64\x42\x5c\x92\x94\x1b\x77\xfe\xee\xd4\x78\x1a\x31\x01\x21\x53\xce\x30\x51\xe8\x54\xa1\x06\xb3\x22\x0f\x16\xec\xe8\x98\x72\xba\xb3\xdb\x59\xd3\xc2\x10\x16\x1b\xe0\x6c\x61\xf7\x86\x7f\x94\x7e\x80\xf7\x4c\x9b\xc2\x0c\xac\xb5\x1e\x14\x52\x7a\xf7\x4e\x14\x26\x44\x61\x60\xcf\xa3\x54\x05\x8b\x49\x84\x7d\x88\x99\x22\xc2\x30\xdd\xab\xc7\x80\x7c\x7e\x9a\x72\x5e\xb8\xf0\x68\x39\x96\x66\xaa\xd0\x7a\x4b\xb9\x8a\xca\x38\x26\x22\x7c\xd0\x70\x00\xbd\xea\x76\x47\x7a\x55\x4e\x79\x1d\x5d\x5a\xfb\xd6\x55\x02\xcf\xe4\xfa\x8b\x0e\x1e\x34\x19\x78\x1d\xe9\x20\x64\xaa\x72\x7a\xb1\x25\x9e\x12\xb3\xea\x43\x2f\xd7\x66\x50\x27\x68\xe1\xaa\xb4\x6a\x16\x87\x30\x90\xe2\xbd\x01\x12\x86\xf0\xce\xa3\x29\x99\x90\x88\x38\xeb\x85\xaf\xcc\xeb\x9c\x49\x41\xf8\xbb\x7f\x02\x33\x70\xc7\x38\x07\x4e\xe8\xda\x6f\x0e\x28\x8c\xda\x6c\x61\xa9\xba\x57\xb1\x7f\x28\xe9\x1a\x95\x96\x74\xbd\x85\x28\x23\xca\x12\xf6\xfc\xc2\xa3\xda\xca\x02\x84\xcb\x68\x0b\xb5\x3d\xee\xea\xec\x21\x2c\xa5\xf2\x26\xc5\x44\xe4\x6c\xca\x6f\xc1\xd9\xa2\x97\x9b\x4e\xcf\x9d\xad\xf6\x76\xe3\xe2\x47\xcd\x32\x8a\x4d\x33\xa2\x02\xce\x16\x3b\x36\x0e\x9a\x4b\x4a\xa1\x31\xdb\x42\x56\x9d\x09\x5a\x7a\xb0\x4c\x36\x0d\xb1\x3b\x49\xd9\x88\x49\x53\xc5\xcc\xc6\xba\x2d\xde\x9b\xaa\x93\x27\x8a\x65\x8c\x63\x84\x61\x2d\x68\x03\xa0\xc8\xda\x96\xf7\xed\xe6\xeb\x70\x3e\x9e\x0c\x86\xf3\xf1\xe9\xe5\xb0\x02\xe3\xa2\xc7\x6f\x4a\xc6\xf5\x00\xb2\x64\xc8\xc3\x19\x2e\x9b\x61\xa5\x9a\xfc\xb3\x93\xc6\xa4\x23\xf2\x92\xda\xd4\x79\x64\x35\x6e\xa3\x7c\x8b\x9b\xdb\xd1\xec\xfa\x62\x78\x3d\x1f\x8c\xae\x4e\xbf\x5e\x0c\xe7\xdf\x6e\x2f\x1f\x67\xc9\xa7\x99\x4b\x92\x7c\xc3\x4d\x07\x67\x35\x05\x06\x7e\x71\x63\x89\x0b\xb4\x21\xd3\x36\x39\xce\xd7\x59\xdc\x98\x96\x89\xf7\x89\x86\x3e\x9b\x4c\x5f\xcd\x46\x93\xdb\xf9\xd5\xcd\x74\x3a\x99\x5d\xbf\x19\xdb\x5a\x31\x99\xcd\x75\x9a\x24\x52\x99\xe7\x31\x3e\x98\x7c\x1f\x5f\x4c\x4e\x07\xf3\xe9\x6c\x72\x3d\x39\x9b\x5c\xbc\x9d\xce\xe5\x9d\xe0\x92\x84\xf3\x44\x49\x23\xa9\xe4\xcf\x13\xe0\x62\x72\x7e\x31\xbc\x1d\xbe\x1d\xdf\x5c\x46\x1c\x33\x7c\x26\xbb\x67\xa7\x17\xa3\xb3\xc9\xfc\xea\xe6\xeb\x78\xf8\x76\x86\x42\x09\x67\x54\x06\x3a\x5d\x08\x7c\xa2\xa1\x8c\x2e\x4f\xcf\x87\xf3\xd9\xf0\x7c\xf8\xc7\x74\x7e\x3d\x3b\x1d\x5f\x5d\x9c\x5e\x8f\x26\xe3\x37\xe3\xdd\xc5\xec\xb9\xc2\x08\xef\x93\xb9\x51\x44\x68\xee\x92\xd6\xf3\xf4\x3f\x3b\xfd\x3e\x1f\x0c\x6f\x47\x67\xc3\xab\x37\x93\x40\x91\xbb\x79\x88\xb6\xca\xd5\xcf\x74\xd2\x3c\x24\x5e\x4c\xce\xcf\x47\xe3\xf3\x37\x0f\x8b\x5c\x46\x11\x13\xcd\x25\xfb\x5a\xfc\xf4\x66\x7e\x39\x19\xbc\xa1\x87\xd2\x24\x0d\x62\x19\x3e\xd5\x45\x6d\x3a\x74\x26\x32\x99\x58\x95\xcf\xde\x8c\xdf\xbc\xa0\x9b\x2b\x29\xcd\xbc\x5e\xf7\x3d\x41\xcf\xde\x51\x2b\x1e\x7a\xd5\x25\x44\x1f\x7a\x68\x68\x51\x6b\xe4\x05\x51\xd1\x0c\xd0\x56\x23\x50\xd6\x61\xbe\x80\xda\xbb\x88\x3e\x84\x91\x00\x4a\x34\xc2\x9d\xed\x23\xfe\x44\x6a\x80\x4b\x4a\x78\x59\xbb\x3b\x04\x3b\x7b\x47\x84\xb1\x0d\x83\x6d\x4a\x99\x01\x21\x0d\xc8\xe5\x92\x51\x46\x38\xdf\x00\xc9\x08\xe3\xae\x71\x95\x02\x5f\xa0\x46\xcf\x05\xd9\xa7\x3c\xaf\xd6\x68\x56\x67\x45\x11\xf9\x17\xc6\x69\xab\x48\xab\x0d\xd6\x69\xf5\x46\xf7\x96\xba\x47\x23\x25\xd3\xa4\x45\xd8\x18\xae\x93\xda\xb2\x30\x96\x61\xca\x6b\x91\xc3\x13\xb6\xc7\x15\x92\x70\x22\xf8\xa6\x65\x28\x55\x48\xdb\xbe\xb7\xb0\x1a\x83\x7b\x01\xbd\x76\x7f\xd1\xee\x62\x7e\xac\x6c\xee\xa6\x6e\x1a\x36\x6c\x31\xf8\x36\xb5\x6d\x5d\x1e\xa1\x0e\x6c\x4f\x83\x46\x57\xdc\xc2\x76\xaa\x5c\x46\xae\x07\x66\x65\x77\xbb\x42\x85\xb0\x40\x4a\xdc\xcd\x8c\x59\xa1\xba\x63\x1a\xcb\x8e\xd7\xa9\x2a\x51\x32\x4c\x29\x02\x2a\x25\x55\x15\x92\xb3\x35\x82\x59\xb1\x8a\xf1\x1e\xc2\x4d\x7e\xdb\x23\x6d\x13\x1c\xe4\xd7\x32\x74\x45\x54\x88\x19\x2c\x19\x47\x78\xef\x75\x20\xa3\x5e\x16\xeb\x1e\x59\x86\x9f\x7f\x59\x2c\x16\xc1\x17\xfc\xf5\x73\x70\x72\x82\x9f\x83\x5f\x7f\xf9\xf7\x49\x70\xfc\xe9\x5f\x9f\x8e\x09\x3d\x3e\x3e\x3e\xfe\xd4\xa3\x4c\x29\xa9\x83\x2c\x9e\x1f\x1f\x71\x19\xbd\xef\xc3\x58\x82\x4e\xe9\xca\x23\x4a\x55\x76\xee\x9b\x76\x53\x15\xeb\x60\x7b\x37\x57\x61\xa5\xdd\x03\xe6\xca\x7c\x9c\xba\x7d\x68\x4f\xe9\xca\x9e\xd3\x57\x59\x4f\x61\x02\xb5\x9e\x2a\xb9\xc0\x2a\x09\xde\x3f\xdc\x13\xfa\xbf\x56\x38\xf2\x2c\xf6\x16\x4c\xf4\x2a\xe1\xc8\x8f\x06\xb4\x31\xa0\x25\x25\x06\x02\xb8\x19\x8f\xfe\xe8\x37\x0d\xb0\x57\x35\xb8\x40\x49\xf8\x8f\x95\xac\x27\x52\xce\x1b\x81\xbc\xf3\xb6\xe3\x67\x0f\xe4\xaf\x1d\xa1\xdf\x3e\x94\x1d\xfa\x40\xec\xae\xc1\xaa\x51\x1e\x88\xc2\xf2\xea\x11\x16\x1b\xd0\x69\x82\x2a\x66\xdb\x82\xe0\xcf\x96\x20\xde\xee\x16\xa4\xc0\xdd\x7a\x34\x3f\x55\xe0\xaf\xa3\xa4\xda\xf1\x60\x43\x84\xbb\xcd\x53\x02\x0d\xea\xf2\x62\x2f\xbf\xd1\xeb\x79\xb3\xef\xd9\x65\xad\x8d\xf6\xb8\x35\xec\x96\x3b\xdf\xa4\x97\xc8\xb0\x6d\x32\x16\xd5\x4e\x74\xde\x3e\xee\xa3\xe9\xe7\xc7\xfa\xa6\x2f\x37\x2a\xd4\x26\xa7\x6e\x38\xb0\xbf\x83\x4a\x4f\xd8\x4e\x1e\x4e\x9a\xc7\x79\xa9\x69\xe3\xb0\x48\xcb\x4b\x97\xd1\x48\x24\xa4\x36\x8c\x42\x92\xaa\x44\x6a\x7c\x8d\x0c\x25\xd0\xec\xbc\xf3\x2d\xec\xce\xad\xfb\x81\x93\x69\x15\xa1\x8f\x17\xaa\x3f\x77\x5a\x8c\x54\x42\xe7\x2b\x24\xdc\xac\xe6\x89\xdd\x0c\x02\x12\x86\x2a\x4f\x93\x56\x65\xb9\x21\x55\xef\x97\x2b\x76\xba\x77\x22\x7c\x7e\xcb\x91\xc5\xfa\xa9\xed\xc6\x4b\x04\xc3\xdf\xa5\x36\xd7\xf2\xac\xf1\x4d\xef\xc7\xc3\xe1\xcb\xb8\xf8\xcb\x86\xa3\xed\xb2\x3e\x2d\x21\x6d\x4b\x9c\xbb\x53\xae\x3f\xd1\xca\xc7\x33\x8b\x5a\xa9\xee\x6d\x18\x59\x49\x6d\x40\x91\x3b\xf0\x17\x41\x40\x28\x45\xad\x4b\x7b\x74\x9f\x5a\x2d\x7e\xd5\xbb\xda\x1c\x36\xa5\xd9\x49\xd8\xed\xce\x1d\x71\x60\x27\x4a\x57\x85\xd1\xa5\xa6\x9d\x20\xb5\xf2\xa1\x55\x51\xec\x24\xad\x56\x4d\xcd\x3a\xea\x10\xae\x27\x83\x49\x1f\x42\x57\xaf\xd9\xe6\x86\xca\x10\xf3\x2f\x4f\xe0\x73\xb0\xab\x56\xad\x95\xb8\x26\xeb\x81\x70\xc5\xb4\xaf\xdb\xf2\x6a\x0b\xce\x66\x23\xdb\x63\xdd\x6f\x80\x09\x6d\x08\xf7\x19\xc5\x16\xb4\xd5\x0d\x99\xf0\x47\xe9\x0b\xbd\xf2\x63\xf6\xd1\x3e\xa2\xec\xfa\xe0\xb5\xe5\x9b\xd9\xa3\x78\x5d\x51\xa2\x2b\x46\xec\x05\xd4\x74\xf6\xae\x10\xf0\x38\x50\x25\x2a\x34\x3f\xe2\xed\x24\xfe\x81\xaa\x68\xcf\x9a\x68\x2f\x25\x74\x46\xa4\xad\xf1\x68\x1f\xc8\xe6\xc1\xd4\xbe\x1d\xee\xa3\xcf\xb2\x18\xaa\xc6\xd3\xae\x38\xbc\x17\xd8\xce\x53\x7e\x0a\x58\x57\x21\xbc\xab\x0c\xde\x8b\xbb\x0e\xb5\x37\x6a\xb8\xbd\xf8\xaa\x17\x4a\xdd\x45\xd6\x4e\xa0\xad\xfd\x64\xab\x9b\x0c\x1e\xee\x81\xfb\xdb\x32\x75\xe0\xeb\xd5\xce\x52\x75\x77\x41\xdb\x7c\x5d\xa5\x16\x84\x1e\x91\xd4\xac\xa4\x62\xff\x73\x6b\x8e\xd6\x5f\xf4\x11\x93\xbd\xec\x64\x81\x86\x14\xef\xae\xf2\x87\x47\x33\xc9\xf1\x2b\x13\x21\x13\xd1\x8e\x07\x58\x4a\x72\xcc\x2f\xb0\x49\xc2\xce\x6d\x6e\xd8\xb1\xd3\x01\x40\x6b\x8f\x16\xa4\x4e\x17\x7f\x22\xb5\x25\x4e\x90\xaf\xbe\xaa\xbd\xf4\xd9\xff\x11\x98\xd5\x40\x7b\xbf\xa7\xe9\xe4\x19\x6f\xcf\x94\x4d\x6e\x76\x7d\x50\xea\x24\x4f\xf1\x01\xbc\x7b\xe7\x7e\x28\xd4\x32\x55\x14\xcb\xf1\xf2\xd5\x93\xce\x07\xdc\xdb\x24\xf7\x3b\x43\xb5\x78\x58\xe7\xee\xe3\xf2\xff\x44\x68\x5e\xe2\x94\x3b\x64\x2c\xd9\x09\x6c\x41\x8e\xaa\x90\xa9\x21\x51\x2e\x4f\x4d\x9a\x86\x2c\x25\xf7\x9e\x5d\xfb\x2f\x67\xda\xff\xb8\x23\x86\xae\x5e\x49\x82\xc2\x7d\x52\x8d\xca\xce\xfc\xb0\x20\x81\xed\x67\x94\x8f\x49\x0d\xa1\x5e\xd5\xd3\x8a\x2c\x66\x0d\x22\x58\xe4\xcb\x5e\xd0\xed\x5a\x47\x5d\xf5\xbf\xa7\x80\x9f\xe7\x85\xa1\x87\xf5\xbe\xd0\xf7\x66\xfc\xba\xa1\x28\x7e\x38\xe4\x57\xd0\xcf\x36\x43\xfa\x9b\x84\xa9\x80\xaa\x70\xbb\xd1\x93\x84\xe1\xbd\x41\xe1\xde\x12\xe6\x98\x5d\x8e\x90\x6a\x23\xe3\x62\x30\x44\xf7\xe8\x31\x4f\x45\x15\x5f\xc8\x83\x53\x7b\x9b\xa2\x93\x5e\x7f\xd1\x1d\xe8\xf9\xac\xcb\x63\x31\x49\x12\x26\x22\x5d\x9d\x28\x2d\xb4\x98\xa9\x6c\x59\xc6\x92\x57\xf7\xc3\x9a\x3e\x5f\xde\xbc\x2c\xec\xcb\x9a\x54\xe3\x71\x55\x27\xe0\x33\xb2\xdb\xff\x03\x00\x00\xff\xff\x3a\x1f\x89\x21\xb5\x2d\x00\x00")

func deployDataVirtletDsYamlBytes() ([]byte, error) {
	return bindataRead(
		_deployDataVirtletDsYaml,
		"deploy/data/virtlet-ds.yaml",
	)
}

func deployDataVirtletDsYaml() (*asset, error) {
	bytes, err := deployDataVirtletDsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "deploy/data/virtlet-ds.yaml", size: 11701, mode: os.FileMode(420), modTime: time.Unix(1522279343, 0)}
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
	"deploy/data/virtlet-ds.yaml": deployDataVirtletDsYaml,
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
	"deploy": &bintree{nil, map[string]*bintree{
		"data": &bintree{nil, map[string]*bintree{
			"virtlet-ds.yaml": &bintree{deployDataVirtletDsYaml, map[string]*bintree{}},
		}},
	}},
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
