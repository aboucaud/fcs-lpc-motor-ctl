// Code generated by go-bindata.
// sources:
// root-fs/bower.json
// root-fs/favicon.ico
// root-fs/fcs-lpc-motor.html
// root-fs/index.html
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _bowerJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x5f\x6f\xd3\x30\x10\x7f\xef\xa7\xa8\xca\x5b\x45\x9c\x76\x2b\x43\x4c\x08\xa1\x09\xde\x91\xc6\xdb\xc4\x26\xc7\xb9\x26\x46\xb6\xcf\xd8\x17\x6d\xd5\xb4\xef\x8e\x1d\x37\x6d\x52\x92\x96\xa7\x36\xfe\xfd\xf1\xcf\x77\x67\xbf\xce\xe6\xf3\x85\xe1\x1a\x16\xb7\xf3\xc5\x56\xf8\x4c\x59\x91\x69\x24\x74\x8b\xf7\x11\xaa\x51\x83\xe5\x55\x0b\xd7\x44\xd6\xdf\xe6\x79\x25\xa9\x6e\x0a\x26\x50\xe7\x15\x66\xca\x7b\xca\x47\x94\xbc\xa1\x1a\x9d\x0f\xc2\x87\xf0\x19\x16\xee\xa1\xe0\x9e\x24\x98\xf9\x9d\x34\x40\xf3\xcf\x45\xfc\xf9\x2a\xc0\x19\x26\xea\x2f\x8b\xc0\xfa\xd5\x2a\x4b\xf0\xc2\x49\x4b\x12\x4d\xdc\x36\xd9\x69\x2e\xfb\x5f\x58\x36\x0a\x7e\xee\x6c\x0c\xf6\x90\x64\x4a\x0a\x30\xbe\x4d\x7a\x77\xff\x2d\xf1\x64\x65\xd0\xc1\x31\xc3\x72\x99\xb3\x65\x0b\xc5\x63\x63\x09\x4f\xc9\xc9\x77\x6b\x05\x3e\x83\x7b\x0a\x47\xb3\x68\xc0\xd0\x61\x9d\xc0\x53\xff\xbf\xef\xc7\xb5\x60\x4a\x30\x42\x42\x3c\xed\x6b\x22\x59\x6e\xc1\x65\x1a\x4c\x13\x03\xfd\x40\xb5\xd3\xe0\xbe\x2b\xd0\xd1\x35\x3f\xa2\xef\x1e\xd7\xec\x8a\x5d\x75\xde\x09\x28\x1d\x0f\x31\x32\xcb\x0d\xa8\x69\x79\x9f\x15\x6d\x56\xec\x66\x68\x53\x03\x2f\x2f\xdb\xf4\x59\xd1\x66\xcd\xae\x87\x36\x5e\xc9\x40\x98\x36\x48\x78\x4a\xf0\x69\x28\x25\x44\x55\xf0\x33\xda\x3d\x21\xed\xbb\xe9\xc4\x0e\x8b\x30\x48\xd9\x16\x0d\x6d\xb9\x68\x5b\xfa\xb8\x62\x9b\xd3\x60\x45\x43\x94\x66\x64\xdc\x3c\xe1\x29\xd8\x7a\x3d\xd4\x4a\x02\x3d\xad\x8c\xe8\x78\x2d\x84\x43\xa5\xfe\xb3\xb2\x23\xe4\x7d\x98\x0f\x27\x61\x04\x9a\x8b\xa7\xe9\x91\x46\xbb\xad\x39\x81\x93\xfc\x4c\x9e\x8e\x31\x2a\x17\xdc\x95\xd3\xd2\x88\xa6\x82\xac\x4f\x3b\x5c\x55\x0a\x2e\x86\x1f\xd0\xf6\x45\xd8\x9c\x38\xf1\xc2\x9f\x31\x08\x68\xd4\x5d\xb3\x43\xed\xa4\x0b\x05\x89\x8f\xd3\xa8\xec\x88\xa6\xed\x3e\x0e\x64\x5b\x74\xa3\xed\x3f\x80\xfb\x8c\x37\x03\x55\x78\xd3\xb0\x72\xf8\x9c\x11\xbc\x10\x77\xc0\x27\x2d\xfe\x61\xee\xfd\x56\x9d\xdf\xef\x3f\x0d\xb8\x5d\x3b\xd9\xe1\xfa\xb3\xd5\x30\x9d\x82\x97\x4c\xf1\x1d\x36\x34\x1d\xf2\xc8\x49\x6f\xc8\xc9\xa4\x96\xa1\xd1\x58\x9d\x79\x3d\x5a\x3c\xc5\xda\x0c\x76\x8f\x83\x36\x5d\xd2\x16\xed\xee\x46\x50\xbd\xcd\xde\x66\x7f\x03\x00\x00\xff\xff\xd7\x73\x78\xe9\x46\x06\x00\x00")

func bowerJsonBytes() ([]byte, error) {
	return bindataRead(
		_bowerJson,
		"bower.json",
	)
}

func bowerJson() (*asset, error) {
	bytes, err := bowerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bower.json", size: 1606, mode: os.FileMode(420), modTime: time.Unix(1465892593, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _faviconIco = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x94\x7d\x4c\x55\x65\x1c\xc7\x9f\xab\xe6\xda\x5a\x41\xb9\xb5\x85\x93\xdd\xd6\x8b\x28\xe5\xac\x10\x73\x72\x31\x5e\xe6\xc2\xd8\xa5\x49\x2f\xda\x35\xbb\xad\xc1\x6e\xbc\x94\xcb\x90\xba\x8b\x09\x82\x12\x6f\x37\x27\x83\x08\x0a\x48\x97\x97\x80\x41\xca\xa0\x04\xe2\xc2\x15\x24\x0b\x3d\xbc\x5c\xf4\xc2\xbd\x89\x88\xbc\xc4\x25\x2e\x01\x5a\x9b\x7f\x7c\x3a\xf7\xdc\x51\xba\x66\xbf\xb3\xef\xce\x73\x7e\xe7\xfb\x79\x9e\xf3\xfc\x9e\xe7\x39\x42\xa8\xe4\xcb\xd7\xd7\x73\x57\x0b\xc3\x0a\x21\x1e\x16\x42\x04\xc8\x92\x53\x72\xc6\x9b\x5f\x8a\x55\xf7\x79\x75\xb7\xe8\xf8\x32\x5e\xdf\x64\x8a\xb1\x34\x14\x44\xd2\x60\x7a\x81\xea\xc3\xc1\x14\x25\xf8\x91\x1e\x2b\x2c\x29\x31\x42\x7f\x37\x6e\xe4\xdc\x71\xf5\x40\x63\xae\x64\x6f\x2f\x67\x66\xb4\x97\xb9\xdf\x1c\xcc\x4e\xd8\x98\x19\xeb\x65\x7a\x54\xc2\x71\xa1\x91\xca\x2c\x2d\x89\x51\x42\xd2\x47\xdc\xf6\x41\x72\xcc\x8f\x9d\x57\xbb\xec\x2d\xee\xb9\x49\x27\x8b\xf3\x73\xcc\x4e\x5e\xc1\x3d\x7d\x8d\xe6\xd2\x5d\x54\x1d\x0c\x60\xe2\xca\x45\xfe\x98\x9d\xc6\xed\x9a\xa0\xa7\xa3\x8a\xf8\xe8\x07\xdc\xb1\x5b\xff\xed\xe3\xaf\x99\x41\xe9\xd6\x0d\x17\x37\x17\x17\xb0\xb7\x95\xd1\x59\x99\x44\x73\xf1\x1b\xb8\xc6\x87\x39\x9e\xea\x47\x4d\x76\x28\x8b\x0b\x0b\x98\x12\xd7\xe2\xb4\x59\xf9\xa9\xad\x0a\xed\x66\x21\x79\xd8\x89\xde\x7a\xfd\xfc\x78\x1f\xb7\xfe\x9c\x67\xdc\xd6\xc2\xe5\xb6\x52\x86\x3a\x4f\x32\xd4\x5d\xab\x30\xb5\x47\x82\xa9\xce\xd6\x28\xed\x8f\x5f\x16\x64\xc6\xad\x65\xa8\xdf\x4a\xe1\x11\x03\x21\xeb\x84\xde\xde\x5a\x68\xb9\x31\x3b\xc6\xef\x23\x3d\x4c\x0e\x9f\x57\x7c\x1e\xb9\xc6\x1d\xf4\xb6\x16\x52\x9a\x74\x2f\xdf\x1e\xf6\xf2\xa9\x31\x82\x84\x17\x05\x15\x26\x03\x03\x17\xad\x04\x3f\x2e\x2c\x7d\x8d\xf9\xdc\x9c\x9b\xc4\xf6\xc3\x51\xa6\xaf\xda\x18\xe9\x6d\xa6\xc9\xa4\xa5\x26\xe3\x59\x2e\x75\x7e\x4d\xd9\xbe\x55\x9c\x3c\xe4\xe5\x3f\xd4\x0a\xf2\xf7\x6b\x38\x51\x94\x82\xd3\x2e\xf1\x5a\x54\x10\x3f\xd7\xa4\xb1\x30\x73\x0d\x6b\x45\x22\xce\x5f\x4e\x33\x25\xf7\xd1\xdf\x5a\x42\xbf\xa5\x44\x61\x06\xbb\xcc\xf4\x9c\xf1\xb6\x4f\x57\x1a\x19\x1e\xe8\x22\x73\x7f\x34\x99\x07\x74\x04\x3d\x71\x3f\xed\xe5\x71\xf2\x5a\x39\x69\xfc\x4c\x4b\x5f\x4b\xb1\xe2\xb3\x9d\xfd\x8a\x73\xf5\x69\x58\xab\x8d\xca\xb3\xd4\x61\xa6\x28\x55\x43\xf9\xa7\x3a\xec\xfd\x5d\xbc\x1e\xee\xc3\xf3\x4f\x0a\xd6\xf9\xa9\x38\x95\xb7\x55\x5e\xe7\x41\xea\xb2\x37\xf1\x7d\x71\x8c\xe2\xaf\xcf\x0f\xe7\xd8\x3b\x2a\x72\x75\x02\x6b\x5d\x8e\x52\xb7\xf7\x5f\x12\x1c\xd0\xf9\x2b\xfc\x7b\x6f\x69\xd8\xb0\x46\xb0\xf1\xd1\x65\x7c\xfe\xee\x4a\xcb\x84\xe3\x2c\x56\x73\x12\x75\x79\x61\x0a\x6f\xce\xd2\x90\xb3\x5b\x70\xe8\x15\x41\x6e\x9c\x3f\x3d\xed\xe6\x7f\xea\xea\x51\xcc\x36\x7f\x02\x1e\x11\x6c\x59\xbf\xdc\x52\xf0\xa6\xd0\x37\x97\x1b\xb8\x3e\x64\xa5\x36\x77\x87\xf2\xbe\xf2\x13\x0d\x07\x77\x0a\x3e\x92\xc7\x6d\x90\xe7\xec\x9a\xba\x8e\xf9\x0b\x23\x19\x1f\x44\xf3\x6a\xa4\x3f\x81\xab\x05\xcf\x3c\xb6\x8c\xf0\xe7\xee\x51\xf6\x73\xce\x5e\x5f\xc9\xd6\x59\xc5\xb8\xf3\x82\xc2\x57\x64\x46\x2b\xb5\x4e\xde\x21\xf8\xe6\x58\x02\x9d\x2d\x66\x74\x91\x3e\x68\x43\x7c\xe8\x6a\x6f\xe2\xa9\x35\x2a\xb6\x6d\x5c\x21\x2d\xed\x3f\x79\x1c\x75\x41\xf2\x06\xf7\xaf\x97\xba\x99\x99\x1a\xc5\x39\xd8\x4d\x7a\x5c\x20\x6f\x47\x08\x76\x87\x0a\xf2\xd2\x74\x94\x98\x8c\x1c\xcd\x36\x92\x9a\xbc\x87\x90\xa7\x57\xb8\xb7\x6f\x5e\x79\xc7\x19\x48\x8a\x12\xea\x7d\x3b\x7d\xa5\x53\x27\x32\x18\x75\x4a\x5c\x75\x48\x74\xfd\x58\xc5\x99\xef\xca\x70\x5c\x96\x88\x08\xf2\x63\xfd\x6a\x95\x67\xce\x52\xe4\xa6\x3b\xd9\xdb\x63\x57\xa8\xd0\xef\xd9\xfe\xa0\x25\x25\x3e\x8c\xbc\x74\x03\x59\x46\x03\x7b\x63\xc3\xd8\x12\xf8\x90\x25\x78\xed\xf2\xff\x9c\xdf\x39\x1f\xb5\xba\x7d\xb9\x5a\x9d\xae\xf2\x4a\xfe\x63\xfc\xaf\x96\x7c\x1e\xc6\xc3\xfe\x1d\x00\x00\xff\xff\xb8\x6d\x5d\xdb\x7e\x04\x00\x00")

func faviconIcoBytes() ([]byte, error) {
	return bindataRead(
		_faviconIco,
		"favicon.ico",
	)
}

func faviconIco() (*asset, error) {
	bytes, err := faviconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "favicon.ico", size: 1150, mode: os.FileMode(420), modTime: time.Unix(1465892593, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fcsLpcMotorHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3c\x6d\x73\xdb\x36\x93\x9f\xa5\x5f\x81\x30\x77\x53\xa7\x09\x29\xd9\xb1\x7b\xa9\x22\xe9\x2e\x75\xdd\x69\x6f\xf2\x36\xb1\xef\xfa\xb4\x99\x4c\x86\x22\x21\x89\x31\x45\x30\x24\x65\x5b\xf1\xf8\x3f\x3d\xbf\xe1\xf9\x65\xcf\x2e\x16\x20\x41\x9a\xd4\x5b\xe4\x64\x32\x6d\x66\x62\x81\x78\xd9\x5d\x2c\x16\xfb\x46\x80\xfd\x7b\x3f\xbf\x3a\x3e\xfb\xe3\xf5\x09\x9b\x66\xb3\x70\xd8\xee\xdf\xb3\x6d\x76\x2c\xe2\x45\x12\x4c\xa6\x19\x3b\xe8\xee\xff\xc0\xce\xa6\x9c\x4d\x84\x1d\xa6\x69\xc6\x9e\xcd\xb3\xa9\x48\x52\x87\x3d\x0b\x43\x26\xfb\xa4\x2c\xe1\x29\x4f\x2e\xb8\xef\xb4\x5b\x30\xf8\xff\x52\xce\xc4\x98\x65\xd3\x20\x65\xa9\x98\x27\x1e\x67\x9e\xf0\x39\x83\xc7\x89\xb8\xe0\x49\xc4\x7d\x36\x5a\x30\x97\xfd\x74\xfa\xb3\x9d\x66\x8b\x90\xcb\x61\x61\xe0\xf1\x08\x86\x66\x53\x37\x63\x9e\x1b\xb1\x11\x67\x63\x31\x8f\x7c\x16\x44\x50\xc9\xd9\xf3\xdf\x8e\x4f\x5e\x9e\x9e\xb0\x71\x10\x72\xa7\x6d\xdb\xc3\x76\xbb\x1f\x06\xd1\x39\xa0\x0f\x07\x56\x30\x8b\x45\x92\x59\x6c\x9a\xf0\xf1\xc0\x1a\x89\x4b\x9e\xbc\xf7\x04\x54\x46\x3c\xca\xd2\x4e\x2c\xc2\xc5\x8c\x27\xfa\xd7\xc1\xc9\x5a\x08\x21\xf5\x92\x20\xce\x58\x9a\x78\x35\xa3\x3e\x7c\x9c\xf3\x64\xd1\xf1\x83\x34\x53\x65\xe7\x43\x6a\x0d\xfb\x1d\x1a\xb5\x11\x05\x41\x22\x22\x7b\x1c\xf2\x2b\x3b\x74\x17\x62\x9e\xdd\xaa\xb0\xbd\xd0\x4d\x53\x9e\x16\xb4\x6d\x06\xdb\x9d\x67\x62\x92\x88\x4b\x3b\xe3\x57\x99\x9b\x70\xb7\xa1\x5a\xc3\xdf\x94\x74\x91\xcc\x8a\xd2\x96\x40\x02\x4f\x44\x66\x71\x4b\x30\xb1\x3b\xe1\x66\x71\x0b\x96\xc5\x6e\xcc\x13\x7b\x34\xcf\x32\x11\x95\x1e\x36\x27\x89\x46\x7b\x6e\xe2\x1b\xc5\x6d\xc1\xf8\x81\x1b\x8a\x49\xe9\x61\x6b\x50\x89\x0b\xf5\xc0\xa1\x88\x87\x35\x55\xdb\x82\xc5\x85\x2b\x33\xce\xa8\xd9\x1a\x68\x14\xc3\x96\x30\xca\xbb\x00\xb4\xbd\xc4\x2b\x58\x19\x9f\x19\xc5\x6d\xc1\xcc\xdc\x8c\x27\xb0\x92\x95\xc7\x6d\xc1\x81\xee\x11\x61\x68\x4f\xb9\xeb\x57\x56\xb7\xa6\x65\x6b\x24\x61\xe0\xa3\xba\x34\x1e\xb6\x05\x95\xb9\x23\xb3\xb8\x35\x18\x31\x99\x84\xbc\x2c\x78\xa5\xba\xed\x01\x8b\x70\xe4\x26\xe5\xa7\x42\xa3\xf8\x62\x66\xcf\x84\x3f\x0f\xc1\x82\xf9\x03\x6b\xec\xa5\x76\x18\x7b\x50\x95\x89\xc4\x86\xdd\x19\x44\xd0\xad\xd5\x07\xf9\x88\x43\x58\x59\x28\xb7\xfa\xd2\xaa\x0d\xdb\xbd\xa9\x00\x8b\x79\xdd\x6e\x81\x01\x81\xc6\x45\x8f\x8d\x42\xe1\x9d\x3f\x6d\xb7\x46\xe2\xca\x4e\x83\x4f\x41\x34\x81\x3a\x91\xe0\x72\x41\x15\x34\xa0\xc8\xda\x6e\x18\x4c\xa2\x1e\x03\x7b\x08\x92\x02\x95\x33\x37\x01\x34\x3d\x76\x14\x5f\xc9\xa7\x2b\xfb\x32\xf0\xb3\x69\x8f\x1d\x1c\x75\xa9\x2a\x88\xf2\xaa\xae\xac\xba\x01\x1c\xc4\x95\x48\x90\x91\x45\x3a\x5a\xf6\x25\x1f\x9d\x07\x80\x21\x8e\xb9\x9b\xb8\x91\xc7\x7b\x2c\x02\x6e\x00\x8c\x16\xd1\xd1\x63\x5d\x7c\x88\x5d\xdf\x97\xd4\xc9\xa7\x91\xeb\x9d\x83\x0d\x01\x63\xdc\x63\x19\x0c\x4b\x63\xd8\x56\x51\x06\x4d\x37\x38\xdd\x8e\x9a\x2f\x96\xd1\x38\xb0\x19\x07\x27\x01\x98\x15\xc3\xfc\x2d\xe6\x7a\x59\x20\xa2\x81\xd5\xc9\xb9\x05\xfd\x82\xd9\x84\xb9\x17\x2e\xec\x4f\xb2\xbd\xd3\x2c\x8b\x7b\x9d\xce\xe5\xe5\xa5\x83\x7e\x86\x23\x92\x49\x27\x85\x5d\x97\x76\x7c\x3e\x76\xe7\x61\xd6\x41\xbb\x9f\x22\x0c\xf1\x7e\xdf\x89\xa3\x09\x00\x0e\xb3\x81\x45\x40\xac\x0e\x81\x35\x94\x00\x8b\xdc\x19\x1f\x58\x73\xf0\x4d\xb0\x64\xb1\xd0\x1d\xa1\x50\x80\x7f\x92\x58\xec\xc2\x0d\xe7\xd0\x7a\x11\x00\x12\x01\xcf\x09\xff\x38\x0f\x12\x70\x4d\xc0\xba\x1b\x40\x9a\xa0\xc6\x60\xaa\x2f\x81\x61\x39\xd4\xd7\x79\x45\xb6\x88\x4b\x1d\x34\xe4\x3a\xc0\xb4\x46\x4c\x5a\xfe\x81\xa5\x97\x4a\xc3\x48\xe7\xa3\x59\x90\x81\xbf\x61\xda\x27\x96\xb8\x41\x0a\xd0\x4e\x65\x63\xbf\x64\xbb\x00\x85\x2a\xc8\x65\xc1\xb5\x40\xd9\xec\x14\xc2\x09\x4f\xda\x77\x69\xb5\x5e\x93\x33\xb4\x87\x92\xd1\x0a\xd2\x5e\xad\x70\x3f\xc2\x46\x68\x78\x4f\xd4\xbc\x97\xb5\x3d\x36\x9e\x47\x72\x59\xf7\xf8\x05\x48\xc2\x03\x29\x5d\xad\x16\xda\x74\x01\xee\x19\x74\xda\x43\x68\xef\xa7\x6e\xe4\x87\x5c\x0f\xb2\x1e\x52\xef\xa7\xb2\xf3\x05\xac\xbd\x14\x97\x01\xf3\x85\x37\x9f\x41\x8b\x33\xe1\xd9\x49\xc8\xb1\xf8\xd3\xe2\x37\x1f\x60\x40\xbb\xa5\xfa\x97\x81\x43\x03\x02\xc4\xdf\x87\x16\x03\x07\x92\x84\xae\xc7\xde\x52\xa5\x43\xcf\x0f\xad\x77\xd8\xa8\x64\x50\xb7\xd1\x23\xb4\x69\xd8\xc6\x08\x20\x87\x04\xd7\x68\xa1\xfe\xd8\xa2\x84\xb8\x86\x22\x40\x33\x1c\xee\x80\x0a\xe2\xf3\x9e\xaa\xea\x74\x12\x9e\xcd\x93\x88\x01\x37\x1d\xe2\xe6\x73\xa4\x60\xcf\xe0\xe4\x8d\x5c\x23\x2d\x71\xff\x0f\xea\xc3\x77\x11\xf2\x6b\x37\x03\x05\x82\x7c\x77\xbe\xff\x6f\xab\xd4\xe9\x24\x49\x44\xf2\x82\xa7\x29\xb8\x4f\xd0\x3e\x12\x42\xc8\x0e\x37\x08\xb1\xf0\x6f\xfb\x9d\x42\xfd\xad\xa3\x0c\xc1\x8d\xfd\x8b\x6b\xc3\x92\x2b\x89\xfa\x04\x11\x4d\x39\x06\x4a\x3d\x76\xb8\xef\x3c\xfe\x2f\xa2\xbb\xa5\x28\x64\x4f\x88\x68\x03\xaa\xed\x89\x50\x00\x6a\xd8\x1f\x7b\xb6\x4d\x00\x01\x92\x7d\xd4\xed\xca\xf5\x56\xcd\x97\x53\x50\x8f\x12\xe9\x86\x3a\x58\x2d\xd2\x67\xa9\x1f\x3d\x06\xe8\xb2\x86\xcf\x25\xcc\x3b\x56\x45\x48\xf5\x6a\x09\x5d\x21\xa2\xb0\xd2\x52\x3e\x49\x06\x02\x98\x80\x37\x4f\x33\x18\x41\x13\x37\x84\x54\x2d\xcf\x7e\xb7\xfb\x9f\x80\x4f\x2f\xa0\x7a\xd4\x22\xd8\x65\x18\x5d\x2d\x93\xe0\xff\x01\x8c\xe1\x02\x96\x51\xc5\x78\x10\x3b\x07\x9f\x44\x94\xb9\xe1\x03\x14\x4b\x29\x68\xfe\x82\xb4\x67\x2e\x65\x87\x4a\x24\x64\x7b\xa3\x6b\xa7\x06\x09\x30\x60\xb0\xb4\x3d\x70\xad\x44\x38\x97\x12\xd1\x6a\x65\x22\xee\x49\x59\x6d\xc9\x10\x5d\x95\x47\x02\x16\x64\xa6\xa4\xb8\x15\xf2\x71\xa6\xcb\xcb\x85\x6f\x92\xf0\x85\x0d\x5b\xe9\x11\xbb\xcf\x39\x7f\xb0\x06\x69\x0e\x68\xc6\x0c\xf6\x44\x65\x62\x4f\x6e\xcd\x0b\x3d\x42\xea\x74\x9b\x84\x10\x49\x9f\x24\xee\x42\x92\xa8\x89\xc1\x11\x76\xca\x43\x2e\x25\xda\x06\x5f\x4d\xf7\x07\x59\xac\x42\x27\x67\xae\x09\xc1\x28\x74\xa5\xfe\xa9\x1b\xe2\x64\x41\xa6\xd4\x84\xb1\xdc\xe6\x04\x1c\xd8\xf7\x1e\x57\xc0\xab\xeb\x8c\x71\x7d\xc1\x28\x87\x94\x17\x75\xad\x55\x6a\xaa\xdf\xcc\x0d\x22\x73\x75\x31\x5f\x32\x0e\x21\x78\x07\x6d\x49\x7c\xce\xfb\x76\xbe\xc7\x0e\xcc\x81\x1e\x59\xe0\xb9\x21\xf0\x84\x8c\x14\x59\x62\x56\x10\x7d\x70\x18\x5f\x15\x7f\x24\x33\x99\x04\x51\x33\xda\xc6\x85\x03\x22\x34\xb1\x2d\x56\x9d\x99\x1e\xa1\xec\x53\xd1\x4e\x73\xb1\x3f\xc0\x86\x0a\xc6\x01\xf7\x1f\x68\x4c\xf0\xe7\xfb\x4e\xbb\x6d\xea\xa8\xbe\xa9\x06\xb4\x7f\x55\x27\x48\xe3\xe0\xaa\x50\x36\xc8\x54\x4b\x8e\xd0\x43\xd4\x6a\x49\x6d\x96\xbb\x69\x45\x34\xca\xb0\x3c\xb0\xc0\x9d\x98\x5b\xac\x14\x01\x53\xe4\x50\xb8\x64\xc5\x18\x05\xcb\x0f\x2e\x34\x5a\xb5\xce\x24\x10\x34\x4b\xd0\x79\xa7\xa7\x67\xec\x97\xe3\x53\xf6\x02\x15\x4b\x0a\x6a\x28\xb8\x50\x43\xeb\x54\x17\x60\xaa\xad\x96\x73\xe9\x54\x26\xd3\xae\x10\xa0\x36\x93\x95\x2b\x34\x29\x26\xf9\x4a\x59\xa5\xe9\x9b\x21\x3e\x35\x10\x2c\xaa\xcf\x41\xaa\x69\x50\x07\xd9\x63\x88\x8e\x71\x8f\xf5\x61\xba\x51\x8e\x69\x2e\x9d\x65\xb9\x5f\x94\xb2\x7f\xd9\x79\x06\x0b\x09\x7d\x86\xc6\x9c\x01\xc2\x34\xc9\x8b\x2a\xce\x05\xae\xeb\xaa\x7c\x6d\x60\xdd\x87\x92\x63\xf6\x3f\x40\x9b\x3e\x3e\x74\xf6\x8f\x9e\xc0\xff\x23\x67\xff\x87\xde\x51\xf7\x60\xbf\xa6\xee\xa0\x6f\xc4\xe0\x4b\x00\xfe\x59\x07\xf0\x71\xdd\x60\x5d\x47\x14\xea\x4a\xc9\x58\xb9\x98\x60\x0d\xdc\x6c\x9e\xaa\x35\x2b\xd7\xd5\xf4\x76\x43\xd8\x12\xe5\xce\x54\xa5\xb8\x2f\xd9\xc4\x8c\xa5\xc0\xe5\xd3\x0b\x51\xec\x78\xab\xc2\x3f\xa9\x1a\x49\xd1\x71\x58\x8d\xeb\x6b\x5d\xbe\xb9\xb1\x68\x5f\x54\x99\x01\x23\x86\xc7\x62\x36\x03\x2f\x31\xed\x17\x21\x77\x5d\xb7\x93\x2b\x28\x65\x35\x9d\x8c\xaa\xb4\xe0\x4d\x91\x6a\x6b\xa2\xa8\x40\x62\x08\x2e\xa9\x0c\xa6\x55\x86\xda\x3c\x76\x45\xf8\x6a\xc7\x14\x86\xb2\x71\x54\x69\x11\xbc\x2c\x64\xb2\x38\xb0\xae\xac\xf2\x62\x40\xd3\xea\x41\x9f\x96\x0f\xba\x67\xdb\xf4\xc4\x4a\x72\x22\x9d\x10\x30\x6d\x15\x51\xc9\xab\x0b\x08\x98\xaf\xce\xa1\x99\x1b\x07\xdd\xb0\x64\xa7\xac\x98\x89\x08\x63\xda\x26\x76\xa8\xe6\xf5\x06\x57\xd9\xb2\x7c\x30\xf8\xd3\x9e\x3b\x2b\x8f\x50\x75\x0d\x93\xaf\x3c\x95\xf8\x42\x32\x9a\xb2\x57\x2f\x9f\xff\xb1\x4b\xfe\xac\xbf\x6a\x4b\x68\xed\x1b\xe9\xe7\xd2\x46\x27\x5d\x5c\x93\x6f\x55\xba\x5e\xf6\x69\x17\x5d\x6a\xac\xde\x67\x38\xca\xe8\xe0\x52\xa0\x97\x08\xe4\x5e\xc0\xd3\x9e\xb2\xe4\x7a\xb3\xc2\x33\xfa\xf7\x3d\xf6\x72\x3e\x1b\xf1\xe4\x11\xa5\x42\xc0\x51\xa0\x30\x92\xfe\x26\x40\xce\xa2\x29\xca\xc7\xc0\x1d\x0d\xc3\xd2\xc0\x5d\x1b\x0f\xe5\x27\x60\xd9\x09\x22\x30\x57\xbf\x9e\xbd\x78\x0e\x23\x31\x9e\xc5\xca\x22\x82\xad\x78\xf6\xeb\x04\x9f\x25\xdd\x6c\x11\xab\x9a\x5c\xfb\x52\x64\x06\x2e\x2d\x97\xae\xd2\xda\x61\x57\x39\xea\x02\xff\xe9\x36\xbc\xb7\x18\x65\x5d\xf0\x77\xab\xe1\xca\xee\x0a\x72\x0e\x0a\x5f\x1b\xe0\xd0\x8a\xbb\x4c\x8e\x9c\x4d\xee\x7a\xa9\x8a\x7c\xfb\x72\x1d\x7a\xfe\x54\x03\x70\xb3\x29\xc2\x1b\xc3\xde\x80\x6d\x48\x9d\x23\x08\xc4\xdc\x50\x86\xc4\xf5\xfe\x58\x6e\x26\x24\x39\x28\x96\x40\xcb\xc0\x2a\x5c\x1d\x76\x4a\xdc\x66\x20\x4f\x17\x2e\x85\x95\x07\x2a\xa0\x34\x9d\x16\x18\x6e\x6b\xcf\x45\xed\x0a\x30\x2a\x21\xd7\xbb\x25\x43\xd8\xf9\x76\xca\x92\x21\xd4\xc0\x8e\x84\x3f\xf9\x86\x83\xb2\x72\x16\x1a\xea\xff\x34\xeb\xa1\x9c\xe4\x5b\xd1\x84\xde\xcf\x30\xc8\xaa\xa0\x62\x72\xf2\x03\xcb\x70\xc7\x91\xc5\x16\x58\xd0\x28\xba\x85\xae\x1c\x03\x6b\xf1\x23\xb9\xb3\xaf\x6c\x11\x85\xe0\x89\x41\xe0\x2c\xdd\xca\x94\xf9\x41\x8a\x53\xf5\x2b\xd1\xb2\x1c\x73\x19\x64\xde\x94\x49\x09\x50\xde\x54\x25\x6e\xde\x04\xf5\xa7\xbb\x42\x6d\x30\x73\x25\xc7\x48\x20\x36\xe5\x99\x54\x31\x5f\x85\x65\x77\x86\x79\x39\x9b\x5e\x90\xed\x2c\x8d\x59\x87\x51\xca\xe6\xee\x9e\xe0\x75\x78\x75\x67\xc8\x4d\x01\x5b\x2d\x61\x8b\xc8\xdb\x98\x6f\x29\x0c\xfa\x3a\x4c\xbb\x1b\xcc\x1b\x71\xec\x85\xf0\xf9\x16\x92\xe6\xdf\x81\x12\x5b\x4f\xcc\xee\x02\xf3\x46\x1c\x7b\xf3\xfa\x45\xba\x31\xc7\x92\x78\x06\x96\x50\x91\x99\x93\xad\xfd\x53\x80\x4b\x28\x76\xc2\xa3\x5d\xe1\xda\x88\x2b\xaf\x55\x7e\x71\x63\xce\xb8\xd1\x04\x53\xc9\x5f\x86\x35\x3b\x43\xb6\x11\x6f\xce\xd8\xde\xbf\xfe\x79\xfc\x60\x63\xd6\xa0\xb3\x65\x77\x6f\x93\x9b\xce\xc7\x10\xd7\x0f\xac\x63\xeb\x8e\xd8\x74\x27\x88\xbf\x1c\xcb\xf6\xbf\x16\xcb\x76\x8d\xf8\xcb\xb1\xec\xe0\x6b\xb1\x6c\xd7\x88\xbf\x1c\xcb\x1e\x7f\x2d\x96\xed\x1a\x71\x39\x1e\x2a\x42\x20\x78\x28\xe2\xb0\x3c\xf9\xf4\xdc\x4d\x33\x46\x04\xb1\x79\xec\x43\x2c\xd8\xcb\xdb\xca\x39\x61\x45\x35\x75\x32\x13\xc2\x65\x80\x72\xd0\xd0\xb6\xed\xfa\xd6\x5a\x86\x50\x12\x03\xc2\xf6\xd8\x5a\x65\xfb\x99\x80\x60\xd7\x8d\x65\xcb\x29\x8f\x7c\x95\xf8\x04\x5f\xf1\xec\xd5\xeb\x0a\x5b\x8c\xcc\x4b\x91\x78\xc1\x00\x75\xe3\x4c\x4b\x39\xe1\xd0\x90\x69\x91\x9d\x7a\x10\x27\x27\x10\x3a\x17\x39\x95\x32\xa1\x46\x72\x45\xe7\x55\x30\x25\x92\x42\x8f\x53\x98\xff\xde\x35\x41\xb1\xae\xac\x1b\x7d\x3a\xa0\xae\xf9\x93\x6e\x5e\x2b\x85\x72\xcf\xb6\x1f\xee\xee\x5f\xbb\x85\x59\x01\x62\x09\xc3\x50\x3f\x11\x61\xaa\x4e\xe6\xd6\x24\x6b\x68\x31\x6c\x6f\xe6\x37\x9c\x11\x68\xce\xda\xd0\x09\x34\x64\x93\xce\xa0\x50\x0d\x9d\xa5\x21\xbe\x17\x6f\x6c\x65\x3e\xe4\xe6\x56\x92\x06\x7b\xdd\x7e\x13\x5f\xe9\xf4\x96\xc4\xee\x9d\x23\xb3\x36\xe3\x79\x58\x9f\xcc\xb9\x7f\x78\xf0\xe4\x68\x7c\x58\xe4\x85\xee\x8f\xc7\xe3\xa7\x4b\xe1\x9b\xef\xa5\x8c\x89\x18\xd5\x30\x9b\x73\x7b\x79\x9e\xc9\xd8\x06\x06\x8c\xdb\xc3\xcd\x64\x53\x8c\xf5\xee\x01\xa5\xb1\x2a\x43\xc6\xb0\x02\xf6\x39\x5f\x8c\x04\xa6\x6b\xc6\x02\xb8\x4f\xdc\x5c\x9e\xbf\x2a\x81\x2c\xa5\xc6\xd8\x3d\x3a\x74\xe7\xd2\x81\x88\xa7\xed\x0a\x42\x62\x6f\x1d\xca\x75\x31\x96\x31\x34\xe2\xbe\xd9\x01\xeb\x2b\xe7\x2a\x74\x7a\x57\x16\xae\xaf\xe5\x0b\xc1\x9b\x9b\x22\xed\xbb\x34\x37\x5d\x7e\x4b\xa7\x04\x5a\x6f\x0d\x2f\x0b\xed\x6b\xda\xd2\x37\x37\x50\x82\x3d\x02\xbf\xd4\xc9\x62\x30\x77\x7d\x5e\xed\xfa\x5a\x16\xf0\x15\xd0\x2c\x88\xf0\x19\x7e\xe4\x93\x7b\x25\x9f\xdc\x2b\x7c\xd2\x06\x64\x60\xbd\x7d\xab\xcb\xef\xde\x59\x8c\xfb\x01\xe9\xfe\x7e\xe9\x68\x67\x9d\x4a\xa6\x85\x42\x00\x6a\x47\xbc\x6b\x06\x9b\xbf\xa7\xa4\x2d\xd3\xa8\x9b\x57\xcd\x57\x1d\x38\xa9\x3d\xda\x56\xd1\xe1\x2a\x8d\x6e\x6a\xef\x65\xca\x9b\x55\x35\xd0\x3a\xaa\x5b\xf6\x69\x41\xef\x1e\x63\xa5\x2a\x3d\x77\x3d\xaa\x45\x09\xf5\x9f\x84\x08\xb9\x1b\x51\x9f\x96\x4a\xab\x8f\xdd\x30\xe5\x54\x45\x66\xa0\x45\xfc\xdc\x6e\xac\x6a\x50\xa9\x7b\x22\x19\xdf\xf2\x97\x2a\xdc\xab\x52\xc5\x86\xd6\x47\xf5\xd0\x16\x06\xef\x71\x38\xb2\xf8\x48\x9e\x7d\xa4\x0a\x5c\xb0\x0d\xec\x4e\xb3\x35\xc0\xbb\x23\x4d\x67\xc6\xbe\x21\x7b\xb0\xfc\xa5\xc1\xdf\xc6\xe1\x6f\xe3\xb0\x81\x71\xf8\xf2\x1a\x58\xee\xc2\x5c\xff\xbe\x39\x39\x3d\x39\xbb\x3b\xf5\x4b\x5b\xfe\xdb\x54\xc0\x1b\xea\xd2\x37\x38\xd5\x46\x4d\x6a\x11\x27\x76\xa3\x48\x69\xc8\xb7\xaf\x49\xff\xf6\xac\xff\xe2\xca\x73\x1b\x25\xb9\xb6\xfe\x33\x35\xf3\x5d\xe9\x37\xb5\x13\x97\x2b\xb8\x92\x3a\x93\x24\x55\xea\x94\xd6\x33\xeb\xd6\x54\x3e\x78\x3c\x03\x46\xb3\x81\x56\x70\x0a\xa5\xa1\x7f\x4c\xc5\x67\x81\x29\xb0\x74\x8d\x38\xe7\x78\xb9\x03\xd4\x97\x2c\xaa\x6a\xa9\xad\x18\xd3\x9e\x9f\x52\x88\x45\x56\xc2\x9b\xba\x11\x36\x90\xd2\xdb\xfb\xdf\xd3\x57\x2f\x9d\x54\xd2\x1d\x8c\x17\x7b\xd0\xf0\x60\x7d\x15\xb7\xec\x9c\x07\x52\xfa\x85\x15\xdc\xee\x0f\x65\x54\x44\x7f\xe9\x61\x8b\xc2\x44\xb3\x63\x95\x5f\xf9\xfc\x73\x17\x35\x87\x21\x86\xfd\x5b\xf2\xab\x4f\x80\xe5\x14\x58\x28\x54\xb0\x1c\x01\xa8\x91\xa9\xc0\x0b\x55\x52\x6c\x81\x5a\xa8\x61\xbf\x62\x8d\x3a\x50\x55\x82\x53\xf3\x76\x7e\x6d\x5c\xb1\x80\xe9\xc2\x3f\x85\x08\xdf\x69\x32\xfd\x42\x6a\x25\xb2\x35\x72\x9e\xfa\x58\x5f\xbf\x1c\x1c\x36\x3a\x49\xf2\xd8\x80\xd2\x38\x0d\x24\xab\x4e\x2a\x5c\xef\xaa\x30\xbd\xab\x02\xf4\x87\xfb\x56\x31\x99\xf2\xfa\xd2\xc0\xf2\xa4\x00\x22\xad\xaf\x26\x58\x9f\x22\x5c\x93\xe0\x78\x56\x78\x8c\x8d\x04\x63\xa7\x06\x72\x0f\xc0\x2a\xe4\x04\xe3\xfb\xd1\x06\xfa\xd6\xa4\x47\xbe\xa5\xb3\xf5\x8d\x85\xd5\xa4\x55\xfb\x57\xa8\xb4\x7f\xcc\xd9\xfa\x63\x41\xa6\x96\x0f\xb6\xe7\x73\xcc\x14\xa7\x0f\x1a\x89\xce\xc5\xa0\x38\xd7\xd7\xaf\xfa\xa9\xcb\x1c\x66\xd8\x8a\xd1\x72\xde\x52\xaf\xf2\x99\x6b\x7d\x81\xaf\x44\x94\xec\xa9\x45\x35\xb7\x86\x9f\x9b\xb2\x66\x55\xe5\xb9\x51\xc2\xfa\x33\xce\xe3\xe9\xa3\x94\xab\x2f\x83\x35\xdd\x9c\xb9\x43\x05\xbc\x3b\xa7\xb4\xea\x43\xaa\xfb\x5e\x35\xd7\xbd\x2a\xb7\xbd\xb6\x3f\x75\xd8\xa2\x5b\xa8\x65\xab\x85\x77\x92\x9a\xef\x20\xe1\x28\xe5\x96\xe5\x63\xf3\x3b\x78\x81\x3c\x3e\x66\xeb\xab\x78\xeb\x40\xd1\xd7\xea\xc9\x9b\x6d\xea\x9f\xdf\xa0\xdb\x8f\xaf\x58\x2a\xc0\xe6\xca\xdb\x43\x0a\x52\xfd\xa7\x29\xcc\x69\xa9\xab\x7b\x2b\x28\xaa\xd0\x22\x31\xd6\x20\xdc\xc4\xe6\xa6\xec\x34\x17\xdf\x2d\x6c\xad\xa1\x4b\xa6\x87\xc3\xdf\x24\xc7\x09\x20\xfb\x25\x08\x79\xbf\x03\xb5\x5a\xd7\xe0\x3d\xbd\xfc\xa6\x45\x1c\x0a\xd7\x57\x5b\xc7\xc6\xcb\xce\xf2\x43\x1b\xa5\x2b\x1a\x65\x4d\xc2\x9a\x3c\x64\xec\x4d\x4b\x4d\x17\xf9\x10\x98\x8a\xfc\xa1\x24\x49\xb2\xf4\xed\x68\xc2\x4a\x3d\x80\xcb\x7c\x70\xd0\x05\xcd\x54\x5c\x51\xde\xc8\xe3\xf6\xc2\xc0\x3b\x1f\x58\xef\x57\x26\x72\x5b\xe5\xd3\xd4\xf9\xfd\xc0\x72\x7d\xad\x82\x06\xf6\x15\x77\x1c\xaa\xbc\x04\x82\xf2\x4f\x94\x58\x55\xd6\x7a\x6a\x58\x85\xaf\x15\xa1\x56\xf7\xb9\xa5\xca\x57\x03\xa4\x95\xb1\x41\x50\xd3\xc1\x21\x53\x3f\x6e\x78\xe9\x2e\x00\x14\x40\xce\x6c\x39\x86\xc1\x86\xf2\xf8\x54\x84\x20\x83\x03\x6b\x02\xa6\xa3\xeb\xe0\x55\x14\x2b\x4f\xb6\x6b\x1c\x15\xd4\x2b\xf9\xaa\xd8\x2a\x4f\x4c\xcb\xa9\x68\x06\xa8\x63\xd9\x2b\x18\xbd\x39\x16\x69\x8d\x6e\x21\x91\xc9\x8c\x5a\x1c\xf5\xeb\xb7\x5b\xf3\x55\xd8\x94\x06\x23\xa6\x62\x17\x75\xb8\x5d\x85\x4d\xfa\x70\xbb\x65\x99\xf9\x1b\x25\xa0\x0d\xb1\x13\xec\x79\xbc\x93\x2e\xa3\x9d\xff\x70\xf2\x2d\x53\x84\x3b\xb4\x08\xb4\xa9\x71\x4f\xef\xe1\x08\xd9\x31\x75\xf0\x9e\xd1\x5e\x77\x49\xbc\x53\xf3\x4a\xb6\xf2\x1e\x55\xb2\x9d\x71\xbc\xa1\x9d\x2e\x3b\xd0\xde\xf8\xa2\xd5\xbc\x84\x84\x1f\xc4\x28\xd8\xad\xcd\x2f\xad\x07\x7d\x5e\xc6\xb8\xed\x7a\x24\x2d\x89\x56\xdc\x07\x5a\x8f\xe6\x75\x2b\x8c\xd5\x72\xab\x26\xaf\x73\xaa\xee\x86\x69\xab\x87\x75\xd3\xce\x15\x76\xbf\x44\x6c\x65\x8e\xc4\xa5\x62\x8b\x51\x2f\xe4\xb2\x39\x69\x2d\x61\xa6\x80\x55\xe5\x8b\xb8\x85\x0b\x76\x4b\xb4\xd6\x91\x2c\xf9\x47\x59\x55\x79\xb9\xde\x10\x2e\x31\xfa\xa0\xe4\x0b\xc5\x2b\xe1\x1f\x41\xba\x2a\x31\x32\x74\x81\x5d\xf7\x91\xc4\x46\x0a\x61\x82\xd7\x2b\xb0\x1a\x4a\x79\x6d\x84\x11\x90\x71\xeb\x42\x7e\x93\xea\x54\x5e\xec\x10\xc9\xde\x77\xf7\x6f\x71\xe6\x3b\x82\x88\xe3\x4a\xb7\x2f\xac\xfe\xf4\x60\x28\xe9\x04\x25\x7a\x30\x94\x0a\xd6\x7a\x08\x63\x1e\x5a\xb4\x6d\x51\xf7\xca\xff\xd0\xfa\x86\x03\x9e\x34\x33\x7b\x02\xad\xba\xa7\x55\x60\x00\xbe\x45\xf4\xe1\x02\x64\x07\x0a\xfe\x52\xb9\xbf\xab\xa3\x08\xea\xcc\x72\xfd\x51\x84\xda\xcb\x4c\xb5\x3e\xea\x37\xe1\x7f\x3a\xa0\x8b\xb2\xe2\x9b\x06\x6a\x23\x1f\x2a\xe7\x49\x7b\xa2\xda\x99\x92\x57\x41\xd0\xcc\x83\x5b\x46\x77\x81\xa5\xf7\x97\x5f\x6f\x77\x0e\xf8\x4c\xde\x0d\xd9\x2e\x47\xf1\x42\x1f\x16\xff\x5c\xb7\xe9\x0c\x96\x83\x27\x6e\x36\x4f\x4a\x0e\x13\xc2\x28\x14\x40\x64\xc4\x63\x59\x31\x20\x8f\xb7\x0a\xce\xa8\x43\x4a\xc6\xf5\x2c\x7d\x0b\xeb\x16\xe6\x67\xd1\x64\x1e\xc2\x4e\x8b\xf3\x73\xaf\xeb\xa0\x2f\x62\xd4\xcf\xc1\x4d\x27\x90\xd7\xc1\x47\x47\x81\xd7\xc6\xb5\x7b\x6b\xac\xaf\x05\x6c\x1c\x50\xb6\x2f\x83\xc8\x17\x97\x8e\x88\xd0\x86\xe2\x3d\x30\xd3\x04\xdf\x6c\x62\x29\xe9\x66\xe1\xaa\x4d\x9e\xdf\x3f\xfc\x86\xf7\xb8\x37\x75\x93\x6c\xf5\xf6\x86\xae\x63\x94\x85\xbc\xbb\xf1\x25\x94\x27\x95\x21\x1a\x44\xf5\x53\x16\x9b\xc6\x4a\xbf\xf3\xd1\xb1\x3b\xdb\x66\xc7\xe3\x17\xa0\xb4\x74\xd3\x22\x81\x10\x83\x77\x0c\xc0\xe4\x17\xa1\xd4\xf7\x9d\x7e\x97\x4d\x0c\x04\x5b\x7f\xe3\x69\xa7\x82\x4c\x98\xad\x2d\xc4\xb3\x22\x9f\xff\x0e\x00\x00\xff\xff\x76\x3c\x6d\x8e\x49\x53\x00\x00")

func fcsLpcMotorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_fcsLpcMotorHtml,
		"fcs-lpc-motor.html",
	)
}

func fcsLpcMotorHtml() (*asset, error) {
	bytes, err := fcsLpcMotorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fcs-lpc-motor.html", size: 21321, mode: os.FileMode(420), modTime: time.Unix(1553030360, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x59\xff\x6e\xdb\x38\xf2\xff\xdb\x7e\x0a\x96\x7f\x6c\x6c\xd8\x92\xd3\xef\xf7\xb0\x38\xc4\x76\x0e\x69\x9a\x62\x7b\xdb\x4d\x8a\x38\xc5\xe1\xae\x0d\x0a\x5a\xa2\x6c\x36\x94\xa8\x13\x29\x3b\xde\xae\xdf\xe9\x9e\xe1\x9e\xec\x66\x48\x5a\x92\x63\x27\x4d\xb2\xc5\x1e\x70\x05\x1a\x8b\xe4\x70\x38\xf3\x99\x1f\x9c\x91\x46\x2f\x5e\x5f\x9c\x5e\xfd\xfd\xfd\x19\x99\x9b\x54\x1e\xb7\x47\x2f\x82\x80\x9c\xaa\x7c\x55\x88\xd9\xdc\x90\xff\x3b\x7c\xf9\x23\xb9\x9a\x73\x32\x53\x81\xd4\xda\x90\x93\xd2\xcc\x55\xa1\x43\x72\x22\x25\xb1\x34\x9a\x14\x5c\xf3\x62\xc1\xe3\xb0\xdd\x82\xcd\x1f\x34\x27\x2a\x21\x66\x2e\x34\xd1\xaa\x2c\x22\x4e\x22\x15\x73\x02\xc3\x99\x5a\xf0\x22\xe3\x31\x99\xae\x08\x23\xaf\x26\xaf\x03\x6d\x56\x92\xdb\x6d\x52\x44\x3c\x83\xad\x66\xce\x0c\x89\x58\x46\xa6\x9c\x24\xaa\xcc\x62\x22\x32\x98\xe4\xe4\xdd\xdb\xd3\xb3\xf3\xc9\x19\x49\x84\xe4\x61\x3b\x08\x40\x56\x27\x72\x6b\x34\xe7\x2c\x86\xdf\xd6\x28\xe5\x86\x91\x8c\xa5\x7c\x4c\x17\x82\x2f\x73\x55\x18\x0a\xa7\x67\x86\x67\x66\x4c\x97\x22\x36\xf3\x71\xcc\x17\x70\x54\x60\x07\x7d\x92\x8a\x4c\xa4\x65\x1a\xe8\x88\x49\x3e\x7e\x19\x1e\xf6\xe1\x3c\x61\x04\x93\xcd\xa9\x12\x14\xb4\x63\x36\x85\xa9\x15\xd7\xb4\x3e\x2e\x9a\xb3\x42\x73\x60\x5f\x9a\x24\xf8\xb3\x5b\x30\xc2\x48\x7e\xfc\xe6\x74\x42\xde\xbd\x3f\x25\x57\x5c\x9b\x29\xcf\xa2\xf9\x68\xe0\x16\x90\x44\x47\x85\xc8\x0d\xd1\x45\x34\xa6\x53\xb5\xe4\xc5\xe7\x48\xa5\xb9\xca\x40\x52\x3d\x58\xf2\x69\x3d\xfa\x72\x67\x1c\x48\x61\x78\x08\x82\x87\x5f\x40\x8e\xd1\xc0\x71\xb2\x4c\xa5\xc8\x6e\xc0\x1c\x72\x4c\x45\xea\x94\x9f\x17\x3c\xd9\x73\x40\xce\x72\xd4\x08\xd1\xdf\x1e\x04\x91\x64\x5a\x73\x1d\x22\xb6\xf4\x41\xa6\x49\x04\xa2\xe4\x51\x90\x2a\xa3\x8a\x06\xbd\x65\x74\xdc\xc6\x09\xf2\xb5\xdd\x42\x9b\x27\x52\x2d\x83\xd5\x11\x61\xa5\x51\xc3\xf6\x1a\xa8\x5a\x53\x15\xaf\x70\x19\xfe\x25\x60\xa0\x20\x61\xa9\x90\x40\x72\x70\xa9\xa6\xc0\xf0\xa0\x4f\x0e\x7e\xe2\x72\xc1\x8d\x88\x18\x39\xe7\x25\x87\x99\x6a\xa2\x4f\x4e\x0a\x30\x51\x9f\x68\x96\xe9\x00\x8c\x23\x92\x61\xcd\x6a\xc9\xd1\x2f\x8f\xc8\xff\x1f\x1e\xda\x59\x3c\x0f\x60\x72\x62\xc1\x93\x77\x97\x51\x0d\x1c\x05\x03\x13\x6d\x0a\x11\x19\x3a\x6c\x2f\x58\x41\x40\x39\x32\x46\xf9\xd0\xf4\x47\x84\x9e\x0f\x4e\x68\xbf\xdd\x32\xea\x86\x67\x30\xc4\x67\x16\xc7\xb8\xf2\xf5\x6b\x78\x02\x4f\xeb\x35\xce\x01\x78\xe6\x73\x99\xc7\xcc\xf0\x7a\x53\x01\xe7\x81\x66\xa8\xec\xed\x11\x49\x98\xd4\x1c\x66\x5b\xbf\xd6\xcf\x6b\xf8\x0f\x7e\x94\x39\xa2\x28\x8d\xf5\x11\xc9\x4a\x29\x91\x0c\x78\xb1\x7a\x84\x9e\xc0\xd2\x6a\x0c\x1b\xd7\xc3\x76\xbb\x93\x94\x59\x64\x84\xca\x3a\x5d\xe4\x80\x0a\x58\x49\x41\x85\x19\x37\xa7\x4a\xdd\x08\xde\xa1\xe0\x8f\x9f\xaf\x2e\x7e\x3e\x3b\xa7\xdd\xa1\x23\x42\xe5\x76\x69\x3e\x4c\xce\x2e\x2d\x09\x80\x10\x7a\x12\xfc\xf1\x33\x1b\xce\xf6\x17\xac\xd9\xe9\x76\xdb\x6d\x4b\x9a\x4b\xc5\xe2\x89\x45\xf5\x0d\x44\x29\xd0\x54\x72\x61\xd4\x56\xb2\x21\x20\x96\x6b\xc6\x97\x04\x29\x2f\xed\x44\x07\xcf\x74\x6b\xa1\xca\x90\x17\x87\xf0\x6f\x30\xb1\x0c\x2c\x07\x80\xc8\x99\xa7\xd5\x32\xab\x1c\xb0\x26\x34\x32\x12\xc1\x6e\xb5\xac\x3f\x02\xfa\xb7\x6e\xe8\x4d\x56\x49\x6e\x27\x31\x47\xe0\x26\x27\x72\x00\x91\x91\xb2\x2c\xd6\x6e\x87\xc3\xdf\x4b\x19\x42\x76\x2b\xa5\xc1\x05\xf4\x23\x64\x83\x96\x0a\x91\x28\xd4\x20\x60\xe7\xaf\x93\x8b\xf3\x10\x9d\x27\x9b\x89\x64\xd5\x81\x85\x2e\x2a\xb2\xae\x74\xc1\x9f\x13\xfd\x4a\x64\xac\x58\x4d\x2c\x9d\x83\x03\x23\xa1\x01\xdc\xa9\x17\xa2\xa9\x31\x5f\x40\xb8\x56\xb8\x25\xaa\x48\x61\xf5\xbd\x92\xab\x14\xf8\xc6\x2a\xf5\x04\xa1\x54\x90\x9c\xae\x58\x01\x96\x0c\x73\x56\xc0\xdc\x99\xe4\x29\xfc\x80\x73\xdc\x01\xcc\xe1\x55\xc1\xb5\x8d\xd6\x3e\xb0\x1c\x56\x7b\xa1\x72\x48\xa1\x58\xa1\xe1\xb7\x06\x0e\x66\xe1\x82\xc9\xd2\xfa\x74\xfb\xf1\x60\x79\x1c\xf0\x22\x31\x0f\xc3\xf0\x14\xdd\x1d\xbf\x4e\xcd\x1f\x25\xf0\xec\xf7\x73\x1f\x0c\xf0\xb6\xd0\x0a\x6e\x18\xa9\x66\x1d\x8a\x1b\x40\x52\xe2\x95\x0e\xc3\x90\x74\x70\x8e\x17\x63\xda\xb3\xdb\x42\x37\xec\xd1\x6e\x15\x54\x16\x51\xe0\xef\xd6\xed\xc8\xaf\x60\x28\x7b\x2b\x78\xd8\xed\x4f\x7f\xd7\x2a\xf7\xdb\xc1\x71\xc5\x67\xc4\x18\x18\x23\x53\x07\x3a\xb0\x3e\x2f\xd3\x29\x04\x52\xac\xa2\xd2\x22\x00\x98\x78\x30\x5e\xad\xde\xc6\x1d\x9b\xb3\xe1\x8c\x80\xf6\xec\xc9\x3d\x1a\x6c\xf4\x40\x8e\x30\xd4\x52\x80\x3a\xb4\xeb\x38\xa2\x4a\xdf\xc0\x04\x84\xee\xdd\x31\x2a\x4a\x84\x56\x7d\x8c\x03\x6c\x68\x1b\x16\xba\x44\xab\xfd\xef\xd8\xe7\x0f\xc7\x6f\x62\x54\xfe\x34\xf8\x26\x57\x17\xef\x1b\x18\xfe\x41\x48\x51\x0d\x82\xd2\xff\x06\x4a\x50\xbe\x40\x6a\x98\xff\xa2\xb2\x26\x4e\x96\x68\x93\x6a\x21\xd9\x19\x4c\x41\x1f\xa9\xe1\x29\x94\x47\xcc\x94\x05\xa7\x7d\x42\x73\xa5\x05\x92\xe3\x73\x91\xa7\x9a\x5e\x7b\x48\xa0\xde\x35\x0a\xb7\xd8\x80\x74\x23\xbc\x31\x01\xc1\x0e\xae\x0b\x58\x3a\x1c\xc2\xcf\xc8\x31\x0f\x25\xcf\x66\x66\x0e\x33\xbd\x5e\x7d\xb1\x09\xcc\x4c\x76\xfd\xa3\x40\xce\x95\x30\xc8\xf8\xa1\xa8\x4e\x55\x06\xb1\x6c\xcf\xae\x43\x5b\xc4\x68\xca\x16\xee\x0f\x45\x96\xf1\xe2\xa7\xab\x5f\xde\x01\x27\x27\xdd\x47\x11\x5f\xdb\x8b\x6a\x1b\x96\xbf\xd9\x12\xe3\x5e\x64\x32\x2c\xe6\xbf\x21\x8c\xab\x52\xa0\xa4\x84\x0b\x21\xb5\xee\x84\xbb\x42\x26\x51\x0d\x7a\x0e\xda\x9c\x2c\x98\xb0\xd5\x34\xdd\x2c\x42\x25\x8c\x8b\xb6\xdc\x11\x29\x9b\xf1\xc1\x97\x9c\xcf\x86\x53\xa6\xf9\x8f\x7f\xea\x3b\xd5\x9c\x05\xbd\x58\x8d\xb2\x25\x42\x8f\xaa\x25\x84\x01\xb0\xb2\x93\xa4\x47\xe8\x98\x7a\x13\x45\xac\x29\x78\x64\xf7\x86\x3a\x87\x7a\xba\x73\x30\x3c\xe8\x3a\x6b\x39\x63\x8d\xd1\x54\xa3\x88\xed\xb7\x12\x4a\x1a\x31\x6f\xa1\xe5\x1c\x6b\x9d\x4e\x84\x3e\x58\x9c\x98\xce\x61\x77\x3c\x3e\x20\x07\x5d\x47\x15\xea\x72\xea\x1c\xb1\xf3\xd2\x1a\x43\x24\x48\x2b\x20\x2b\xdd\x5e\x24\x1d\x27\xf8\x18\x5c\xa3\x0b\x05\x07\x38\x59\xb6\xb5\x05\x97\xbd\x08\xfd\xc8\x3f\x54\xc5\x85\xa5\xa6\xd4\x82\xb2\x04\x86\x6a\xe9\xcb\xa6\x9d\x9a\xa9\x8a\x10\x1f\xb7\x58\x77\x81\x99\x27\x2a\xba\x81\x0b\x92\x2e\xf5\xd1\x60\x40\x01\x2a\xa4\xc3\xb2\x16\x51\x1b\x20\x6d\x55\x06\x56\xf1\xf5\xd8\xdd\x48\xbb\xbd\x7b\xb9\xf1\xab\x47\xed\x5f\xc0\x3d\xa4\xe8\x76\x7c\x5b\xef\x56\x99\xca\x79\x76\x47\xc5\xf5\x2e\x55\x24\x95\xe6\x3b\x64\xbb\x74\x29\xd7\x1a\xbc\x6d\x7f\xc6\xb4\xd4\x8d\xa2\xde\x4b\xff\x1a\x1e\x6d\xb5\xda\x4c\x86\x36\xe7\xe4\xd8\x0c\x3a\x06\xf6\x80\x8a\x68\x93\x4b\xeb\x08\xdd\xac\x54\x9c\x1f\x0c\x29\x6d\x20\x05\xe9\xc0\x11\xd3\xca\x93\x36\x9b\xc7\xb6\x23\x70\x32\x63\x32\x4d\x19\xb4\xcd\x39\xe8\xd5\x27\x2c\x31\x50\x6c\x47\x05\x87\x72\x1a\x1b\x5b\x4d\x16\xd8\x30\x41\x1b\x85\xca\xf6\x21\xa0\x0d\x36\xd8\x05\x27\x2b\xa8\xa0\xc2\x10\x39\x38\xe7\x1a\xda\xa2\x17\xfe\xb8\x53\xea\x14\x32\xbe\x03\x0b\xa4\xf6\x0f\x57\xa7\xbe\xb6\xb5\xb2\x6d\xa7\xd9\x4e\x05\xc5\x4e\x9e\xb1\x4b\xde\x37\x2c\xc5\xdd\x9b\xc0\x26\x84\x07\xb2\x7f\xd5\x51\xac\x36\xe8\x3a\xa8\x10\x13\x6a\xe7\x69\x6d\x03\x68\xed\x6b\x2b\x24\x3a\x75\x44\x65\x66\x49\x1e\x83\x7f\x5d\x38\x41\xa8\x89\x0c\x2c\x11\x32\x70\x9a\x05\xdf\x30\x75\xd3\x4f\x67\xe7\x24\x6d\x70\xb3\x13\x4f\xe7\xe3\x75\x6c\x72\xf2\x53\x4f\xe7\xa5\x57\x59\xb4\xa3\x20\x4e\x3e\x43\x3d\xbc\x2a\xbb\x5b\x97\x90\xe5\x86\xf3\x4f\xe7\xc6\xb2\x99\xe4\xfb\xd8\xd9\x05\xcc\x1f\xe4\xdf\xff\xb2\x36\xad\xaf\x5f\x97\xd1\x2d\x15\x5e\xe9\xd5\xfd\x5b\x25\xf6\x27\xca\x80\x4c\xf0\x8a\xdd\x23\x85\xe5\x0f\xb7\x43\x0f\xc5\x38\xa5\x55\x18\x3d\xd7\x92\x3b\x07\x80\xe7\x5a\xd7\xc7\x0c\x50\x19\xe5\x59\x5a\x78\x13\x37\x4f\xa0\x17\x3f\x3b\x91\x09\x97\x90\x41\xbf\x13\xd3\xb3\xcb\xcb\x1a\x88\x4a\x6e\x17\x2d\xcf\x92\xbc\x8a\xbf\x6d\xd9\xdd\xec\xf0\xe9\xfc\x36\x01\xb8\x03\xb6\xa3\xfc\x7d\x88\xdc\x23\x6c\x92\x7c\x4f\x69\xed\xfb\xa6\x67\x3b\x5b\xbc\x37\xa0\x70\xfe\x39\x49\x28\xe6\x7b\x72\x59\x95\xff\x61\xf4\xd1\xd2\x5e\x37\x13\x9d\xab\x7f\x6d\xd5\x8d\x80\xec\x69\x56\x21\x5b\xb8\x57\x34\x7b\xd6\x6c\xe8\x07\x75\x79\x0e\x64\x55\x01\x1d\x0b\x8d\xb5\x66\xdc\xd4\x8a\xbc\x18\xdb\x6a\x9e\x92\xdf\x7e\x23\x2f\xb6\x84\x28\x98\xd0\x77\x88\xc7\x1b\xe2\x1f\x7e\x68\x6a\xb3\xa7\xb8\x07\x0d\x76\x8b\xc6\xba\xb6\x87\x65\x5f\x37\xde\x8b\xa9\x88\x9b\x8d\x78\x53\x76\xff\xf8\x88\xdd\xe5\x34\x15\xe6\x77\xef\xae\x80\x70\x0f\xce\xb7\x5a\xeb\x9d\x9e\xeb\x9b\x35\x99\xa7\xda\x5f\x93\xed\x90\x3d\x5c\x92\x21\x94\x6a\xfa\xe5\x9b\x15\x97\x85\x1b\x08\x43\x61\x05\xdf\x2a\x2a\x0a\x1e\x71\x70\xcd\x18\xdf\x8b\xed\xe9\x2d\x61\x57\xb7\x2a\xb1\x60\xf0\x91\xf2\xa2\xa0\xd7\xd6\x65\xe8\xdd\x6c\xf5\xcf\x92\x17\xab\x09\x97\x3c\x02\x47\xdc\x34\x63\xf0\xa8\x03\x26\x79\xe1\x8d\x90\x4b\xb6\x3a\x2b\x0a\x20\x40\xde\x1e\xc8\x96\x5e\x0a\x13\xcd\xfd\x11\x10\x45\x33\x4e\xaf\x3d\xfb\x08\xfa\x1e\x42\x67\x3c\x0b\x62\x05\x29\xe2\xc8\xbe\xde\xbe\x37\x0a\xb5\x80\x6e\x2b\xb7\xd1\xdb\x8c\x3b\xfb\x7e\xd9\xbd\x19\xdf\xd5\xa4\x51\x30\x3e\xc4\x39\x4b\x4a\x0d\x16\x08\x50\x14\xda\xdb\x0c\x3f\x43\x6b\xb9\x95\x2d\x76\x11\xf4\x47\x75\xdd\xf9\xeb\x6d\x29\x52\x3d\xdb\x23\x85\x7d\x71\xcd\x6f\x6d\x97\x38\xca\x0b\x7c\x27\xbf\x92\x7c\xfc\x89\xe2\x24\xc0\x29\x66\xd9\x11\x91\x3c\x31\x9f\xe8\xb1\xcb\x9a\xbe\xb0\xd3\x33\x6f\x6b\x78\xf2\x1d\x1d\xfd\x94\x51\x7f\xb6\x8b\xd3\x46\x8c\x22\xd5\x76\x8c\x12\x47\xd8\xb2\xa7\xf7\xc6\x48\x01\x61\x8a\x95\xc4\x68\x5a\x54\x67\x39\x25\x2a\x22\x3a\x1a\x80\x90\xc7\xf4\x3b\x41\x88\x6c\xb7\x0f\x9a\x42\xaa\xb9\xf1\x53\xce\x21\xb0\x95\x6f\x7a\xc4\x23\x4e\xb5\x5b\x1e\x38\xd6\x39\xdf\x62\x46\xaf\x87\xf7\x9d\xfa\xab\xc8\xb7\x0e\x75\x39\x75\xf1\x50\xcf\x52\x1d\x5f\x70\xfc\x68\x74\x47\x80\x86\xed\x62\xb9\xd5\x94\xc3\xd9\x86\x7b\x4e\x10\x4c\xaa\x48\x2b\x2b\xc6\x32\xb4\x31\x5d\xf3\x76\x89\x2a\x40\xaa\x00\x3b\x48\x17\xe9\x8e\x34\xe5\x66\xae\xe2\x31\x66\x6d\x43\xeb\x69\x66\xf3\xc9\x98\xce\x8d\xc9\xa1\xef\xac\xbe\xde\x0c\xa0\x7d\xb6\xcd\xf3\x5f\xf0\xdd\x3c\xbe\x3e\xcc\xf0\xb3\xe5\x87\xcb\xb7\xe8\xb0\x21\x7e\xf2\x6a\xca\xd1\xbc\x76\x47\xee\x03\xda\xb4\x34\x46\x65\x9b\xcb\x03\x93\x9d\x88\x6e\xc0\x77\x9d\x90\x6f\x40\xc6\x0e\xbe\x91\xe9\xd1\xee\x27\x4a\xec\x67\x36\x58\x8c\x94\x54\x45\x52\x4a\x98\x12\x31\x12\x6f\x54\xdb\x88\xb3\x0d\x5c\x8f\x82\xe7\xbf\xf6\x4b\xe4\x1f\x6f\xdf\xa3\x77\x76\x26\x22\x2d\xa5\xed\xe6\x48\x40\xfc\x19\xa3\x41\x53\xaa\xca\x83\xc1\x6a\x21\xcb\x73\x7c\x11\x3e\x17\x32\xee\xc4\xb2\x3b\x24\xbb\x66\x5f\x6f\x7e\x7c\xfc\x6c\xf7\xf0\xf7\xe6\xfb\xfd\xb4\xf7\x75\xe2\xfb\xa9\x1f\x4e\xfe\x0d\xa1\xd6\xf8\x2d\x6f\xf3\x0d\x0f\xe6\x46\xf6\x93\x62\x99\x41\x6f\xa9\xa4\xcd\xed\x16\x63\x0a\x99\x57\x95\x06\x9b\x5e\xfc\x74\x28\x49\x04\xcc\x00\x18\xf7\x43\x8f\x91\xdb\x08\x9d\x19\xf0\xb7\xc9\x1b\xe0\xb1\x1f\x32\x5b\xa3\xad\x2f\x9c\x38\x7f\x3c\x1a\xec\xce\xe1\xfe\x01\x30\x40\x4e\xa3\x01\x0a\x71\xdc\x1e\x0d\xdc\xb7\xe9\xff\x04\x00\x00\xff\xff\x02\xa0\x29\x09\x5f\x1f\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 8031, mode: os.FileMode(420), modTime: time.Unix(1553030546, 0)}
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
	"bower.json": bowerJson,
	"favicon.ico": faviconIco,
	"fcs-lpc-motor.html": fcsLpcMotorHtml,
	"index.html": indexHtml,
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
	"bower.json": &bintree{bowerJson, map[string]*bintree{}},
	"favicon.ico": &bintree{faviconIco, map[string]*bintree{}},
	"fcs-lpc-motor.html": &bintree{fcsLpcMotorHtml, map[string]*bintree{}},
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: k}
	}
	panic("unreachable")
}
