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

var _fcsLpcMotorHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3c\xfd\x72\xdb\x36\xf2\x7f\x4b\x4f\x81\x30\xbf\xdf\xd4\x6e\x42\x4a\x56\xec\x5e\xaa\x48\xba\x4b\xdd\x74\xda\x9b\x24\xcd\xd4\xb9\xeb\xb5\x99\x4c\x06\x22\x21\x89\x31\x45\x30\x24\x64\x5b\xf1\xe8\x9d\xee\x19\xee\xc9\x6e\x17\x1f\x24\x48\x51\x9f\x91\x93\xe9\x5c\x33\x13\x1b\x5c\x02\xbb\x8b\xc5\x62\x3f\x80\xa5\x7b\xf7\xbe\xff\xf9\xfc\xf5\x6f\xaf\x9e\x91\x89\x98\x46\x83\x66\xef\x9e\xeb\x92\x73\x9e\xcc\xd3\x70\x3c\x11\xa4\xd3\x3e\xf9\x86\xbc\x9e\x30\x32\xe6\x6e\x94\x65\x82\x3c\x9d\x89\x09\x4f\x33\x8f\x3c\x8d\x22\x22\xfb\x64\x24\x65\x19\x4b\xaf\x58\xe0\x35\x1b\x30\xf8\x1f\x19\x23\x7c\x44\xc4\x24\xcc\x48\xc6\x67\xa9\xcf\x88\xcf\x03\x46\xe0\x71\xcc\xaf\x58\x1a\xb3\x80\x0c\xe7\x84\x92\xef\x2e\xbe\x77\x33\x31\x8f\x98\x1c\x16\x85\x3e\x8b\x61\xa8\x98\x50\x41\x7c\x1a\x93\x21\x23\x23\x3e\x8b\x03\x12\xc6\x00\x64\xe4\xf9\x4f\xe7\xcf\x5e\x5e\x3c\x23\xa3\x30\x62\x5e\xd3\x75\x07\xcd\x66\x2f\x0a\xe3\x4b\x20\x1f\xf5\x9d\x70\x9a\xf0\x54\x38\x64\x92\xb2\x51\xdf\x19\xf2\x6b\x96\xbe\xf3\x39\x00\x63\x16\x8b\xac\x95\xf0\x68\x3e\x65\xa9\xf9\xed\xe1\x64\x1d\xc4\x90\xf9\x69\x98\x08\x92\xa5\x7e\xcd\xa8\xf7\x1f\x66\x2c\x9d\xb7\x82\x30\x13\xba\xed\xbd\xcf\x9c\x41\xaf\xa5\x46\xed\xc4\x41\x98\xf2\xd8\x1d\x45\xec\xc6\x8d\xe8\x9c\xcf\xc4\x12\xc0\xf5\x23\x9a\x65\x2c\x2b\x78\xdb\x0d\x37\x9d\x09\x3e\x4e\xf9\xb5\x2b\xd8\x8d\xa0\x29\xa3\x2b\xc0\x06\xff\xae\xac\xf3\x74\x5a\xb4\xf6\x44\x12\xfa\x3c\xb6\x9b\x7b\xa2\x49\xe8\x98\xd9\xcd\x3d\x44\x96\xd0\x84\xa5\xee\x70\x26\x04\x8f\x4b\x0f\xbb\xb3\xa4\x46\xfb\x34\x0d\xac\xe6\xbe\x68\x82\x90\x46\x7c\x5c\x7a\xd8\x1b\x55\x4a\x01\x0e\x12\x8a\x59\x54\x03\xda\x17\x2d\x2e\x5c\x59\x70\x16\x64\x6f\xa4\x71\x02\x5b\xc2\x6a\x1f\x02\xd1\xfe\x1a\xaf\x71\x09\x36\xb5\x9a\xfb\xa2\x99\x52\xc1\x52\x58\xc9\xca\xe3\xbe\xe8\xc0\xf6\xf0\x28\x72\x27\x8c\x06\x95\xd5\xad\x79\xb3\x37\x91\x28\x0c\xd0\x5c\x5a\x0f\xfb\xa2\x12\x74\x68\x37\xf7\x46\xc3\xc7\xe3\x88\x95\x15\xaf\x04\xdb\x1f\x31\x8f\x86\x34\x2d\x3f\x15\x16\x25\xe0\x53\x77\xca\x83\x59\x04\x1e\x2c\xe8\x3b\x23\x3f\x73\xa3\xc4\x07\x90\xe0\xa9\x0b\xbb\x33\x8c\xa1\x5b\xa3\x07\xfa\x91\x44\xb0\xb2\xd0\x6e\xf4\xa4\x57\x1b\x34\xbb\x13\x0e\x1e\xf3\xb6\xd9\x00\x07\x02\x2f\xe7\x5d\x32\x8c\xb8\x7f\xf9\xa4\xd9\x18\xf2\x1b\x37\x0b\x3f\x86\xf1\x18\x60\x3c\xc5\xe5\x02\x10\xbc\x40\x95\x75\x69\x14\x8e\xe3\x2e\x01\x7f\x08\x9a\x02\xc0\x29\x4d\x81\x4c\x97\x9c\x25\x37\xf2\xe9\xc6\xbd\x0e\x03\x31\xe9\x92\xce\x59\x5b\x81\xc2\x38\x07\xb5\x25\x68\x01\x34\x94\x54\x62\xae\x9c\x2c\xf2\xd1\x70\xaf\xd9\xf0\x32\x04\x0a\x49\xc2\x68\x4a\x63\x9f\x75\x49\x0c\xd2\x00\x1c\x0d\xc5\x47\x97\xb4\xf1\x21\xa1\x41\x20\xb9\x93\x4f\x43\xea\x5f\x82\x0f\x01\x67\xdc\x25\x02\x86\x65\x09\x6c\xab\x58\xc0\xab\x05\x4e\xb7\xa5\xe7\x8b\x6d\x74\x0e\x64\xca\x20\x48\x00\x61\x25\x30\x7f\x87\x50\x5f\x84\x3c\xee\x3b\xad\x5c\x5a\xd0\x2f\x9c\x8e\x09\xbd\xa2\xb0\x3f\x95\xef\x9d\x08\x91\x74\x5b\xad\xeb\xeb\x6b\x0f\xe3\x0c\x8f\xa7\xe3\x56\x06\xbb\x2e\x6b\x05\x6c\x44\x67\x91\x68\xa1\xdf\xcf\x10\x07\x7f\x77\xe2\x25\xf1\x18\x10\x47\xa2\xef\x28\x24\x4e\x4b\xa1\xb5\x8c\x00\x89\xe9\x94\xf5\x9d\x19\xc4\x26\xd8\x72\x48\x44\x87\xa8\x14\x10\x9f\xa4\x0e\xb9\xa2\xd1\x0c\xde\x5e\x85\x40\x84\xc3\x73\xca\x3e\xcc\xc2\x14\x42\x13\xf0\xee\x16\x92\x55\x58\x13\x70\xd5\xd7\x20\xb0\x1c\xeb\xab\x1c\x20\xe6\x49\xa9\x83\xc1\x5c\x87\x58\xad\x11\x91\x9e\xbf\xef\x98\xa5\x32\x38\xb2\xd9\x70\x1a\x0a\x88\x37\x6c\xff\x44\x52\x1a\x66\x80\xed\x42\xbe\xec\x95\x7c\x17\x90\xd0\x0d\xb9\x2c\xb8\x16\xa8\x9b\xad\x42\x39\xe1\xc9\xc4\x2e\x8d\xc6\x2b\x15\x0c\x1d\xa1\x66\x34\xc2\xac\x5b\xab\xdc\x0f\xf1\x25\xbc\x78\xa7\xb8\x79\x27\xa1\x5d\x32\x9a\xc5\x72\x59\x8f\xd8\x15\x68\xc2\xb1\xd4\xae\x46\x03\x7d\x3a\x87\xf0\x0c\x3a\x1d\x21\xb6\x77\x13\x1a\x07\x11\x33\x83\x9c\x07\xaa\xf7\x13\xd9\xf9\x0a\xd6\x5e\xaa\x4b\x9f\x04\xdc\x9f\x4d\xe1\x8d\x37\x66\xe2\x59\xc4\xb0\xf9\xdd\xfc\xa7\x00\x70\xc0\x7b\x47\xf7\x2f\x23\x87\x17\x88\x10\x7f\x3f\x70\x08\x04\x90\x4a\xe9\xba\xe4\x8d\x02\x7a\xea\xf9\x81\xf3\x16\x5f\x6a\x1d\x34\xef\xd4\x23\xbc\x33\xb8\xad\x11\xc0\x8e\x52\x5c\xeb\x8d\xea\x8f\x6f\xb4\x12\xd7\x70\x04\x64\x06\x83\x03\x70\xa1\xe4\x7c\xa4\x41\xad\x56\xca\xc4\x2c\x8d\x09\x48\xd3\x53\xd2\x7c\x8e\x1c\x1c\x59\x92\x5c\xc8\x35\x32\x1a\xf7\x4f\x30\x1f\x01\x45\xcc\xaf\xa8\x00\x03\x82\x72\xf7\xbe\xfe\xab\x53\xea\xf4\x2c\x4d\x79\xfa\x82\x65\x19\x84\x4f\xf0\x7e\xc8\x39\x97\x1d\x16\x88\xb1\x88\x6f\x7b\xad\xc2\xfc\x6d\x63\x0c\x21\x8c\xfd\x1f\xb7\x86\xa5\x50\x12\xed\x09\x12\x9a\x30\x4c\x94\xba\xe4\xf4\xc4\x7b\xf4\x17\xc5\x77\x43\x73\x48\x1e\x2b\xa6\x2d\xac\xae\xcf\x23\x0e\xa4\x61\x7f\x1c\xb9\xae\x42\x08\x98\xdc\xb3\x76\x5b\xae\xb7\x7e\x7d\x3d\x01\xf3\x28\x89\xee\x68\x83\xf5\x22\x7d\x92\xf9\x31\x63\x80\x2f\x67\xf0\x5c\xe2\xbc\x63\x53\x84\x5c\x6f\xd6\xd0\x0d\x2a\x0a\x2b\x2d\xf5\x53\xe9\x40\x08\x13\xf0\x67\x99\x80\x11\x6a\xe2\x96\x92\xea\xe5\x39\x69\xb7\xff\x1f\xe8\x99\x05\xd4\x8f\x46\x05\xdb\x04\xb3\xab\x75\x1a\xfc\x37\xa0\x18\xcd\x61\x19\x75\x8e\x07\xb9\x73\xf8\x91\xc7\x82\x46\xc7\xa8\x96\x52\xd1\x82\xb9\xb2\x9e\xb9\x96\x9d\x6a\x95\x90\xef\x57\x86\x76\x7a\x10\x07\x07\x06\x4b\xdb\x85\xd0\x8a\x47\x33\xa9\x11\x8d\x86\xe0\x49\x57\xea\x6a\x43\xa6\xe8\xba\x3d\xe4\xb0\x20\x53\xad\xc5\x8d\x88\x8d\x84\x69\xaf\x57\xbe\x71\xca\xe6\x2e\x6c\xa5\x87\xe4\x3e\x63\xec\x78\x0b\xd6\x3c\xb0\x8c\x02\xf6\x44\x65\x62\x8f\x97\xe6\x85\x11\xa1\xea\xb4\xcc\x42\x84\xac\x8f\x53\x3a\x97\x2c\x1a\x66\x70\x84\x9b\xb1\x88\x49\x8d\x76\x21\x56\x33\xfd\x41\x17\xab\xd8\x55\x30\xb7\x8a\xc0\x30\xa2\xd2\xfe\xd4\x0d\xf1\x44\x28\xb4\x99\xb0\x96\xdb\x9e\x80\x07\xfb\xde\x67\x1a\x79\x75\x9d\x31\xaf\x2f\x04\xe5\x29\xe3\xa5\xba\xd6\x1a\x35\xdd\x6f\x4a\xc3\xd8\x5e\x5d\x3c\x2f\x19\x45\x90\xbc\x83\xb5\x54\x72\xce\xfb\xb6\xbe\xc6\x0e\xc4\x83\x1e\x22\xf4\x69\x04\x32\x51\x4e\x4a\x79\x62\x52\x30\xdd\x39\x4d\x6e\x8a\x1f\x52\x98\x44\xa2\xa8\x19\xed\xe2\xc2\x01\x13\x86\xd9\x06\xa9\xce\xcc\x8c\xd0\xfe\xa9\x78\xaf\xe6\xe2\xbe\x87\x0d\x15\x8e\x42\x16\x1c\x1b\x4a\xf0\xe3\xeb\x56\xb3\x69\xdb\xa8\x9e\x6d\x06\x4c\x7c\x55\xa7\x48\xa3\xf0\xa6\x30\x36\x28\x54\x47\x8e\x30\x43\xf4\x6a\x49\x6b\x96\x87\x69\x45\x36\x4a\xb0\xdd\x77\x20\x9c\x98\x39\xa4\x94\x01\xab\xcc\xa1\x08\xc9\x8a\x31\x1a\x57\x10\x5e\x19\xb2\x7a\x9d\x95\x42\xa8\x59\x82\xcd\xbb\xb8\x78\x4d\x7e\x38\xbf\x20\x2f\xd0\xb0\x64\x60\x86\xc2\x2b\x3d\xb4\xce\x74\x01\xa5\x5a\xb0\x9c\x4b\xab\x32\x99\x66\x85\x01\xbd\x99\x9c\xdc\xa0\x49\x35\xc9\x57\xca\x29\x4d\xdf\x4e\xf1\xd5\x0b\x85\x4b\xc1\x73\x94\x7a\x1a\xaa\x83\xec\x31\xc0\xc0\xb8\x4b\x7a\x30\xdd\x38\xa7\x34\x93\xc1\xb2\xdc\x2f\xda\xd8\xbf\x6c\x3d\x85\x85\x84\x3e\x03\x6b\xce\x80\x61\x92\xe6\x4d\x9d\xe7\x82\xd4\x0d\x28\x5f\x1b\x58\xf7\x81\x94\x98\xfb\x2f\xb0\xa6\x8f\x4e\xbd\x93\xb3\xc7\xf0\xff\xcc\x3b\xf9\xa6\x7b\xd6\xee\x9c\xd4\xc0\x3a\x3d\x2b\x07\x5f\x83\xf0\xf7\x3a\x84\x8f\xea\x06\x1b\x98\xe2\xd0\x00\xa5\x60\xe5\x62\x82\x37\xa0\x62\x96\xe9\x35\x2b\xc3\x6a\x7a\xd3\x08\xb6\x44\xb9\xb3\x02\x69\xe9\x4b\x31\x11\x6b\x29\x70\xf9\xcc\x42\x14\x3b\xde\xa9\xc8\x4f\x9a\x46\x65\xe8\x18\xac\xc6\xed\xad\x69\x2f\x16\x8e\xda\x17\x55\x61\xc0\x88\xc1\x39\x9f\x4e\x21\x4a\xcc\x7a\x45\xca\x5d\xd7\xed\xd9\x0d\xb4\x44\x4d\x27\x0b\x94\x15\xb2\x29\x8e\xda\x56\x71\x54\x10\xb1\x14\x57\x99\x0c\x62\x4c\x86\xde\x3c\x6e\x45\xf9\x6a\xc7\x14\x8e\x72\xe5\xa8\xd2\x22\xf8\x22\x22\xb2\xd9\x77\x6e\x9c\xf2\x62\xc0\xab\xcd\x83\x3e\xae\x1f\x74\xcf\x75\xd5\x13\x29\xe9\x89\x0c\x42\xc0\xb5\x55\x54\x25\x07\x17\x18\xf0\xbc\x3a\xc7\x66\x6f\x1c\x0c\xc3\xd2\x83\x8a\x62\xca\x63\xcc\x69\x57\x89\x43\xbf\xde\x6e\x70\x55\x2c\xeb\x07\x43\x3c\xed\xd3\x69\x79\x84\x86\xad\x98\x7c\xe5\xa9\x24\x17\xa5\xa3\x19\xf9\xf9\xe5\xf3\xdf\x0e\x29\x9f\xed\x57\x6d\x0d\xaf\x3d\xeb\xf8\xb9\xb4\xd1\x95\x2d\xae\x39\x6f\xd5\xb6\x5e\xf6\x69\x16\x5d\x6a\xbc\xde\x27\x04\xca\x18\xe0\xaa\x44\x2f\xe5\x28\xbd\x90\x65\x5d\xed\xc9\xcd\x66\x85\x67\x8c\xef\xbb\xe4\xe5\x6c\x3a\x64\xe9\x43\x75\x14\x02\x81\x82\x4a\x23\xd5\xcf\x14\xd8\x99\xaf\xca\xf2\x31\x71\x47\xc7\xb0\x36\x71\x37\xce\x43\xc7\x09\xd8\xf6\xc2\x18\xdc\xd5\x8f\xaf\x5f\x3c\x87\x91\x98\xcf\x22\xb0\xc8\x60\x2b\x91\xfd\x36\xc9\x67\xc9\x36\x3b\x4a\x54\xab\x42\xfb\x52\x66\x06\x21\x2d\x93\xa1\xd2\xd6\x69\x57\x39\xeb\x82\xf8\x69\x19\xdf\x1b\xcc\xb2\xae\xd8\xdb\xcd\x78\x65\x77\x8d\x39\x47\x85\xd7\x06\x38\xb4\x12\x2e\xab\x40\xce\x55\xe1\x7a\x09\xa4\x62\xfb\x32\x0c\x23\x7f\x05\x01\xbc\x62\x82\xf8\x46\xb0\x37\x60\x1b\xaa\xce\x31\x24\x62\x34\x92\x29\x71\x7d\x3c\x96\xbb\x09\xc9\x0e\xaa\x25\xf0\xd2\x77\x8a\x50\x87\x5c\x28\x69\x13\xd0\xa7\x2b\xaa\xd2\xca\x8e\x4e\x28\xed\xa0\x05\x86\xbb\x26\x72\xd1\xbb\x02\x9c\x4a\xc4\xcc\x6e\x11\x88\x3b\xdf\x4e\x22\x1d\x00\x04\x76\x24\xfc\xc8\x37\x1c\xb4\x75\xb0\xb0\x02\xfe\xbb\x0d\x87\x76\x9a\x6f\x45\x1b\x7b\x4f\x60\x92\x55\x21\x45\xe4\xe4\xfb\x8e\x15\x8e\xa3\x88\x1d\xf0\xa0\x71\xbc\x44\xae\x9c\x03\x1b\xf5\x53\x7a\xe7\xde\xb8\x3c\x8e\x20\x12\x83\xc4\x59\x86\x95\x19\x09\xc2\x0c\xa7\x1a\x54\xb2\x65\x39\xe6\x3a\x14\xfe\x84\x48\x0d\xd0\xd1\x54\x25\x6f\xde\x85\xf4\xc7\xbb\x22\x6d\x09\x73\xa3\xc4\x94\x42\xec\x2a\x33\x69\x62\xbe\x88\xc8\xee\x86\xf2\x2e\x12\x7b\xc1\x03\xb6\xab\xbc\xc0\xf4\x7d\x19\x0d\xbb\x13\xc2\xbb\x48\xeb\x97\x57\x2f\x76\xd7\xae\x64\x0a\x26\x4a\xf3\x98\xf3\x6c\x02\x07\x40\xab\x28\x1c\x44\x9d\x0e\x43\x6a\x17\x89\xbc\xd2\xa7\x3e\xbb\x4a\x85\xc6\x63\x3c\xdf\xfb\x2c\x62\x39\x14\xad\x5d\xe4\xf2\x9a\x1c\xfd\xe7\xdf\xe7\xc7\xbb\x8a\x05\xbd\x9f\xdb\x5e\xe6\x35\x9b\x8d\x20\xd1\xea\x3b\xe7\xce\xdd\x88\xe8\x2e\xe8\x7e\x2e\x71\x9d\x7c\x21\x71\x1d\x98\xee\xe7\x12\x57\xe7\x0b\x89\xeb\xc0\x74\x3f\x97\xb8\x1e\x7d\x21\x71\x1d\x98\x6e\x39\x24\x2d\xa2\x50\x78\x28\x42\xe1\x3c\xff\x7f\x4e\x33\x41\x14\x3f\x64\x96\x04\x10\x8e\x77\xf3\x77\xe5\x63\x39\xcd\xb4\xea\xa4\x98\x90\x47\x72\x56\x8a\x59\x64\x98\x18\x89\xef\x9c\x52\x96\x33\xab\x15\x29\xa5\xec\xd4\x85\x84\x20\x85\x1c\x61\xeb\x3c\xee\x9e\xeb\x3e\x38\xdc\xbf\x66\x03\x53\x13\xc5\x2e\xc1\x7c\x03\xd2\xe9\x4c\x97\x07\xd6\x64\x8c\x6a\x91\x5c\x7f\x1a\xac\xb8\xa8\x5c\x9d\x3a\xaa\x32\x18\x9c\xba\x49\xe3\x14\x44\x5d\xe8\x2b\x99\x14\xd7\x46\x32\x29\x5b\x2c\x65\x8a\xd8\x6b\xf9\x3a\xb0\xd2\xe9\x8d\xd2\xc1\xb7\x9e\x4c\x1d\x47\xb3\xa8\x3e\xa3\xbc\x7f\xda\x79\x7c\x36\x3a\x2d\x92\xd3\xfb\xa3\xd1\xe8\xc9\x5a\xfc\xf6\xe1\xb8\x35\x11\x0b\x0c\xb3\xb9\x74\xd7\x27\xbb\x56\xac\x67\xe1\x58\x1e\x6e\x67\xbc\x09\xc2\x69\x47\xe5\xd2\x95\x21\x23\x58\x01\xf7\x92\xcd\x87\x1c\x73\xc6\x11\x07\xe9\x2b\x69\xae\x4f\xa2\x4b\x28\x4b\xf9\x39\xb9\xa7\x2a\x7f\xa8\xba\x95\x7d\xd2\xac\x10\x54\xe2\xad\x23\xb9\x2d\xc5\x32\x85\x95\xb4\x17\x07\x10\x7d\xe5\x72\xd7\x9c\x31\xc9\xc6\xed\xad\xbc\x95\x58\x2c\x8a\xb3\xa7\xb5\x07\x64\xe5\xab\x02\xad\xd0\x66\x6b\xf8\x22\x72\x6f\x6f\xe5\x4e\x5a\x2c\xa0\x05\x7b\x04\x7e\xab\x4e\x0e\x81\xb9\x9b\xa2\x99\xdb\x5b\xd9\xc0\x73\xe8\x69\x18\xe3\x33\xfc\x92\x4f\xf4\x46\x3e\xd1\x1b\x7c\x32\x16\xb4\xef\xbc\x79\x63\xda\x6f\xdf\x3a\x84\x05\xa1\xb2\x7e\xbd\x52\x7d\x59\x89\xb5\xd2\x65\x33\x22\xd0\x3b\xe2\xed\x6a\xb4\xf9\x65\x89\xda\x32\x0e\x01\x81\x0a\x9a\xc8\xb9\x5d\xb0\x38\xd0\x27\xe3\xce\xa6\xf9\xea\x5b\xef\xda\xfa\x9a\x8a\x7d\xd5\x67\x79\xb6\x65\x5d\x67\x58\x49\xd5\x02\x6d\x63\x56\x65\x9f\x06\xf4\xee\x12\x52\x02\x99\xb9\x9b\x51\x0d\x75\xaa\xf7\x1d\xe7\x11\xa3\xb1\xea\xd3\xd0\x67\x7b\x23\x1a\x65\x4c\x81\xd4\xf9\x5e\x43\xc9\x73\xbf\xb1\xfa\x85\x3e\x3f\x54\x2c\xe3\x55\x63\x09\x40\x6f\x4a\x80\x85\xa9\x23\xb2\x96\xc2\x3a\x5f\x34\x47\x8b\x78\x2a\x98\x15\x3d\x8e\xd4\xf2\x74\xb1\x98\xdc\x93\xcd\x87\xb2\x00\x4b\x01\x70\xc1\x8e\x3f\xe9\xfc\x50\xaf\x85\x1a\xf2\xa7\x43\xf8\xd3\x21\xfc\xc1\x1d\xc2\x96\x86\xbf\xb6\x94\x67\xa3\xe1\xb4\xbd\xcd\x5d\x59\x45\xbd\x13\xd7\x1b\xc6\x92\x19\x94\x2c\x55\x60\xda\x5a\xda\xb0\x2d\xed\x0f\x5e\x6d\xc0\x68\xd2\x37\x86\x51\x93\xb4\x0c\x90\x6d\x30\x1d\xf0\x21\x8e\x81\xf0\x4b\x86\x85\x91\x60\xc1\x64\x53\x83\xa5\xb9\x22\xc4\x18\x2c\x6d\x48\x9f\xe4\xd6\xce\x9f\xd0\x18\x5f\x28\xbb\x77\xf4\xf7\x8b\x9f\x5f\x7a\x99\xe4\x3b\x1c\xcd\x8f\xe0\xc5\xf1\xf6\x26\x6e\xdd\x1d\x09\x72\xfa\x99\x0d\xdc\xe1\x2f\x34\x2a\xaa\xbf\xf6\xa2\xa2\xf0\xed\xe4\x5c\xa7\x05\x9f\x7e\x67\x51\x73\x91\x30\xe8\x2d\xe9\xaf\xb9\x3d\xcd\x39\x70\x50\xa9\x60\x39\x42\x30\x23\x13\x8e\xc5\xc8\x52\x6d\x81\x5b\x80\x90\x1f\x11\xa2\x2f\x23\x4b\x78\x6a\x72\xd8\xad\x69\x25\x1c\xa6\x0b\xff\x34\x21\x3c\x71\x26\xe6\xd8\x70\x23\xb1\x2d\x92\x55\x73\x25\xde\x2b\xc7\x34\x2b\xa2\x2b\x7d\xe4\xae\x2d\xce\x0a\x96\x75\x27\x1d\x65\xb6\x75\x74\xd9\xd6\x71\xe5\x83\x13\xa7\x98\x4c\x79\x7d\xd5\xc0\xf2\xa4\x00\xa3\x5a\x5f\xc3\xb0\xb9\x81\xdf\x92\xe1\x64\x5a\x84\x9a\x2b\x19\xc6\x4e\x2b\xd8\xed\x80\x57\xc8\x19\xc6\x13\xec\x15\xfc\x6d\xc9\x8f\x3c\x4d\x75\x4d\xb5\xdf\x66\xd6\xaa\xfd\x2b\x5c\xba\xdf\xe6\x62\xfd\xb6\x60\xd3\xe8\x07\x39\x0a\x18\x9e\xe2\x67\xc7\xab\x84\xfa\xa9\x67\x0c\xa4\x6a\x9a\x0e\x7a\xc2\xb0\xe6\xa6\xd8\x5c\xf2\x6f\x2e\x53\x5e\x55\xd3\x79\x87\xe6\xed\x70\x21\x1f\x20\x53\x9f\x2b\x2c\x95\xb2\x6e\x9c\x56\x3e\x2e\xaf\xd3\x0e\xe5\x15\xa3\xbb\xa9\x5c\x3b\xc7\x60\xbe\xbb\x5a\x27\xc5\xbc\xc0\xfa\x24\xb9\x21\x19\x07\xb7\x22\x8b\x4b\x15\xe7\xb5\xdf\x2d\x5a\x53\xd1\x65\xdd\xeb\x38\x29\xf3\x20\x49\x2d\x53\xda\xc5\x9d\x64\xe4\x22\xd7\x9d\x3d\xdc\x48\x51\x3a\xd2\x9b\x9c\x0e\x7e\x92\x32\x56\x08\xc9\x0f\x61\xc4\x7a\x2d\x80\xea\xf7\xb2\x7c\x3b\x2f\xc0\x4b\x22\x4e\x03\xad\xb7\x2e\x7e\x03\x23\xbf\xbf\x2c\x55\xee\x95\x8b\xfa\xc8\xaa\xe0\x0f\x7b\xab\xc5\x55\xf5\xdd\x88\x4c\x67\xc3\xd0\x92\x2c\x39\xe6\xa3\x19\x45\x55\xf5\x00\x09\xb3\x7e\xa7\x0d\x96\xa0\xf8\x72\x65\xa7\x60\xd2\x8f\x42\xff\xb2\xef\xbc\xdb\x98\x5a\x37\xca\x45\x36\x79\xd9\x78\x19\x5e\xb8\xa0\xb2\x50\x8b\xd2\xb7\xaa\x2c\x81\xa1\xfc\xcb\x55\xa7\x2a\x5a\x5f\x0f\xab\xc8\xb5\xa2\xca\xfa\x33\x1f\x69\x98\xf5\x00\x69\x40\x5d\x50\xd1\xac\x7f\x4a\xf4\x2f\x1a\x5d\xd3\x39\xa0\x02\xcc\xc2\x95\x63\x08\xec\x21\x9f\x4d\x78\x04\x2a\xd8\x77\xc6\x4c\x90\xb6\x87\x15\x8a\x4e\x7e\xfc\x61\x68\x54\x48\x6f\x94\xab\x16\xab\x2c\xa4\x91\x53\x31\x02\xd0\xd5\x3a\x1b\x04\xbd\x3b\x15\xfc\xce\x5b\x2c\x11\xf9\x05\xa1\xb5\x34\xea\xd7\xef\xb0\xbe\xa3\x30\xe8\x2b\x3c\x88\x0e\xcb\x75\xcd\x93\xce\x08\x4c\xcd\x93\xe3\xd8\x45\x4f\x5a\x41\x57\xa4\x05\xb0\xe7\xf1\x53\x25\x19\xc8\xff\x9f\x97\x6f\x99\x22\x92\x57\x8b\xa0\x36\x35\xee\xe9\x23\x1c\x21\x3b\x66\x1e\x96\x9f\x1e\xb5\xd7\x84\xf2\x35\x87\xe4\x95\x93\x6d\x29\x76\xc2\xf0\xc3\x9d\x6c\x5d\x9d\xd3\xca\xa3\x6f\xbb\x36\x15\xbf\x93\x2c\xc4\x6d\x7c\x9f\x2e\x5a\x93\x5f\x1d\x5b\xe6\xf6\x4c\x3b\x0e\x65\xb0\x3b\xc6\x8c\xe6\xb0\x0d\x95\x53\xeb\x53\x68\x59\xe5\xaf\xbb\x5b\x75\x56\xf5\xb8\x16\xcd\xdc\x60\xf7\x4a\xcc\x56\xe6\xa8\xa4\x54\x6c\x31\xd5\x0b\xa5\x6c\x4f\xda\x68\x98\xad\x60\x55\xfd\x52\xd2\xc2\x05\x5b\x52\xad\x6d\x34\x4b\xfe\xd0\x8e\x54\x7e\x73\x65\x29\x17\x1f\xbe\xd7\xfa\x85\xea\x95\xb2\x0f\xa0\x5d\x95\xf4\x0f\xba\xc0\xae\xfb\xa0\xd4\x46\x2a\x61\x8a\x55\x77\x08\x86\x56\x0e\x8d\x31\xb8\xb7\x8a\xf1\xe4\x9f\x2a\xb8\x90\xf5\x7e\x3c\x3d\xfa\xea\xfe\x92\x64\xbe\x52\x18\x71\x5c\xa9\x28\xcf\xe9\x4d\x3a\x03\xc9\x27\x18\xd1\xce\x40\x1a\x58\xe7\x01\x8c\x79\xe0\xa8\x6d\x8b\xb6\x57\xfe\x87\xb7\xbf\x30\xa0\x93\x09\xbb\x27\xf0\x6a\x7a\x3a\x05\x05\x90\x5b\xac\xbe\x67\x43\x71\xa0\xe2\xaf\xd5\xfb\xbb\xba\x1c\xd2\xb5\xab\xf5\x97\x43\xb5\x35\xae\xb5\x01\xe2\x1f\x22\xf8\xf3\xc0\x16\x89\xe2\x53\x37\xbd\x91\x4f\x75\xdc\x64\xce\x03\x4d\x1c\x25\x2b\x04\xd1\xcd\x43\x38\xa6\x3e\x11\x91\x01\x5f\xfe\xd5\x93\xd7\x61\xd3\x5d\xe3\x25\x2b\x3d\x7b\xa1\x04\xfa\xe9\x61\xd3\x6b\x58\x0e\x96\x52\x31\x4b\x4b\x01\x13\xe2\x28\x0c\x40\x6c\x65\x4d\xa2\x18\x90\x5f\x10\x14\x92\xd1\xb7\xb7\x56\xd5\xae\x29\xce\x5d\xa2\xfc\x34\x1e\xcf\x22\xd8\x69\x49\x5e\x78\xb3\x0d\xf9\x22\xfd\xfa\x14\xda\xaa\xfc\x69\x1b\x7a\xaa\x12\x69\x6b\x5a\x87\xf7\xc6\x7a\xeb\xec\x9e\xcd\x35\xaf\xc3\x38\xe0\xd7\x1e\x8f\xd1\x87\x62\x79\xb0\xed\x82\x17\xbb\x78\x4a\x55\x70\xbe\x69\x93\xe7\x65\xe9\x7f\xe0\x3d\xee\x4f\x68\x2a\x36\x6f\x6f\xe8\x3a\x42\x5d\xc8\xbb\x5b\x1f\xc8\x3e\xae\x0c\x31\x28\xaa\x5f\x38\xee\x9a\x2b\xfd\xca\x86\xe7\x74\xba\xcf\x8e\xc7\x3f\x0c\x60\xb4\x5b\x2d\x12\x28\x31\x44\xc7\x80\x4c\xfe\xa1\x00\xfd\xd9\xff\xaf\xf2\x15\x01\xc5\x36\x9f\xfe\x1f\x54\x91\x15\x65\x67\x0f\xf5\xac\xe8\xe7\x7f\x03\x00\x00\xff\xff\xaa\x46\x18\x51\x60\x49\x00\x00")

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

	info := bindataFileInfo{name: "fcs-lpc-motor.html", size: 18784, mode: os.FileMode(420), modTime: time.Unix(1553008059, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x59\x6d\x73\xdb\xb8\xf1\x7f\x2d\x7d\x0a\x04\x2f\xce\xd2\x48\xa4\x7c\xff\x7f\xe7\xa6\x63\x89\xee\xf8\x1c\x67\x2e\x6d\xce\xc9\xc4\xce\x74\xda\x24\x93\x81\x48\x50\x42\x0c\x12\x2c\x01\x4a\xd1\xe5\xf4\x9d\xfa\x19\xfa\xc9\xba\x0b\x80\x14\x29\x29\x7e\xb8\xe6\x85\x4d\x02\x58\xec\x2e\x7e\xfb\x08\x6a\xf6\xec\xf9\xeb\xcb\xdb\x7f\xbc\xb9\x22\x4b\x93\xc9\xf3\xfe\xec\x59\x10\x90\x4b\x55\x6c\x4a\xb1\x58\x1a\xf2\x7f\xa7\x3f\xfe\x44\x6e\x97\x9c\x2c\x54\x20\xb5\x36\xe4\xa2\x32\x4b\x55\xea\x90\x5c\x48\x49\x2c\x8d\x26\x25\xd7\xbc\x5c\xf1\x24\xec\x13\x02\xbb\xdf\x69\x4e\x54\x4a\xcc\x52\x68\xa2\x55\x55\xc6\x9c\xc4\x2a\xe1\x04\x86\x0b\xb5\xe2\x65\xce\x13\x32\xdf\x10\x46\x7e\xbe\x79\x1e\x68\xb3\x91\xdc\xed\x93\x22\xe6\x39\xec\x35\x4b\x66\x48\xcc\x72\x32\xe7\x24\x55\x55\x9e\x10\x91\xc3\x24\x27\xaf\x5e\x5e\x5e\x5d\xdf\x5c\x91\x54\x48\xee\x64\x81\xbe\x4e\x6d\x42\x66\x4b\xce\x12\x7c\x81\xd7\x8c\x1b\x46\x72\x96\xf1\x88\xae\x04\x5f\x17\xaa\x34\x14\x74\xc8\x0d\xcf\x4d\x44\xd7\x22\x31\xcb\x28\xe1\x2b\x90\x17\xd8\xc1\x98\x64\x22\x17\x59\x95\x05\x3a\x66\x92\x47\x3f\x86\xa7\x63\x10\x2a\x8c\x60\xb2\x3d\x55\xc1\x39\xed\x98\xcd\x61\x6a\xc3\x35\x6d\x0b\x8c\x97\xac\xd4\x1c\x04\x54\x26\x0d\xfe\x5c\x2f\x19\x61\x24\x3f\x7f\x71\x79\x43\x5e\xbd\xb9\x24\xb7\x5c\x9b\x39\xcf\xe3\xe5\x6c\xe2\x16\x1c\x91\x8e\x4b\x51\x18\xa2\xcb\x38\xa2\x73\xb5\xe6\xe5\xa7\x58\x65\x85\xca\x41\x5f\x3d\x59\xf3\xf9\x6e\xf4\x79\x6f\x1c\x48\x61\x78\x08\xea\x87\x9f\x41\x9b\xd9\xc4\x71\xf2\x6c\xa5\xc8\xef\xc0\x3a\x32\xa2\x22\x73\x20\x2c\x4b\x9e\x1e\x11\x51\xb0\x02\x4f\x86\xb6\xe8\x0e\x82\x58\x32\xad\xb9\x0e\x11\x66\xfa\x00\xdb\x34\x06\x75\x8a\x38\xc8\x94\x51\x65\x67\x87\x65\xe6\xde\x89\x75\x34\xf2\xd5\x0f\x08\x41\x97\x48\xa5\x5a\x07\x9b\x33\xc2\x2a\xa3\xa6\x7e\x65\xeb\x9f\x73\x95\x6c\x5a\xe4\x29\x98\x31\x48\x59\x26\x24\xd0\x9f\xbc\x55\x73\x10\x76\x32\x26\x27\xbf\x70\xb9\xe2\x46\xc4\x8c\x5c\xf3\x8a\xc3\x4c\x33\x31\x26\x17\x25\x18\x72\x4c\x34\xcb\x75\x00\x26\x14\xe9\xb4\xcb\x6e\xcd\xd1\x91\xcf\xc8\xff\x9f\x9e\x76\xa5\x03\x9e\xb5\xe6\xb3\x49\xed\x60\xb3\x1a\x63\x0a\xfe\x40\xb4\x29\x45\x6c\xe8\xb4\xbf\x62\x25\x01\x04\x48\x04\xca\xf6\xd0\x53\xce\x08\xbd\x9e\x5c\xd0\x71\xbf\x67\xd4\x1d\xcf\x61\x88\xef\x2c\x49\x70\xe5\xeb\xd7\xf0\x02\xde\xb6\x5b\x9c\x03\x8c\xcd\xa7\xaa\x48\x98\xe1\xbb\x4d\x25\x88\x83\x23\x02\xb3\xde\x97\x33\x92\x32\xa9\x39\xcc\xf6\x7e\xdb\xbd\x6f\xe1\x0f\x9c\x2e\x77\x44\x71\x96\xe8\x33\x92\x57\x52\x22\x19\xf0\x62\xbb\x11\xba\x0c\xcb\x9a\x31\x6c\xdc\x4e\xfb\xfd\x41\x5a\xe5\xb1\x11\x2a\x1f\x0c\x91\x03\x1e\xc0\x6a\x0a\x47\x58\x70\x73\xa9\xd4\x9d\xe0\x03\x0a\xae\xfb\xe9\xf6\xf5\xdf\xae\xae\xe9\x70\xea\x88\xf0\x70\x87\x34\xef\x6e\xae\xde\x5a\x12\x00\x21\xf4\x24\xf8\xf0\x33\x35\x67\xfb\x9c\xf6\xb7\x83\xe1\xb0\xdf\xb7\xa4\x85\x54\x2c\xb9\xb1\x98\xbe\x80\xc8\x06\x9a\x46\x2f\x8c\xf4\x46\x37\x04\xc4\x72\xcd\xf9\x9a\x20\xe5\x5b\x3b\x31\x40\x99\x6e\x2d\x54\x39\xf2\xe2\x90\x32\x5a\x4c\x2c\x03\xcb\x01\x20\x72\xe6\xe9\xf5\xcc\xa6\x00\xac\x09\x8d\x8d\x44\xb0\x7b\x3d\xeb\xb4\x80\xfe\x17\x37\xf4\x26\x6b\x34\xb7\x93\x98\x52\x70\x93\x53\x39\x80\x00\xca\x58\x9e\x68\xb7\xc3\xe1\xef\xb5\x0c\x21\x27\x56\xd2\xe0\xc2\x16\xfe\x90\x0d\x5a\x2a\x44\xa2\x50\x83\x82\x83\xbf\xde\xbc\xbe\x0e\xd1\x79\xf2\x85\x48\x37\x03\x58\x18\xe2\x41\xb6\xcd\x59\xf0\x71\xa1\x7f\x16\x39\x2b\x37\x37\x96\xce\xc1\x01\xd8\xb5\x81\xbb\xf4\x4a\xb4\x4f\xcc\x57\x10\xd5\x0d\x6e\xa9\x2a\x33\x58\x7d\xa3\xe4\x26\x03\xbe\x89\xca\x3c\x41\x28\x15\xe4\xb2\x5b\x56\x82\x25\xc3\x82\x95\x30\x77\x25\x79\x06\x0f\x70\x8e\x3d\xc0\x1c\x5e\x0d\x5c\x5d\xb4\x8e\x81\xe5\xb0\x3a\x0a\x95\x43\x0a\xd5\x0a\x0d\xff\x62\x40\x30\x0b\x57\x4c\x56\xd6\xa7\xfb\x8f\x07\xcb\xe3\x80\xe5\xc7\xdc\x0f\xc3\x53\xce\xee\xf8\x0d\x76\xfc\x51\x03\xcf\xfe\x38\xf7\xc9\x04\x8b\x8b\x56\x50\x95\xa4\x5a\x0c\x28\x6e\x00\x4d\x89\x3f\x74\x18\x86\x64\x80\x73\xbc\x8c\xe8\xc8\x6e\x0b\xdd\x70\x44\x87\x4d\x50\x59\x44\x81\xbf\x5b\xb7\x23\xbf\x82\xa1\xec\xad\xe0\x61\xb7\x8f\xf1\xa1\x55\xbe\x6d\x07\xc7\x15\xdf\x11\x63\x60\x8c\x4c\x1d\xe8\xc0\xfa\xba\xca\xe6\x10\x48\x89\x8a\x2b\x8b\x00\x60\xe2\xc1\xf8\x79\xf3\x32\x19\xd8\xc4\x0e\x32\x02\x3a\xb2\x92\x47\x34\xa8\xcf\x81\x1c\x61\xa8\xa5\x80\xe3\xd0\xa1\xe3\x88\x47\x7a\x00\x13\x50\x7a\xb4\x67\x54\xd4\x08\xad\xfa\x18\x07\xa8\x69\x1b\x0f\x48\xc1\x68\xcb\x5f\x55\xde\x36\x90\x25\xaa\x83\x00\xdc\xd0\xa0\x73\xbc\xa7\x86\x67\x50\xdf\x98\xa9\x4a\x4e\xc7\x84\x16\x4a\x0b\x24\xc7\xf7\xb2\xc8\x34\xfd\xe8\x61\x87\xf6\xc5\x28\xdc\x62\xa1\x72\x23\xcc\x65\x60\xa5\x01\xae\x0b\x58\x3a\x9d\xc2\x63\xe6\x98\x87\x92\xe7\x0b\xb3\x84\x99\xd1\x68\x97\x72\x04\xfa\x8c\x5d\x7f\x2f\x90\x73\xa3\x0c\x32\xbe\x0f\xef\x4c\xe5\x80\xb2\x95\xbd\x03\x5d\x24\x88\x6d\x0f\xf7\x87\x22\xcf\x79\xf9\xcb\xed\xaf\xaf\x80\x93\xd3\xee\xbd\x48\x3e\xda\x14\xd2\x85\xe5\xef\x36\xf9\x7f\x13\x99\x1c\x7b\xb3\x07\x94\x71\xf5\x03\x7a\x02\x08\xd5\xcc\xba\x2c\xee\x0a\x99\xc4\x63\xd0\x6b\x38\xcd\xc5\x8a\x09\xdb\x16\xd1\x7a\x11\x9a\x19\x5c\xb4\x85\x48\x64\x6c\xc1\x27\x9f\x0b\xbe\x98\xce\x99\xe6\x3f\xfd\x69\xec\x8e\xe6\x2c\xe8\xd5\x6a\x15\x94\x18\xfd\x6a\xe8\xab\xbe\x55\x12\xc6\xc0\xcd\xce\x93\x11\xa1\x11\x9d\x36\x6b\x31\x6b\xab\x1f\x5b\x0e\xa1\x2e\xa0\x31\x1a\x9c\x4c\x4f\x86\x8e\x10\xcc\xe6\xac\x16\xa1\xcd\x66\x31\xdb\x37\x57\xdd\x10\x58\x8e\x28\x8a\x59\x83\xd5\xd3\xeb\x25\x96\xa6\x41\x8c\x8e\x59\x5e\x98\xc1\xe9\x30\x8a\x4e\xc8\xc9\xd0\xd1\x86\xba\x9a\x3b\xef\x1c\xfc\x38\xdc\x6d\x12\x29\xee\x10\x10\xea\x5f\x5e\xa7\x03\x77\xa6\x08\xbc\x66\x08\x55\x02\xfc\x2f\xef\x6c\xc4\x65\xaf\xd4\x38\xf6\x2f\x9e\x97\x6b\x46\xfc\x1e\x4a\x2d\x6a\x6b\x60\xab\xd6\xbe\xe2\x1d\x94\xbb\x26\x84\x7c\xf2\xc0\x92\x09\x7e\x70\xa3\xe2\x3b\xc8\x6d\x74\xad\xcf\x26\x13\x0a\x40\x22\x1d\x76\x24\x88\xe9\x04\x69\x9b\x0a\xde\x04\xe0\x63\x77\x23\x6d\x77\xf7\xba\x76\xbc\x47\xed\x5f\x41\x0a\x51\xb4\x9b\x00\xac\xfb\xab\x5c\x15\x3c\xdf\x3b\xe2\xf6\x90\x2a\x96\x4a\xf3\x03\xb2\x43\xba\x8c\x6b\x0d\xee\x78\x3c\x97\x5b\xea\x56\x3f\xe6\xb5\x7f\x0e\xaf\xb6\xd1\x68\x67\x64\x9b\x94\x0a\x6c\xfa\x1d\x03\x2b\xa0\x21\xaa\x13\xfa\x2e\x84\xeb\x95\x86\xf3\xbd\x31\xa7\x0d\xe4\x28\x1d\x38\x62\x0b\x6c\x0f\xfd\xa9\xde\x1c\xd9\x66\xce\xe9\x8c\xd9\x36\x63\x70\x4b\x2a\xe0\x5c\x63\xc2\x52\x03\x7d\x52\x5c\x72\xe8\x84\xf0\x0a\xa3\xc9\x0a\x9b\x5e\x68\x85\xf1\xb0\x63\x88\x78\x83\xf7\xa9\x92\x93\x0d\x14\xbf\x30\x44\x0e\xce\xb9\xa6\xb6\x5f\x81\x7f\x4e\xca\x2e\xc7\x44\x7b\xb0\x40\x7d\x79\x77\x7b\xe9\xdb\x12\xab\x5b\x37\x0f\x0f\x1a\x28\x0e\x12\x91\x5d\xf2\xbe\x61\x29\xf6\x4b\x85\xcd\x18\xdf\x2c\x0f\xbd\x9e\x47\xd1\xf6\xc7\x35\xbe\x0e\x2c\x44\x85\xda\x79\xcc\x42\xbd\xc7\xc0\xbb\x2b\x69\x10\x49\x22\x07\xa0\x43\x06\x3e\xb1\xe2\x35\x67\x37\xfd\x74\x76\x4e\x8d\x16\x37\x3b\xf1\x07\xf8\x60\x4d\x1a\x76\xb2\xbd\xd5\x0b\xe7\x9f\xce\x8d\xe5\x0b\xc9\x8f\xb1\xb3\x0b\x18\x87\xe4\x3f\xff\xb6\xe0\xed\xea\x9c\xcb\x98\x96\x0a\x6b\x67\x53\xe8\x9a\x3a\xf7\x44\x1d\x90\x09\xd6\xb2\x23\x5a\x58\xfe\x90\x77\x47\xa8\xc6\x25\x6d\xdc\x11\x1d\xbf\x65\x8d\x3f\x24\xb6\xb1\x6f\x5b\x2a\x7d\xed\x66\xa7\x4f\xe7\x57\x1b\xf8\xe0\x10\x8e\xd2\xea\x4e\x38\xdc\xd4\xbe\xa3\xb2\x69\xfa\x3d\xb5\xb5\x37\xcd\x06\xe3\xa7\xf1\xcb\xa0\xc4\x1f\x3b\x3c\xce\x3f\xdd\x2d\x3d\xb7\x83\x58\x69\xf2\x07\x0c\xdf\x5b\xe2\x8f\xed\x48\x72\x0d\x96\x6d\xeb\x10\x91\x23\x7d\x2a\x44\x89\xbb\x9d\x1d\x59\xb3\x2e\x1f\xec\xfa\x3f\x20\x6b\x3a\xb4\x44\x68\x6c\x66\x92\xf6\xb1\xc8\xb3\xc8\xb6\x8b\x94\xfc\xfe\x3b\x79\xd6\x51\xa2\x64\x42\xef\x11\x47\x35\xf1\x0f\x3f\xd4\x0a\x77\xa2\x6a\xd7\x3d\xc2\x09\x0e\x7b\xc7\x5d\xf3\x08\xcb\xbe\x75\xfc\x26\xa8\x22\x69\xf7\xe0\x6d\xdd\xfd\xeb\x23\x76\x57\xf3\x4c\x98\xff\x79\x77\x03\x84\x7b\x71\xce\xd5\xdb\x1e\x34\xf5\x0f\xd6\x74\x4f\x75\xbc\xa6\x1f\x90\xdd\x5f\xd2\x11\x4a\x35\xff\xfc\x60\xc5\xb6\x70\x03\x61\x28\xac\xe2\x9d\xa2\x54\xf2\x98\x83\x6f\x26\x78\x25\x3e\x72\x79\x81\x5d\xc3\xa6\x44\xc3\xe0\x3d\xe5\x65\x49\x3f\x5a\x97\xa1\xfb\xe9\xea\x5f\x15\x87\x2b\x3d\x97\x3c\x06\x47\xac\xbb\x7d\x78\xd5\x01\x93\xbc\xf4\x46\x28\x24\xdb\x5c\x95\x25\x10\x20\x6f\x0f\x64\x4f\xaf\x85\x89\x97\x5e\x04\x84\xd1\x82\xd3\x8f\x9e\x7d\x0c\x8d\x35\xa1\x0b\x9e\x07\x89\x82\x1c\x71\x86\x73\xdf\x0e\x43\x2d\xa0\x9d\x2f\x6c\xf8\xb6\x03\xcf\x7e\x5a\xb2\xf6\x3e\x72\x92\x56\xc3\x71\x1f\xe7\x3c\xad\x34\x58\x20\x40\x55\xe8\xa8\x1e\x7e\x82\xbb\x4b\x27\x5d\x1c\x22\xe8\x45\x0d\x9d\xfc\x6d\x57\x8b\x4c\x2f\x8e\x68\x61\xbf\x59\xf1\x2f\xf6\x1a\x32\x2b\x4a\xfc\x1c\xb7\x91\x3c\xfa\x40\x71\x12\xe0\x14\x0b\xb8\x0f\x4b\x9e\x9a\x0f\xf4\xdc\xa5\x4d\xdf\x9c\xe9\x85\xb7\x35\xbc\xf9\xcb\x02\xfd\x90\x53\x2f\xdb\xc5\x69\x2b\x46\x91\xaa\x1b\xa3\xc4\x11\xf6\xac\xf4\x51\x84\x14\x10\xa6\x58\x41\x67\xf3\xb2\x91\xe5\x0e\xd1\x10\xd1\xd9\x04\x94\x3c\xa7\xdf\x09\x42\x64\xdb\x06\x6b\x0e\x89\xe6\x6e\xba\x73\x06\xbc\x27\x3e\xce\x1b\x6a\x81\x76\xc7\x3d\x12\x9d\xdf\xad\x16\xd4\x25\xa4\x03\x89\xbf\x89\xa2\x2d\xd0\xa5\xd2\xd5\x7d\xad\x6e\x23\xba\xe4\xf8\x2d\x79\x4f\xf8\x74\xc7\x46\x76\x6e\x79\x20\xd7\x70\xcf\x08\x42\x48\x95\x59\x6d\xbb\x44\x86\x36\x90\x77\x9c\x5d\x76\x0a\x90\x28\xc0\x6b\x87\x0b\x6f\x4b\x99\x71\xb3\x54\x49\x84\x99\xda\xd0\x66\x96\xd9\x14\x12\xd1\xa5\x31\x05\x5c\x55\x9a\x6f\xb5\x13\xb8\x71\xd9\xfb\xd6\x5f\xf0\x4b\x1c\x7e\xcc\xc9\xf1\x97\x8d\x77\x6f\x5f\xa2\x8f\x86\xf8\x15\xbc\xa5\x44\xbb\xd2\xce\xdc\x47\xf5\x79\x65\x0c\x5c\x74\x7d\x96\xc4\xf4\x26\xe2\x3b\xf0\x56\xa7\xe1\x0b\x50\x70\x80\x97\xfc\x11\x1d\x7e\xa0\xc4\x7e\x7a\x87\xc5\x58\x49\x55\xa6\x95\x84\x29\x91\x20\x71\x7d\xae\x5a\x9b\x2e\x66\x23\x0a\xbe\xfe\xdc\x2f\x91\x7f\xbe\x7c\x83\xfe\x38\xb8\x11\x59\x25\x6d\xff\x4f\x02\xe2\x65\xcc\x26\x6d\xad\x6a\x9f\x05\x7b\x85\xac\x28\xf0\xa3\x17\xdc\x79\x93\x41\x22\x87\x53\xd2\x35\xf6\x7e\x4a\x77\xcd\xfc\xc3\x49\xbd\xa1\x7b\xf0\xaa\xd6\x50\xde\x9b\xd9\x61\x0b\xb0\xdf\xfd\xe8\x01\xd7\x64\x18\xda\x1f\x0c\xaa\x1c\x6e\x1c\x4a\xda\x8c\x6d\x71\xa4\x90\x4f\x55\x65\xf0\x2a\x84\x3f\x0a\x48\x12\x03\x17\x38\xbc\x7b\xd0\x73\xac\xce\x33\xf4\x55\xc0\xd8\xa6\x64\xc0\x00\x66\xe1\xd8\xb3\xce\xef\x1a\x38\x7f\x3e\x9b\x1c\xce\xe1\xfe\x09\x30\x40\x4e\xb3\x09\x2a\x71\xde\x9f\x4d\xdc\xcf\x53\xff\x0d\x00\x00\xff\xff\x22\x34\xab\x27\x66\x1b\x00\x00")

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

	info := bindataFileInfo{name: "index.html", size: 7014, mode: os.FileMode(420), modTime: time.Unix(1508441623, 0)}
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
