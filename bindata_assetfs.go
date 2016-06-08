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

	info := bindataFileInfo{name: "bower.json", size: 1606, mode: os.FileMode(420), modTime: time.Unix(1461078977, 0)}
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

	info := bindataFileInfo{name: "favicon.ico", size: 1150, mode: os.FileMode(420), modTime: time.Unix(1455561960, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fcsLpcMotorHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x5b\x7b\x73\xdb\xb6\xb2\xff\x5b\xfa\x14\x08\x73\xef\xc4\x89\x43\x4a\x56\xec\xdb\x5c\x55\xd2\x39\xa9\xe3\xce\xc9\x19\x27\xf1\xc4\x69\xcf\x74\x3a\x9d\x0c\x45\x42\x12\x63\x92\x60\x48\xd0\xb2\x9a\xf1\x77\x3f\xbb\x58\x80\x2f\x51\xb2\x2d\x3b\x7d\xcd\x34\x02\x17\xc0\x62\xb1\xd8\xc7\x0f\x0f\x8f\x1e\xbd\x7e\x7f\xfc\xf1\x97\xb3\x13\xb6\x90\x51\x38\xe9\x8e\x1e\xd9\x36\x3b\x16\xc9\x2a\x0d\xe6\x0b\xc9\x06\xfd\x83\xff\x63\x1f\x17\x9c\xcd\x85\x1d\x66\x99\x64\xaf\x72\xb9\x10\x69\xe6\xb0\x57\x61\xc8\x54\x9b\x8c\xa5\x3c\xe3\xe9\x25\xf7\x9d\x2e\x63\xd0\xfb\xa7\x8c\x33\x31\x63\x72\x11\x64\x2c\x13\x79\xea\x71\xe6\x09\x9f\x33\xf8\x9c\x8b\x4b\x9e\xc6\xdc\x67\xd3\x15\x73\xd9\x0f\xe7\xaf\xed\x4c\xae\x42\x4e\xfd\xc2\xc0\xe3\x31\xf4\x95\x0b\x57\x32\xcf\x8d\xd9\x94\xb3\x99\xc8\x63\x9f\x05\x31\x10\x39\x3b\x7d\x73\x7c\xf2\xee\xfc\x84\xcd\x82\x90\xd3\x58\x93\x6e\x77\x14\x06\xf1\x05\x88\x10\x8e\xad\x20\x4a\x44\x2a\x2d\xb6\x48\xf9\x6c\x6c\x4d\xc5\x92\xa7\x9f\x3c\x01\xc4\x98\xc7\x32\xeb\x25\x22\x5c\x45\x3c\x35\xbf\x0e\x4e\xd8\x42\x0e\x99\x97\x06\x89\x64\x59\xea\xb5\xf4\xfa\xfc\x25\xe7\xe9\xaa\xe7\x07\x99\xd4\x65\xe7\x73\x66\x4d\x46\x3d\xea\x75\x27\x09\x82\x54\xc4\xf6\x2c\xe4\x57\x76\xe8\xae\x44\x2e\xd7\x08\xb6\x17\xba\x59\xc6\xb3\x52\xb6\xbb\xf1\x76\x73\x29\xe6\xa9\x58\xda\x92\x5f\x49\x37\xe5\xee\x06\xb2\xe1\x7f\x57\xd1\x45\x1a\x95\xa5\x1d\x99\x04\x9e\x88\xab\xc5\x1d\xd9\x24\xee\x9c\x57\x8b\x3b\xa8\x2c\x71\x13\x9e\xda\xd3\x5c\x4a\x11\xd7\x3e\xee\x2e\x12\xf5\xf6\xdc\xd4\xaf\x14\x77\x65\xe3\x07\x6e\x28\xe6\xb5\x8f\x9d\x59\xa5\x2e\xd0\x41\x43\x31\x0f\x5b\x48\xbb\xb2\xc5\x85\xab\x2b\xae\x42\xd9\x99\x69\x9c\x80\x4b\x54\xca\x0f\xc1\x68\x77\x8b\xd7\xbc\x24\x8f\x2a\xc5\x5d\xd9\x44\xae\xe4\x29\xac\x64\xe3\x73\x57\x76\x10\x7b\x44\x18\xda\x0b\xee\xfa\x8d\xd5\x6d\xa9\xd9\x79\x90\x30\xf0\x31\x5c\x56\x3e\x76\x65\x25\xdd\x69\xb5\xb8\x33\x1b\x31\x9f\x87\xbc\x6e\x78\x35\xda\xee\x8c\x45\x38\x75\xd3\xfa\x57\x19\x51\x7c\x11\xd9\x91\xf0\xf3\x10\x92\x98\x3f\xb6\x66\x5e\x66\x87\x89\x07\x24\x29\x52\x1b\xbc\x33\x88\xa1\x59\x67\x04\xf6\x91\x84\xb0\xb2\x50\xee\x8c\x54\x62\x9b\x40\x96\x62\x6c\xb8\x10\x90\x39\xbf\xaa\x32\x63\x90\x49\xa0\xd5\x6a\xc8\xa6\xa1\xf0\x2e\xbe\xd7\xd4\xa9\xb8\xb2\xb3\xe0\xf7\x20\x9e\x43\x85\x48\x71\xf1\x80\x64\x6a\xd1\x8a\x6d\x37\x0c\xe6\xf1\x90\x41\x92\x04\xe3\x31\x35\x91\x9b\xc2\xf0\x43\x76\x94\x5c\x95\xa4\x2b\x7b\x19\xf8\x72\x31\x64\x83\xa3\x7e\x85\x1e\xc4\x05\xbd\x5f\xd0\xaf\xbb\x1d\xad\xbb\x58\x50\x36\x06\x49\x3b\x1d\x7b\xc9\xa7\x17\x01\x0c\x9a\x24\xdc\x4d\xdd\xd8\xe3\x43\x16\x83\xce\xbe\x87\x3a\x92\x6f\xc8\xfa\xf8\x91\xb8\xbe\xaf\xa4\x56\x5f\x53\xd7\xbb\x80\x4c\x03\x49\x7b\xc8\x24\x74\xcb\x12\x70\xbe\x58\x42\xd5\x35\x2a\xa5\xa7\xb5\x82\x65\x00\x1a\xf8\xa3\xd4\x67\x63\x4e\xc6\x9c\xc2\x20\x8a\x28\x8a\xd2\xf2\xa7\x85\x1b\xfb\x21\xff\x44\x2a\x66\x39\xc0\x8c\xd8\x8d\xf8\xd8\xba\x0c\xb2\x00\x74\x6f\xb1\x88\x03\x1c\x81\x25\x49\x40\xc3\x16\x73\x3d\x19\x08\xe8\xda\xd3\x1d\x12\x48\xa6\x4b\x10\xf6\x67\x50\x9d\xef\x62\xdd\x99\x2b\x41\x79\xd0\xc4\x79\x86\x4b\x06\xe3\x07\xd1\x9c\xb9\x97\x2e\x04\x09\x02\x00\x0b\x29\x93\x61\xaf\xb7\x5c\x2e\x1d\x04\x3c\x8e\x48\xe7\x3d\x18\x0c\x72\x8d\xcf\x67\x6e\x1e\xca\x1e\xc2\x8f\x0c\x87\x10\x9f\x0e\x9c\x24\x9e\xc3\xb8\xa1\x1c\x5b\xc4\xc4\xea\xa9\xd5\xef\xd5\xa7\x05\x33\x46\xea\x8f\xef\x3f\xbc\x05\xe3\x7c\xf3\xee\xec\xa7\x8f\x0c\x41\x17\x18\xa4\xd2\xbd\x05\xf3\x3e\x06\xfc\x73\x01\xe3\x03\xd0\x10\x00\x32\xe6\x62\xaf\xff\xd4\x62\x3f\xbf\x3a\xfd\x09\x9a\x7d\xe0\x33\x40\x59\x0b\x34\xed\x1e\x71\x81\x25\xb2\x49\x91\x4a\x6f\x5b\x15\xf1\x0d\x67\xda\x19\x55\x62\x2e\xa3\xd5\x31\xeb\x64\xb1\xd0\x9d\xa2\x0f\x02\x22\x84\xb5\xba\x74\xc3\xbc\xba\x76\x29\xff\x92\x07\x29\x80\x41\x00\x53\x15\x26\x9b\xb8\x9a\xc5\x2c\xb8\x9e\x15\x04\xb9\x4a\x6a\x0d\x0c\xe7\x36\xc6\xa4\x70\xa6\x80\xd6\xd8\x32\x36\x6f\x78\x64\xf9\x34\x0a\x24\xc0\xbb\x2a\x1c\x60\xa9\x1b\x64\xc0\xed\x5c\x55\x8e\x6a\x50\x01\x86\xd0\x05\xb5\xec\xb4\xd8\x50\x28\x63\x01\x7c\x19\xa8\xd8\x39\x23\xe8\xb9\x87\x1e\x16\x64\xc3\xd6\x48\xf2\x1c\xea\xd0\xf6\x49\x14\xb2\xfd\x21\x9b\xe5\xb1\x5a\xd3\x3d\x7e\x09\xfe\xf4\x54\xf9\x68\x07\xe1\x93\x00\x2c\x0c\x6d\xf6\xd6\xfc\x65\xc8\xac\x7d\x6a\x8c\x8e\xd9\xb9\x84\x65\x57\x96\x32\x66\xbe\xf0\xf2\x08\x2a\x9c\x39\x97\x27\x21\xc7\xe2\x0f\xab\x37\x3e\xb0\x80\x7a\x8b\x9a\xd7\x59\x03\x1d\xd9\xe1\xef\xbe\x85\x58\x9d\xcc\x6d\xc8\x7e\x25\xa2\x43\xdf\xfb\xd6\x6f\x58\xa9\xad\xcf\xd4\xd1\x27\xd4\x69\xd6\x95\x0e\x20\x0c\x59\x6c\x59\x41\xad\xb1\x42\x1b\xef\xba\x38\x08\xfc\x27\xf7\x17\x81\x14\xbc\x47\x94\x5e\x2f\xe5\x32\x4f\x63\x06\x6a\x74\x48\x8d\xa7\x38\xfc\x5e\xa9\xc2\xeb\xe7\x2a\xdc\x6d\x08\x29\xa0\x1f\xe7\xd9\x3f\xac\x6a\x9b\x93\x34\x15\xe9\x5b\x9e\x65\x80\x4f\xa1\x7a\x2a\x84\xc0\xfa\x6b\xe4\x56\xee\x1f\x46\xbd\x32\xbd\x60\x9c\xd8\x9e\x6d\x20\x1e\xab\x5c\x43\x91\x3a\x00\x13\xf6\x72\x08\x18\x91\x36\x63\xb4\x42\x93\x6e\x60\x56\x3a\xdc\x1f\xf4\xfb\xff\xab\x66\xb9\xe0\xb8\x63\xab\x10\x4c\xfa\xe8\x33\xdc\x23\x28\xd2\xa6\x4c\x04\x55\xff\x84\xd1\xc3\xd5\x9e\x6d\x76\x2b\xb0\x13\x0c\x7e\x17\xb1\x74\x43\x52\x50\x57\x25\x07\x7f\x45\x83\x17\x99\xe1\x50\x25\x1b\x5d\xbf\x11\xa6\xe8\x4e\x02\xa2\x03\x68\x75\x08\x30\x41\x84\xb9\x54\x09\xa7\x23\x45\x32\x54\xf9\xa5\xa3\xb6\x9c\xba\x3c\x15\xe0\x78\x91\xce\x3c\x9d\x90\xcf\xa4\x29\x97\x69\xc8\xf6\x44\x28\x20\x57\x81\xf9\x83\xdc\x34\xf8\x3c\xe5\x2b\x1b\x52\xe0\x73\xf6\x98\x73\xfe\xf4\x16\xa2\x39\x60\x7f\x12\xec\xa0\x31\xb1\x97\x6b\xf3\x42\x74\x43\x8d\xd6\x45\x08\x51\xf4\x79\xea\xae\x94\x88\x46\x18\xec\x61\x67\x3c\xe4\xca\x44\x6d\xc0\x1d\xa6\x3d\x04\xb0\x26\x77\x02\x26\x9b\x06\x98\x86\x2e\x02\x8a\xd6\x2e\x8e\x0c\xa4\x4e\xed\x95\x45\xaf\x4e\xc0\x81\x5c\xed\x71\xcd\xbc\xb9\xd2\xb8\x47\x2d\x15\xe5\x10\x06\xa1\xa6\x6d\xd8\xc4\xb4\x8b\x5c\x4c\x82\xe5\xea\xe2\xf6\x7f\x16\xc2\x46\x14\xe0\x0f\xe9\xb9\x68\xdb\x7b\x86\x1d\xa0\x5e\x06\x9e\x1b\x82\x46\x28\x0e\x34\x04\x1e\x1c\x26\x57\xe5\x3f\xe5\x40\xcd\x7e\x36\x2e\x18\x0c\xbe\x69\x3e\xa6\xfd\xd3\xba\x5d\x93\xfc\xf6\x67\x70\xa9\x60\x16\x70\x5f\x4f\xb9\xd3\x79\xd6\xeb\x76\xab\xf0\x65\x54\x8d\xef\x26\x5b\xb5\x59\xce\x2c\xb8\x82\x0c\xa7\xb3\x0d\x6a\xd1\x52\x3d\x4c\x17\xbd\x3c\x2a\x35\x15\x49\xaf\xdc\x4a\x31\x2c\x8f\x2d\x88\xd0\x39\xa2\x99\xca\xf6\x8d\x60\x6f\x99\xe0\xca\x3e\x9a\x97\x1f\x5c\x9a\x61\xf5\xc2\x92\x05\xd0\x14\xad\xc9\xe9\xf9\xf9\x47\xf6\xe3\xf1\x39\x7b\x8b\x71\x25\x83\x28\x14\x5c\x52\x86\xec\x35\x44\xeb\x36\xd8\x69\x5f\xb0\x8a\xe8\xa4\x56\xb9\x50\xb8\x55\x9b\x4c\x75\xb7\x49\x15\xc4\x8b\xe8\x05\x4b\x2d\x14\x35\x50\x2d\x26\x08\x1a\x86\x6c\x04\xc2\xc7\xc5\x48\xb9\x02\x12\xca\xdc\xc7\x16\xf8\x87\x35\x79\xd7\x7b\x05\xcb\x02\x6d\x26\xe5\x0c\x90\xc3\x22\x35\x45\x9a\x1f\x3b\xcf\x23\x30\xa3\x95\xa9\x9f\xa6\x2d\x4d\xb5\xcc\xa8\x6e\x43\x2a\x16\x05\x16\x7c\xa2\x58\xd9\x7d\x88\x9f\xff\x7f\xe4\x0c\x06\x07\xce\xc1\xc1\x77\xce\xe0\xf0\x68\x78\xd4\x7f\x31\xaa\xec\x12\xb7\x74\x3e\x68\xeb\x7c\xd8\xd6\xd9\xd0\x48\x1a\x43\x54\xfa\x56\x33\x82\x88\xef\xca\x3c\x83\x79\xaf\xd3\xb4\xa2\x95\x46\x58\x45\xeb\xb8\x52\x46\xe7\xa5\x6f\x5a\x8d\xf9\xab\x20\x46\x21\x89\x83\xe2\xbf\x7e\x35\xe5\xeb\x6b\x8b\x0c\xba\x39\x41\xe8\x31\x39\x16\xa0\xe0\xd8\xcf\x46\xe5\x46\xaf\xad\xd9\xc9\x15\x94\x64\x4b\xa3\x0a\x29\x2b\xe7\x5b\x1e\xf0\x6c\x92\xa8\x1c\xa4\x62\xa3\xe4\xe6\xcc\xb8\xb9\xb6\x7a\xbb\x61\x67\xad\x7d\xca\xa4\xb6\xb1\x57\x6d\x19\x3c\x19\xd6\xd7\x00\x09\xed\x4d\x29\xe5\x43\xea\x68\x2c\x5a\x41\x6e\xef\xe6\x86\x30\x8d\x7a\x17\x22\x95\xd3\xa8\xda\x7e\xcd\xba\x1f\x64\x8a\x91\x88\x11\xb2\xd7\x45\x30\xc4\xf6\x2e\xb0\x7d\xf4\xdc\xa8\xde\x43\xd3\x36\x48\xdd\xf8\xc2\x0f\xb2\x95\x8c\xbd\x7f\x77\xfa\x4b\xdd\xbd\x2b\x07\x7f\x35\x63\xef\x56\x2d\x69\x2d\xf6\xe8\x36\xdd\xb2\x49\x4b\xc8\xde\x19\xbe\x23\x34\x53\xf0\x2f\x15\x28\x76\xc0\xb3\x21\x25\x1f\x63\xac\xf0\x89\xbb\x8c\x21\x7b\x97\x47\x53\x9e\x3e\xa7\x1d\x11\xa4\x35\x85\x2c\xd5\x3f\x29\x08\xb2\xda\x80\xf6\x11\xc1\x63\x00\xdc\x8a\xe0\x4d\x90\xa4\xc4\x86\x45\x27\x88\x21\x2a\xff\xeb\xe3\xdb\x53\xe8\x88\xe8\x16\x89\x06\xcf\x36\xb0\x68\x13\x8c\xb6\x61\xd1\x5a\xa4\xa1\xf3\xf3\x4d\x58\xb4\x76\xa6\x0a\xa8\x8b\xab\x7c\xbe\x1d\x9c\x41\x54\xb7\x8f\xfa\x7d\x94\x4b\xd7\x2e\x17\x01\xc2\x40\x48\xf4\xeb\xfc\x7e\x45\x64\x7f\xc9\x7f\xbb\x99\xaf\x6a\xae\x39\x17\xac\xf0\x94\x16\xbb\x36\x10\x1d\xe1\x0d\x9b\x10\x65\x8d\x44\xf0\xb3\x4e\x43\x70\x4a\x14\xe0\x2b\x17\xc8\x6f\x06\xce\x05\xe6\x4e\x8d\x63\xd8\x6c\xb8\x21\xd6\x96\x00\xa2\x7a\x56\x34\xaa\xc8\x82\x86\x08\x82\x8c\xad\x32\x33\xb3\x73\xd2\x34\x03\x2b\xba\x74\x69\x5f\x3f\x50\x9e\x5a\x4b\xca\xd0\xdb\x36\x99\x99\x8c\x1d\x02\x69\xc8\xb5\x47\x48\x84\xe5\xc6\x53\x24\x78\x32\xc8\xa9\x64\x19\x5b\x15\xfc\x86\x13\xb6\x26\x06\x10\xc8\x05\x36\x6b\x6c\x85\x8d\x15\xd0\xf2\xdb\x22\x0e\x21\xe7\xc3\xee\x59\xc1\x91\x0c\x0f\xb7\x70\x54\x5f\x6f\x99\x0b\xe0\x83\x3d\x96\x81\xf4\x16\x4c\xad\x83\xce\xdb\x8d\x8d\x34\x0e\xd6\x93\xe9\x37\x10\x53\xb9\xd5\x9f\x22\xe5\x07\x21\xd5\x9a\xb1\xd7\x41\x4a\xd0\xf4\x96\x12\xeb\x7e\xb6\x1f\xe0\x51\x09\xc9\x59\xc8\x6d\xc2\x39\x0c\x41\xa3\x3d\x88\xa8\x67\x6f\x6f\xab\xce\x24\xca\xfe\x20\xa1\x5e\xc5\xf3\x3c\x84\xc0\x77\xa6\xf7\x85\xb7\x13\xd0\x8d\xe7\x78\xa6\xf3\x87\x48\xf8\x11\x3c\x99\xa7\x30\x6c\xca\x6f\x27\x1c\xba\xbe\xdd\x5f\x97\x2e\xcb\x67\x00\xae\xc6\xd6\xb1\xf5\x57\x12\xf4\xe0\xef\x22\xe8\xe0\xef\x22\xe8\x8b\x07\x17\x14\x4a\x45\x7c\x87\x72\x11\xf7\x0d\x0e\x3c\x75\x33\xc9\x48\x04\x96\x27\x3e\xa4\x9d\xa1\xae\x62\x8c\x75\xea\x5b\x2c\x2d\x29\x35\xa3\xb1\xd5\xf6\xaa\x5b\x42\xac\x51\xe5\x8e\x13\xcf\xaf\x2a\x60\xa9\xc0\x4a\x15\xa8\x64\x90\x52\x1d\x38\x3c\xef\x22\xf6\xd8\x02\x3d\x1e\xd9\xf6\xfe\xc3\xfd\xd7\xed\x60\x46\x25\x11\x18\xe6\x49\xc0\x7d\x59\x97\x5e\x10\x6c\x06\x39\x00\xe5\x37\xdc\xec\x6c\x86\x3b\x74\x53\x86\x28\xc0\x40\x0f\xa2\xd0\x21\x34\x41\xc2\xf2\x5c\x4e\x01\x89\xeb\x6f\x0a\x49\x6a\x06\x09\x9c\xcd\x21\xe0\xe1\x81\xf3\xe2\xbb\xa3\x96\x46\xbf\x92\x79\xfe\xe6\x28\x18\x35\xcb\xc3\x76\x74\xf5\xf8\x70\xf0\xf2\x68\x76\x58\x02\xb5\xc7\xb3\xd9\xec\xfb\x5b\xf0\x77\x2a\x89\xb6\xa2\x26\xaa\x04\x35\x5d\xb4\xe1\xb7\x04\xe9\xee\x80\x90\x61\xa3\xcb\x0c\xd6\xc6\xbe\xe0\xab\xa9\x40\x14\x34\x03\x70\xac\xa1\xf7\x76\x48\x58\x63\x59\x43\x9b\xec\x11\x5d\x1b\xba\x74\x75\xb5\x3e\x22\x69\xe8\x5e\x63\xd6\xc7\xd8\x3e\x7c\xa1\xc0\xea\xd9\x50\x45\x75\x15\x72\x55\x7f\x25\x72\x6e\xdc\xbd\xdd\x80\x38\x8f\xb5\x7f\xb4\x60\xce\x1b\x40\x27\xed\xdb\xca\x13\x07\x8d\xd2\x98\x02\x60\xd5\x2d\xdc\xd6\x5d\x69\xf3\x40\x42\xbb\x94\x71\x4e\xf0\x4a\x42\x74\xba\xc2\x62\xa0\x55\x73\xbd\x04\x89\x35\xc2\xeb\x43\xfc\x75\x21\xa4\xee\x43\x02\xe3\x7e\x40\x41\x71\x54\xbb\xcd\x6e\x8c\x52\xbb\xec\x29\x0f\xbb\xc8\x05\xf0\x86\xce\xa3\x1b\x3a\xdc\x43\x65\x3c\xf6\xf5\x81\xc7\xde\x57\xbc\x9f\x1a\x3e\x51\x12\x3d\xb9\x7e\x6a\xb5\xc9\xa9\xaf\x95\x5a\x6f\x90\x5a\xf6\xaf\x65\xd1\xec\xe4\x9b\x7a\x6d\xc3\x95\x0f\xa8\xdd\x0a\xfa\xd4\x07\xaa\x5b\x54\x8d\x18\xa1\xd0\x75\x91\xd3\xee\xa4\xf4\x26\x44\xbb\x83\xda\xd7\x44\x5d\x5f\x83\x96\xd9\xec\xb2\x20\xed\x0b\xa1\x50\xf3\x03\xaa\x3e\x89\x6e\x63\xd6\x2f\xfa\xfd\xfe\x9f\xa0\xec\x24\x5a\xd7\x2e\x0a\xfc\x60\xea\x5c\xc7\xfb\x0f\xa7\x5a\xb5\x2b\xb0\xcd\x0d\xd3\x76\x8b\x7e\x59\xc6\x8f\x97\x7f\x82\xa2\xeb\xa2\xae\xe9\xbc\x39\x93\x1d\xc3\x4b\x51\xaa\x23\xba\x4d\xe7\x5f\x75\x4c\xc7\x9a\x40\x89\x10\xdd\x8e\xa7\x49\xe6\x08\x74\xe3\x43\x9a\xe2\x56\x73\xd3\xfd\xe4\x1f\x84\xa0\xee\x07\x8e\x80\x19\xbd\x69\x30\xaf\x81\x6a\x37\xb4\xdb\x9f\x02\x95\x20\x40\x73\xe8\x14\x4f\x89\x02\x75\x0c\x63\xeb\x17\x45\x37\xab\xc8\xbc\x8a\xdb\xa6\xcf\xe2\x99\xcf\x41\x72\xc5\x32\x01\x66\xae\xae\x4b\x69\x0e\xad\xaf\x4a\x91\x5b\xfd\x81\xd1\x36\x49\xea\x32\xa8\xa1\xd6\x47\xba\x1b\x6e\x39\x2f\xac\xe8\x7e\xc0\x65\x71\x38\x79\xa3\x74\x4c\x0c\xd9\x8f\x41\x08\x1b\x3c\xa0\xea\x7a\xf5\x8e\xa2\xb8\x93\x4a\x42\xe1\xfa\xda\x82\x6d\x7c\x32\xa3\x5e\xc7\xd6\x2e\xb3\xea\xf7\x5c\x6c\x4b\xf0\xa2\xb5\xa5\xc7\x28\xc8\x4b\xfb\x3d\x94\x94\x44\x96\x79\x62\x43\x83\x52\x0b\x50\x30\x1f\x0f\xfa\x10\x92\xca\x77\x2e\x77\x80\x35\xb6\x0e\x45\x9f\x6e\x11\x46\x6a\xa7\xf4\xc5\x73\x97\x5b\x60\x16\x50\x5e\x79\x41\xd4\xd4\x24\x08\x54\xbc\x2a\xb6\x9a\x8a\xf5\x74\xb7\x86\x56\x1b\x86\xac\xdf\x04\xa9\x00\xa9\x3b\xa8\x00\x0e\xc9\x7f\x99\x8d\x0f\x99\xfe\x71\xc3\xa5\xbb\x02\x56\xc0\x59\xda\xaa\x0f\x03\x0f\xf2\xf8\x42\x84\x60\x80\x63\x6b\xce\x25\xeb\x3b\x2f\x9d\x83\x23\xab\x88\xf0\x66\x8c\x7b\xc0\x45\x9a\x8a\x51\x80\x3e\xd5\xbf\x41\xd1\x77\x1f\x05\xdf\xe1\xcb\xb5\x41\x3e\x20\x75\x43\x4e\x68\x5b\xbf\x87\x4b\x0a\x65\x50\x7f\xde\x5d\xbb\x15\x91\xe2\x82\xc7\xc5\x8d\xc8\xb9\x4c\xa1\x5d\x71\x23\x62\x59\xea\x72\x02\xfe\xd7\x46\x59\xb9\x12\xa1\xdb\x10\xbc\x0c\x01\x1f\xc7\xe7\x4c\xf8\x87\x06\xce\xff\x38\x85\x8f\xe0\x6e\xaa\x54\x3a\xb9\x30\x7a\xf0\x1e\xb6\x57\xcd\x32\x07\x2f\x5a\xf7\xfa\x4f\x31\x5f\x5d\xaf\xe5\xad\x96\x93\x88\xc6\xf1\x81\x52\x31\xe3\xf8\xbe\x27\xdb\xf6\x18\xe7\xa6\xf3\x05\x75\x8d\xb7\x31\xe5\x55\x1f\x80\x57\x62\xeb\x91\x7a\xbb\x63\xa2\xf3\xc0\xc4\xcc\x82\x76\xc3\xad\xca\xf6\xfd\xa9\x7a\xa4\xa2\x9b\x57\xee\x60\xda\x79\xd5\xe2\x73\x79\xe5\x4f\x02\x37\xa6\x4a\xca\x2a\xbd\x8a\x5a\x35\xec\xea\x56\x66\x45\x4a\xdb\xdd\xa4\x74\xd6\x54\xaf\xb3\x2a\x76\x25\xa6\x9f\x4b\xd3\x4a\xf9\x17\xb0\xac\x7f\x9f\xbf\x7f\xe7\x64\x8a\x4f\x30\x5b\x61\x0b\xf0\xb1\x2f\xea\xb4\x40\xd9\x5f\x8a\x77\x71\x48\x85\x92\x21\xc6\xf8\xa7\x2e\x95\x1b\x3a\xf5\x27\x23\xe7\xea\x06\x50\xa4\x7b\x4f\x1e\xaf\xa9\xe4\x89\xe2\x87\xdd\x6a\x37\x75\xd6\x68\x31\x98\x28\x19\x21\x5e\x0e\x26\x2a\x96\x5a\xfb\xd0\x65\xdf\x22\x0f\xc5\x30\xab\xfe\x87\xda\x0f\x1c\x86\xc9\x64\xb5\x25\x08\x6a\x5a\x5a\xc5\x00\xa0\xaf\x78\xef\xd6\x46\xff\xad\x8e\xdf\xf4\xcd\x71\xfb\xf1\x5b\xeb\x0d\xf3\x16\xf7\xf8\x8b\xa3\x3d\x07\x82\x90\x2c\xdf\x55\x6b\x17\x3e\xd4\xf0\xc8\x1c\x95\x19\xb8\xa4\xee\x0d\x31\x9d\x03\xea\xa2\xf7\x4d\x0a\xd7\x15\xcf\xf6\x9c\x01\x8f\xee\x0e\x8b\xde\x92\x1e\xef\x0f\x8a\x6a\xe7\xdd\x65\x12\x47\x1e\xa5\xb7\xc7\xea\x94\x5b\x37\x2b\xce\xb3\x4b\x35\xe8\xf3\xec\x1b\x77\x83\xc0\xdf\x6c\x06\x93\x62\x33\xb8\x79\x50\xd3\xe6\x7e\x23\xd2\x6e\x7e\xf3\x28\x74\xf3\x75\xeb\x11\x1e\x36\x9f\x6a\x7f\xd0\x1b\xad\xee\x32\x88\x7d\xb1\x74\x44\x8c\x69\x0e\xef\xf6\xab\x19\xf2\xfa\x2e\x09\x8d\x1e\x66\xdc\xe4\x8e\xc5\xf3\x8d\xbf\xb1\x37\x7a\x0b\x37\x95\x37\x3b\x22\x34\x9d\xe1\xe2\x16\xcd\x2b\x7f\x3b\xf1\xb2\xd1\xc5\xb0\x68\x3e\xa5\xbd\xab\x97\xfe\x87\x4f\x8f\xdd\xa8\xd5\x49\x77\x38\x9f\x54\xa7\xeb\xc4\x72\x0f\x40\xe0\x49\x8c\x07\x11\xbd\xd7\x74\x20\xd1\x82\x06\x37\x06\x82\xc2\xfc\x69\xf9\xc1\xde\x01\x0a\x47\xd6\xe4\xe1\xcc\x9a\x18\xdf\xdd\xaa\x1b\x66\xfd\xdf\x00\x00\x00\xff\xff\xcc\xf4\x76\x5f\xca\x3a\x00\x00")

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

	info := bindataFileInfo{name: "fcs-lpc-motor.html", size: 15050, mode: os.FileMode(420), modTime: time.Unix(1465377327, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x59\xef\x72\xdb\xb8\x11\xff\x2c\x3d\x05\x82\x2f\xa6\x46\x22\x65\xa7\xbd\xf6\x6a\x49\xee\xf8\x1c\xdf\x5c\xda\x9c\x93\x89\x9d\xe9\xb4\xb1\x27\x03\x91\xa0\x84\x18\x24\x58\x02\x94\x4e\xcd\xe8\x9d\xfa\x0c\x7d\xb2\xee\x02\x20\x45\xfd\x89\x4e\xe9\xf5\x83\x47\x04\x76\xb1\xbb\xf8\x61\xff\x01\x1e\xbf\x78\xf5\xf6\xe6\xe1\xef\xef\x6e\xc9\xdc\x64\xf2\xaa\x3b\x7e\x11\x86\xe4\x46\x15\xab\x52\xcc\xe6\x86\xbc\x3c\xbf\xf8\x03\x79\x98\x73\x32\x53\xa1\xd4\xda\x90\xeb\xca\xcc\x55\xa9\x23\x72\x2d\x25\xb1\x3c\x9a\x94\x5c\xf3\x72\xc1\x93\xa8\x4b\x08\xac\xfe\xa0\x39\x51\x29\x31\x73\xa1\x89\x56\x55\x19\x73\x12\xab\x84\x13\x18\xce\xd4\x82\x97\x39\x4f\xc8\x74\x45\x18\xf9\xe1\xfe\x55\xa8\xcd\x4a\x72\xb7\x4e\x8a\x98\xe7\xb0\xd6\xcc\x99\x21\x31\xcb\xc9\x94\x93\x54\x55\x79\x42\x44\x0e\x93\x9c\xbc\x79\x7d\x73\x7b\x77\x7f\x4b\x52\x21\xb9\xd3\x05\xf6\x3a\xb3\x09\x19\xcf\x39\x4b\xf0\x03\x3e\x33\x6e\x18\xc9\x59\xc6\x27\x74\x21\xf8\xb2\x50\xa5\xa1\x60\x43\x6e\x78\x6e\x26\x74\x29\x12\x33\x9f\x24\x7c\x01\xfa\x42\x3b\x18\x90\x4c\xe4\x22\xab\xb2\x50\xc7\x4c\xf2\xc9\x45\x74\x3e\x00\xa5\xc2\x08\x26\xdb\x53\x15\xec\xd3\x8e\xd9\x14\xa6\x56\x5c\xd3\xb6\xc2\x78\xce\x4a\xcd\x41\x41\x65\xd2\xf0\xfb\x9a\x64\x84\x91\xfc\xea\xc7\x9b\x7b\xf2\xe6\xdd\x0d\x79\xe0\xda\x4c\x79\x1e\xcf\xc7\x43\x47\x70\x4c\x3a\x2e\x45\x61\x88\x2e\xe3\x09\x9d\xaa\x25\x2f\x3f\xc5\x2a\x2b\x54\x0e\xf6\xea\xe1\x92\x4f\x37\xa3\xcf\x3b\xe3\x50\x0a\xc3\x23\x30\x3f\xfa\x0c\xd6\x8c\x87\x4e\x92\x17\x2b\x45\xfe\x0c\xa7\x23\x27\x54\x64\x0e\x84\x79\xc9\xd3\x03\x2a\x0a\x56\xe0\xce\xf0\x2c\xb6\x07\x61\x2c\x99\xd6\x5c\x47\x08\x33\xfd\x15\xb1\x69\x0c\xe6\x14\x71\x98\x29\xa3\xca\xad\x15\x56\x98\xfb\x26\xd6\xd1\xc8\x17\x3f\x20\x04\x5d\x22\x95\x6a\x19\xae\x2e\x09\xab\x8c\x1a\x79\xca\xda\xff\x4e\x55\xb2\x6a\xb1\xa7\x70\x8c\x61\xca\x32\x21\x81\xff\xec\xbd\x9a\x82\xb2\xb3\x01\x39\xfb\x89\xcb\x05\x37\x22\x66\xe4\x8e\x57\x1c\x66\x9a\x89\x01\xb9\x2e\xe1\x20\x07\x44\xb3\x5c\x87\x70\x84\x22\x1d\x6d\x8b\x5b\x72\x74\xe4\x4b\xf2\xbb\xf3\xf3\x6d\xed\x80\x67\x6d\xf9\x78\x58\x3b\xd8\xb8\xc6\x98\x82\x3f\x10\x6d\x4a\x11\x1b\x3a\xea\x2e\x58\x49\x00\x01\x32\x01\x63\x3b\xe8\x29\x97\x84\xde\x0d\xaf\xe9\xa0\xdb\x31\xea\x99\xe7\x30\xc4\x6f\x96\x24\x48\xf9\xf2\x25\xba\x86\xaf\xf5\x1a\xe7\x00\x63\xf3\xa9\x2a\x12\x66\xf8\x66\x51\x09\xea\x60\x8b\x29\x93\x9a\xc3\x10\xbc\x0b\x44\x80\xe8\x4e\x9c\x25\xfa\x92\xe4\x95\x94\x30\xdd\x81\x45\xac\x19\xad\x07\xdd\xf5\xa8\xdb\x0d\xd2\x2a\x8f\x8d\x50\x79\xd0\xc3\x15\x68\x99\x35\x01\x6c\x9b\x71\x73\xa3\xd4\xb3\xe0\x01\x05\x9f\xfc\xf4\xf0\xf6\xaf\xb7\x77\xb4\x37\x72\x4c\x68\xf5\x3e\xcf\x87\xfb\xdb\xf7\x96\x05\x76\x17\x79\x16\xfc\xf1\x33\xb5\x64\xfb\x3b\xea\xae\x83\x5e\xaf\xdb\xb5\xac\x85\x54\x2c\xb9\xb7\x60\xfd\x08\x21\x0b\x3c\x8d\x5d\x18\xc2\x8d\x6d\xb8\x53\x2b\x35\xe7\x4b\x82\x9c\xef\xed\x44\x80\x3a\x1d\x2d\x52\x39\xca\xe2\x90\x0b\x5a\x42\xac\x00\x2b\x01\x20\x71\xb8\x77\x3a\x66\x55\x00\x88\x84\xc6\x46\x22\x8a\x9d\x1a\xfc\xc6\x54\x3b\x89\xc9\x01\xb9\x9c\x8d\x21\x84\x42\xc6\xf2\x44\xbb\x15\x0e\x60\x6f\x56\x04\xd9\xad\x92\x06\x09\x6b\xf8\x43\x31\x78\x14\x11\x32\x45\x1a\x2c\x0a\xfe\x72\xff\xf6\x2e\x42\x37\xc8\x67\x22\x5d\x05\x40\xe8\xa1\xe5\xeb\xc6\x78\xfc\xb9\xd6\x3f\x88\x9c\x95\xab\x7b\xcb\xe7\xf6\x0f\x60\xb5\x91\xba\xf1\x46\xb4\xb7\xc8\x17\x10\x9f\x0d\x50\xa9\x2a\x33\xa0\xbe\x53\x72\x95\x81\xdc\x44\x65\x9e\x21\x92\x0a\xb2\xd2\x03\x2b\xe1\xe8\xa2\x82\x95\x30\x77\x2b\x79\x06\x3f\xe0\x0d\x3b\x08\x39\x80\x1a\x7c\x0e\xc1\xe3\xd0\x39\x08\x8e\xc3\x06\x0d\x89\x0c\xff\xc5\x80\x2a\x16\x2d\x98\xac\xd0\x45\xd7\xdd\xd3\xe1\xf1\x3b\xc7\xd2\x61\x8e\x6f\xfc\x5b\x76\xeb\xe4\x05\x1b\xf9\x68\x81\x17\xff\x75\x58\x31\x84\xbe\x15\x1c\x2b\x21\xc2\x6f\xdc\x38\x1c\x37\x0a\x71\x48\x80\xa8\xbb\x2a\x9b\x82\x03\x27\x2a\xae\xac\x59\x60\xa8\xb7\xf0\x87\xd5\xeb\x24\xb0\x99\x12\x74\x84\xb4\xbf\x91\xd3\xa7\xa1\x96\x02\xfc\x85\xf6\x9c\x1c\x74\x22\x28\x5a\x5a\x41\xb5\x93\x6a\x16\x50\xdc\x0c\xa0\x48\xfc\x81\x80\x9d\xfd\x1d\x70\xd1\x08\x44\xf7\x94\x83\xa8\x79\x9b\x93\x48\x01\xbc\xf9\xcf\x2a\x6f\x03\x65\x99\x00\x19\x07\x14\xf8\x83\xc1\x53\xfa\x48\x0d\xcf\xa0\x48\x30\x53\x95\x9c\x0e\x08\x2d\x94\x16\xc8\x8f\xdf\x65\x91\x69\xfa\xe4\x33\x0a\xf4\x00\x46\xe1\x12\x0b\x8f\x1b\x61\xde\x50\x25\x09\x90\x2e\x80\x74\x3e\x82\x9f\xb1\x13\x1e\x49\x9e\xcf\xcc\x1c\x66\xfa\xfd\x4d\x78\x0b\x3c\x3c\x4b\xff\x28\x50\x72\x63\x0c\x0a\x3e\x86\x71\xa6\x72\xc0\x58\x24\x08\x65\x07\x17\x44\x22\xcf\x79\xf9\xd3\xc3\xcf\x6f\x60\xa9\x33\xe7\xa3\x48\x9e\x6c\xb8\xd6\x2e\xb3\x14\x26\x9e\xff\x8d\x4f\x6f\x58\xb6\x97\x6d\x50\x31\xd6\x5f\x4b\x3a\xaa\xda\x71\x41\x19\x85\x08\xc9\xe8\xde\x59\x3a\xf2\x25\x40\xd9\x77\x9f\x1b\xcb\xfa\xf4\xc9\xf2\x8b\x94\x04\xbb\x34\x32\x99\x40\x19\x71\xc8\xec\xd3\x08\x1d\x8b\x6c\x06\x70\x4d\x1e\xa9\x53\xfc\x48\x6d\x53\xf1\x48\xe7\xc6\x14\x97\xc3\xe1\xc5\x9f\xbe\x8b\x5e\xbe\xbc\x88\x2e\x2e\xfe\x18\xbd\xfc\xfd\x77\x97\xdf\x9f\x0f\xb3\xcf\xc5\x6c\xb8\x00\xc7\x53\x11\x7e\x3e\xd2\xe1\x15\x45\x3c\x08\x87\xc2\xf3\x55\x45\xb4\x81\xcc\xe3\xd3\xaa\x1a\x31\xba\x73\xcf\xd7\x6c\x44\x0c\xc7\xb0\xc8\xce\x93\x3e\xa1\x13\x3a\x6a\x68\x31\x6b\x23\x19\x5b\x09\x91\x2e\xa0\xad\x09\xce\x46\x67\x3d\xc7\x08\xfe\xe2\xdc\x65\x82\xce\x32\x8e\xd9\xae\x9f\xd4\xe5\xdc\x4a\x44\x55\xcc\x7a\x4a\x3d\xbd\x9c\x63\xfd\x09\x62\x0c\x89\xf2\xda\x04\xe7\xbd\xc9\xe4\x8c\x9c\xf5\x1c\x6f\xa4\xab\xa9\x8b\x8b\xe0\xa2\xb7\x59\x84\xf8\xc7\xb0\xeb\x84\xff\xf2\x36\x0d\xdc\x9e\x00\xfe\xf3\x1e\x54\x06\x70\xfc\x7c\x6b\x21\x92\xbd\x51\x83\xd8\x7f\x78\x59\xae\x95\xf0\x6b\x10\x38\x40\x6d\x09\x62\xd5\xd2\x97\xb5\x3d\x2f\x6b\x82\xd7\x67\x26\xac\x8b\xe0\x90\xf7\x2a\x7e\x86\xec\x46\x97\x1a\x4e\x92\x02\x90\xc8\x87\xfd\x04\x62\x3a\x44\xde\xa6\x4c\x37\xa1\x7f\xea\x6a\xe4\xa5\xdb\x99\xc3\x86\xac\xca\x55\xc1\xf3\x1d\x0b\xd7\xfb\x5c\xb1\x54\x9a\xef\xb1\xed\xf3\x65\x5c\x6b\x36\xe3\x87\x93\xb1\xe5\x6e\x35\x43\xde\xf8\x57\xf0\x69\x9b\x01\xd7\xa4\xd4\xa4\xa3\xe1\xa7\x0d\x24\x27\x1d\x3a\x66\x0b\x8b\x8d\xa7\x7a\xf1\xc4\x76\x4c\x4e\x69\x67\x38\x24\x19\x83\x3b\x46\x01\x86\x0d\x08\x4b\x0d\x34\x23\x71\xc9\xa1\xdd\xc0\x0b\x80\x26\x0b\x6c\x19\xa1\x91\x44\x6b\x07\x24\x87\xb4\x03\xb7\x91\x92\x93\x15\x94\x9f\x28\x42\x09\xee\x70\x47\xbe\x47\x70\x4a\x36\xf1\x32\xd9\xd9\x16\xd4\x92\x0f\x0f\x37\xbe\x13\x68\xf6\xe5\xcf\xda\x66\xe9\x02\x6f\x12\x0e\x18\x0b\x9c\x65\xda\xce\xd2\x41\x33\x0f\xe6\x6f\xa5\x16\xdb\x11\x7e\xb5\x34\x74\x3a\xd8\x34\x9e\x00\x1d\xf8\xa6\xc8\x01\xba\x88\xc1\x31\x2d\x78\x9d\xc3\xdd\xf4\xe8\x44\x21\xb6\x89\xdd\x93\x61\x67\x4f\x16\xa1\x8c\x45\x3e\x4c\x04\x96\xc6\x76\x16\x72\xc2\x3c\xfd\x13\xd0\xb9\xf5\xa8\x93\x25\x63\xa5\x3a\x24\x11\xe6\x4f\x95\xc1\xf2\x99\xe4\x87\x84\x58\x82\x3d\xb6\xa6\xd6\xb9\xe4\x65\xa9\x58\x3f\x9b\x62\xd7\xd4\xba\x93\x34\xe2\x52\x2c\x6a\x07\x74\x5a\xa9\x90\xf8\xfa\x94\xfc\xe7\xdf\x37\xd4\xf9\xa3\xf7\xfc\xd6\xe1\x7d\x83\xb2\xc6\x09\xb6\x92\xff\x5b\x37\x6b\xe5\x6f\xca\xc4\x6f\x11\x98\xa6\x1b\x89\xdd\x4e\xdb\x64\xeb\x2b\xdf\x60\x71\xed\x71\x5b\xf2\xdf\xdb\xc9\xff\xc5\xde\x83\xe2\xee\x20\x05\xb4\x44\xd6\x10\xbb\x00\x85\x69\xf2\xa2\xed\xe8\xde\xf8\x0d\x75\x37\x0a\x8e\x1b\x82\x7d\xa2\x65\xdd\x69\x0f\x31\x85\xd6\x22\xad\x14\xd7\x29\xd9\xfe\x0c\x87\x9d\xcd\xea\x56\x04\xb9\x08\x71\xf7\x9c\x16\x47\x91\xed\x4e\x59\xff\x0d\x37\x8d\x1d\x52\x9f\x1a\x3d\x89\xd0\xf8\x06\x82\xa5\xeb\xc5\xbe\x15\x25\x13\xda\xd2\x5a\x24\x0b\xc1\x81\xb6\x0f\x2c\xde\x6f\xfa\x5a\x6d\x1f\xd0\x7d\xd3\xf7\x75\x9c\x44\xd2\x6e\x9e\x5b\xb6\xd5\x9f\xa7\x2c\xaf\xa6\x99\x30\xbf\x7d\x79\xb3\x77\xf7\x61\x97\xae\x9d\x5b\xaf\xf7\xfa\xf2\x5f\xad\xae\x9e\xeb\x70\x75\xdd\x63\x3b\x5e\x5c\x11\x52\x35\xfd\x7c\xb4\xc6\x34\xb0\x03\x63\x24\xac\xf5\x5b\xb5\x05\xfc\x87\x43\x22\x4f\xf0\x3e\x79\xe0\xfe\x01\xab\x7a\x4d\xad\x85\xc1\x47\xca\xcb\x92\x3e\x61\x3c\xd4\x5d\xeb\x06\xc6\x7f\x56\x1c\xee\xc3\x5c\x82\x4b\x42\x7b\xe7\xfb\x75\xf8\x84\xa4\x2a\x79\xe9\x8f\xa2\x90\x6c\x75\x5b\x96\xc0\x80\xb2\xeb\x0a\xeb\x7a\x74\xaf\x02\x22\x75\xc6\xe9\x93\x17\x1f\x33\x40\x8a\xce\x38\x78\xbb\x82\x8c\x72\x79\xf4\xe8\xa8\x16\xd0\xa2\x17\x36\xb8\xdb\x55\xca\x3e\xba\xb8\x53\xdf\xdf\x49\xab\x73\x38\x26\x39\x4f\x2b\x8d\x41\x87\xa6\xd0\x7e\x3d\xfc\x04\x97\x91\xad\x64\xb2\x8f\xa0\x57\xd5\x73\xfa\xd7\xdb\x56\x64\x7a\x76\xc0\x0a\xfb\xc2\x03\xd7\x71\x7b\x05\x28\x4a\x7c\x95\x5a\x49\x0e\x6d\x3f\x4e\x02\x9c\x62\x06\xb7\x58\xc9\x53\xf3\x48\x6d\x73\x5f\xaf\x01\x69\xfe\xac\xe1\xcb\x77\xdd\xf4\x31\xa7\x5e\xb7\x0b\xd8\x56\xb0\x22\xd7\x76\xb0\x12\xc7\xd8\xb1\xda\xfb\x13\xe4\x80\x70\xc5\x96\x72\x3c\x2d\x1b\x5d\x6e\x13\x0d\x13\x1d\x0f\xc1\xc8\x2b\xfa\x7f\x82\x10\xc5\xb6\xc1\x9a\x42\xca\x79\x1e\x6d\x9c\x01\x2f\x7e\xa7\x79\x43\xad\xd0\xae\x38\xa2\xd1\xf9\xdd\x62\x46\x7d\x62\xda\xd5\xf8\x2f\x51\xb4\x15\xba\x8c\xb9\x38\xd6\xb3\x36\xaa\x4b\x8e\x4f\xaa\x3b\xca\x47\x1b\x31\x72\xeb\xba\x04\x7a\x0d\xf7\x82\x20\x84\x54\x99\xd5\x67\x97\xc8\xc8\x06\xf2\x46\xb2\x4b\x51\x21\x32\x85\xd8\xff\x8b\xa4\xe1\xcc\xb8\x99\xab\x64\x82\xf7\x78\x43\x9b\x59\x66\x53\xc8\xa4\xbe\x3b\x36\x4f\x96\x43\xb8\xba\xd8\x8b\xcb\x9f\xf1\x19\x6b\x42\xfb\x3c\xc7\x07\xfe\x0f\xef\x5f\xa3\x8f\x46\xf8\x18\xdc\x32\x62\xeb\x7a\xea\xde\x96\xa7\x95\x31\x70\x63\xf4\xa9\x12\xd3\x9b\x88\x9f\xf1\xce\x6a\x2d\xfc\x11\x0c\x0c\xf0\xd6\xde\xa7\x3d\xb8\xbf\xda\x17\x68\x20\xc6\x4a\xaa\x32\xad\x24\x4c\xb9\x0b\x6e\xbd\xaf\xda\x9a\x6d\xcc\xfa\x14\x7c\xfd\x95\x27\x91\x7f\xbc\x7e\x87\xfe\x18\xdc\x8b\xac\x92\xb6\x18\x92\x90\x78\x1d\xe3\x61\xdb\xaa\xda\x67\xe1\xbc\x22\x56\x14\xf8\x7e\x04\x97\xc7\x24\x48\x64\x6f\x44\xb6\x0f\x7b\xed\xae\xc0\x9d\xcd\xab\x3b\xdc\xf4\x60\x68\x5f\xac\xab\x1c\x1a\x75\x25\x6d\xae\xb4\x3b\xa0\x90\xc9\x54\x65\xf0\x36\x81\xaf\xd2\x92\xc4\x70\x66\xa0\xd6\xfd\xd0\x2b\xac\x90\x63\xf4\x12\xd8\x9d\x4d\x86\xa0\x1d\x66\x41\xe1\x78\xeb\x61\x1d\xe7\xaf\xc6\xc3\xfd\x39\x5c\x3f\x04\x01\x28\x69\x3c\x44\x23\xae\xba\xe3\xa1\xfb\xff\xc8\x7f\x03\x00\x00\xff\xff\x7f\xad\x44\x02\xe7\x19\x00\x00")

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

	info := bindataFileInfo{name: "index.html", size: 6631, mode: os.FileMode(420), modTime: time.Unix(1465377186, 0)}
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
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: k}
	}
	panic("unreachable")
}
