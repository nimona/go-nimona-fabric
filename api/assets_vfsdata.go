// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package api

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 12, 1, 12, 4, 51, 639000788, time.UTC),
		},
		"/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          time.Date(2018, 12, 6, 23, 28, 52, 373077723, time.UTC),
			uncompressedSize: 6663,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\x5b\x73\xdb\x36\x16\x7e\xf7\xaf\x38\xa1\x77\x4b\x29\x16\x49\x39\x1e\xc7\x35\x4d\x29\xdb\x8b\x3d\x6d\x37\x6d\x77\xd6\x99\xf6\xc1\x71\x67\x20\xf2\x48\x42\x0c\x02\x0c\x00\x4a\x56\x12\xff\xf7\x1d\x80\x14\x09\x8a\xb2\xeb\xdd\xd9\x87\x9d\x95\x32\x19\x06\x38\xd7\xef\x5c\xa9\x24\x2f\x82\x00\xbe\x25\x0a\x33\x10\x1c\x96\x5a\x17\x2a\x8e\xa2\x19\x13\x8b\x50\x23\xc9\xb5\x44\x5c\x8a\x52\x61\x98\x8a\x3c\x22\x3c\xa0\x5c\x4b\x91\x95\xa9\xa6\x82\x07\x5a\x04\x6b\x9c\x29\x91\xde\xa1\x56\x10\x04\xd3\x83\x64\xa9\x73\x36\x3d\x38\x48\x96\x48\xb2\xe9\x01\x40\xa2\xf4\x86\xa1\x79\x02\x78\x39\x82\x97\xf1\x0c\xe7\x42\xa2\x79\x22\x73\x8d\x12\x3e\xdb\x2b\xf3\x09\x72\xf1\x29\x98\x89\xfb\x40\xd1\x4f\x94\x2f\x62\x98\x09\x99\xa1\x34\x47\x17\x2d\xd1\x1a\x67\x77\x54\xff\x29\xdd\xd3\xf7\x0f\x07\xf5\x83\x31\xd7\x31\x61\x2e\xb8\x0e\xe6\x24\xa7\x6c\x13\xc3\x0f\xc8\x56\xa8\x69\x4a\x46\xf0\x8d\xa4\x84\x8d\x40\x11\xae\x02\x85\x92\xce\x2f\xba\x2c\x8a\x7e\xc2\x18\x8e\xc7\xe3\xbf\x3a\x16\x90\xf4\x6e\x21\x45\xc9\xb3\x18\x0e\x4f\x4e\x4e\x7a\xba\x0f\x0b\xb2\xc0\x60\x2d\x49\x51\x74\x70\x58\xd3\x4c\x2f\x63\x78\x7d\x3a\x2e\xee\x1f\x11\x77\x75\x75\xd5\xde\x14\x24\xcb\xac\x9f\xc7\x98\xb7\xa7\x39\x91\x0b\xca\xed\x21\x90\x52\x0b\x17\x1a\x0b\x87\x16\x45\x0c\xa7\xc5\x3d\x28\xc1\x68\x06\x87\xaf\xcf\xd3\xb3\xb3\x93\x1d\x04\x97\x24\x13\xeb\x18\xc6\xf0\xaa\xb8\x87\xe3\x71\x71\x0f\x72\x31\x23\x83\xf1\xc8\x7e\xc3\xaf\x87\x7d\x44\x8f\x1d\x57\x2a\x23\x2a\x55\xe3\x3e\x00\x4a\x13\x5d\xaa\x5d\xfc\x2b\x30\xc7\xe1\xb9\xec\xfb\x13\xcc\x84\xd6\x22\x8f\xe1\xd8\xb9\x6c\x04\x86\xa2\x40\xee\x88\x4b\x05\x13\x32\x86\x85\x44\xe4\x7d\xe2\x94\x09\x93\xf6\x3d\x72\x89\x99\x43\x5c\x3f\x95\x6e\x9a\x30\xaa\x74\x60\x13\x3b\x06\x2e\x38\xf6\x51\x1f\xef\x09\xcf\x78\x6f\xd2\x8c\xc3\xf3\xd3\x7d\xbe\x94\x0c\x18\x75\x74\xb6\x62\x42\x43\x0f\xe3\xf0\xec\xb4\x03\x50\x93\xe5\x35\x40\x6d\x64\x2f\x2f\x2f\xf7\x8b\x8f\xe7\x54\x2a\x1d\xa4\x4b\xca\x5c\x1c\xdc\x04\x79\x8e\x18\xc8\xe8\x2a\xe6\x7a\x59\x09\x1a\xbc\x1a\x3a\xb2\x34\xde\xeb\x40\xac\x50\xce\x99\x49\x24\x64\x8c\x16\x8a\xaa\xd6\xec\xf5\x92\x6a\x0c\x54\x41\x52\x0b\xa6\x29\x87\xf6\xb2\x65\x5c\xd2\x2c\xdb\x13\xc4\x4a\xbf\x2a\x88\x1b\xf6\x8c\xaa\x82\x91\x4d\x0c\x94\x33\xca\x31\x98\x31\x91\xde\x5d\xec\xd6\xd7\x79\xa7\xbc\x6c\x44\xd6\x48\x17\x4b\x6d\x1a\x06\xcb\x2e\x76\xd3\xe2\xf0\xfc\xfc\xfc\x91\x08\x9e\x75\x02\x61\x5d\xd6\x92\x70\x35\x17\x32\x8f\xa1\x34\xf5\x9d\x12\xe5\xa4\x09\x43\xad\x51\x5a\xaf\xab\xca\x2d\xfa\xad\x29\x54\xc8\xb5\x1b\x95\xa6\x01\x04\x5b\x8b\xae\xce\xcc\xb7\x5f\x57\x39\x2a\x65\x7a\x8b\xd1\xff\x48\x35\x1e\x87\x7b\xb3\xce\xd8\x4e\x24\x92\x7e\x3b\xea\x36\xb7\x9d\x6c\xdc\xdf\x0e\xf7\xa4\x67\x27\xa1\xbe\x3f\x37\xdf\x5e\x06\x4b\x92\xd1\x52\xc5\x70\x52\xec\xf6\xf3\xba\x1b\x51\xae\x50\xc3\xd8\x8a\x3a\x6e\x5b\x12\xd4\x7f\xc2\xe3\xa1\x53\x90\x94\x07\xcb\x3a\xac\xc7\xe3\x4e\xc8\x9f\xd3\x51\x28\x2f\x4a\xfd\xff\x03\xc6\x33\x3c\x9e\x95\x5a\x8b\x7f\xa3\x9a\x9e\xb6\xb3\x72\xb3\xdb\x23\x9f\x68\xf2\x0e\x92\xaf\x4d\x93\xeb\x4c\xb4\x3a\xeb\x6d\xc3\xd8\xd3\x71\x61\x1c\xbe\x3a\xed\x17\x22\x61\x74\xc1\x63\x48\x91\x6b\x94\x8f\x8c\xd3\x6f\xbf\x31\xdf\xe7\xb4\x52\xa7\x07\xec\x20\x76\xa3\x37\x05\x4e\x3c\x55\xce\x72\xaa\xbd\xdb\xbd\x85\x1b\xc3\xe1\xd7\xaf\x67\x27\xaf\xb2\xe7\xa8\x3a\xcd\xce\xb2\xe3\xf9\x23\xda\xe2\xa5\x69\x8d\x8e\x12\x61\x5a\x89\xde\xd8\x6e\x74\xea\x60\x56\x4a\x65\x40\x2b\x04\x75\xfd\x6f\xa4\x91\xc7\x96\x9f\x5c\x70\x61\x9b\xf2\x0e\x9c\x19\xa6\x42\x12\xb3\xfd\xed\x86\xb5\x0e\x4f\xd7\xc1\x46\x51\x21\xf1\xc9\x5e\xb6\x58\x0a\xa5\x77\x42\xdb\xaf\x12\x45\xd9\xca\x8d\x62\xbb\xfa\x98\xd5\xe4\xd5\x9e\xf2\x8e\x3b\xa7\x6d\x73\xfd\xa0\x04\x0f\xee\x70\xd3\x1f\xff\x33\x29\xd6\x7b\xb6\x05\xcb\xb0\x22\xac\xc4\x3e\x0b\x27\xab\xcd\x23\x1c\x4a\x4b\xca\x17\x7d\x16\xc1\xe8\x0a\x5b\x1e\xf3\x77\x12\x35\xbb\x72\xa2\x52\x49\x0b\x5d\xad\xcd\x51\x64\xd0\xd3\x7a\x03\x85\xa4\x5c\x1b\x71\x73\x29\x72\xbb\xac\xc7\x51\xf4\x41\xcd\x69\x96\x31\x0c\x39\xea\xa8\xe4\x6f\xaf\x7f\x8a\x2c\x1b\x9d\xc3\xe0\x05\xa3\x33\x49\xe4\x66\x58\x2b\x5a\x11\x09\xf5\x11\x4c\xe0\xf3\xc3\x45\x65\x6c\x7d\x64\x2d\x36\xe7\x35\xb5\xc4\x82\x91\xd4\x44\x60\x5e\x72\xbb\xf1\xc3\x20\x27\x3a\x5d\x8e\xa0\xf8\x91\x67\xc8\xf5\x08\x8a\xbf\xe3\x66\x04\xc5\x6f\x66\x3d\x2e\x2e\x79\xe6\xce\x7e\xa3\xcd\x40\x3c\x01\x3f\xb1\x23\x3a\x65\x44\xa9\xc9\x16\xfa\xa9\x7f\xd1\x21\x5d\x11\xb6\x97\xd4\x82\xbe\x4b\xac\xb4\xdc\x4b\x5c\xe1\xbd\x4b\x6d\x68\x6b\x93\xe1\xcb\x17\xf0\x9d\x6b\x03\x93\x71\x62\xd8\x9c\x80\x25\x97\x70\x64\x8d\x3f\xb2\x2e\x86\x35\x16\x83\xe8\xc6\x8b\xe1\x36\x5a\x8c\xc0\xf7\x87\x70\x04\x7e\x12\x19\x13\xa6\x31\xec\xca\xfc\x8d\xb0\x7d\x32\xed\xc5\xcd\xf8\x16\x26\x13\xf0\x3d\x1f\xde\x58\x57\x62\xe3\xbd\x91\x67\x6e\x1d\xb1\x8e\x50\x89\xba\x94\xbc\x96\x71\xc9\xb3\xca\x91\x76\x01\x1f\xb5\x95\xa6\xf5\xe6\x1f\x26\x55\xdc\xc0\x89\xd9\x87\xdd\xd8\x18\xc0\xde\x52\x8e\x30\x81\xe8\x8f\x01\xbc\x1c\x0e\xbc\x9b\xf7\xeb\xdb\x23\x2f\x86\xe1\x9b\x81\x77\xf3\x87\x77\xfb\xd2\xfb\x72\xf3\x7e\x1d\x1e\x05\xb7\x2f\x87\x6f\x06\x37\xa3\x9b\xcf\xb7\xc3\x37\x7f\x89\xf2\x45\xcf\xae\x9f\xae\x7f\xfd\x25\xac\xd0\xa7\xf3\x8d\x51\x37\x02\x5e\x32\x36\x82\x13\x17\x86\x16\xc8\xaf\x2c\x88\x5f\x91\xbc\xb8\xf0\x87\xed\xf1\xfb\xf7\x5e\x75\xf1\xb1\x14\xfa\xc2\xdf\xcf\x9b\x54\x24\x4c\x77\x58\xa7\xd5\xe9\xe2\x31\xb6\xad\xbf\xa3\x4e\xc2\x6f\xaf\xe5\xb0\x5b\x8e\x0f\xe6\x9f\x49\xb4\xad\xc4\x24\xaa\x5e\x69\x0f\x92\x99\xc8\x36\xb6\x48\x33\xba\x02\x9a\x4d\x3c\xf7\x3d\xce\xab\x6a\x36\x59\x9e\x4c\x7f\xa1\xb9\xe0\x04\x9a\x37\x64\xc0\x7b\x92\x17\x0c\x93\x68\x79\x52\x53\x6d\x25\x54\x2f\x42\xde\xf4\x3b\xc1\x39\xa6\xa6\xc4\xc3\x30\x4c\xa2\x8c\xae\xa6\x55\x85\x26\x76\x95\x33\xa4\xee\x6e\xe7\x01\xb1\xd1\x9d\x78\x87\x1e\xe4\xa8\x97\xc2\x58\x23\x94\xae\xad\x00\x48\x9a\x85\xce\xe1\xf5\xc0\x3a\xbc\x14\x2c\x43\x39\xf1\x7e\x97\x54\x23\x6c\x44\x29\xa1\x26\x80\x25\x4a\x0c\xc3\xd0\x03\x89\x1f\x4b\x2a\x31\x9b\x26\xd1\x56\x52\x23\xba\xda\x8d\x8c\x5c\x89\x29\x2d\x28\x72\xbd\x23\xf9\x9f\xdb\x73\x5f\x99\x5a\x6a\xc5\x41\xd4\x48\xa9\xf7\x8d\xce\xf4\x9c\x5e\x23\xcf\xe0\xe7\xca\x98\x24\xaa\x48\x6a\xc8\x22\xe3\xf8\x16\x95\x92\xb9\x7e\x29\x6f\x9a\x44\x25\xb3\xb1\x69\xb0\xeb\xb4\xd2\x52\xda\x16\xb3\x56\x71\x14\xf9\x70\x04\x4c\xa4\x76\x96\x85\x66\xfa\x70\x92\xa3\x29\xae\xe6\xb0\x10\x52\xc3\x1b\xf0\xe3\x0e\xa9\x3d\x8d\xb7\xe5\x1f\x91\x82\x46\xab\xe3\x48\x69\x89\x24\x57\x11\xb7\x41\x0f\xa9\x88\x6a\xa3\xfc\xba\xc7\xae\x29\xcf\xc4\x3a\x14\x9c\x09\x92\xc1\xc4\xa9\xcc\xb6\x2c\x4d\x51\xda\x40\x4f\x20\x13\x69\x99\x23\xd7\xe1\x02\xf5\x25\x43\xf3\xf8\xed\xe6\xc7\x6c\xe0\xbb\xf1\x6f\xab\xdf\x70\xd6\x37\x57\x14\x59\xf6\x0c\x09\x5d\xe6\x26\x84\x7f\xca\xde\x50\xee\xd5\xae\xde\x52\xa5\x9f\xa1\x5d\x75\xb9\xab\x0a\xb9\xae\x7e\x0f\x78\x82\xbb\x2a\x14\xc3\x5b\x33\x37\x94\x1f\x4b\x94\x9b\x6b\x64\x98\x6a\x21\x07\xbe\x29\x52\x7f\x18\x92\x2c\xbb\x5c\x21\xd7\xc6\x2a\xe4\x28\x07\x7e\xca\x68\x7a\xe7\x8f\x1c\xfc\xd1\xed\x8b\xa6\x75\x63\xa8\x89\x5c\xa0\x0e\xed\x5c\xf9\x85\xe4\x18\x6a\xf1\x56\xac\x51\x7e\x47\x14\x0e\x86\x30\x31\xbd\xdb\xcc\x60\xdf\x65\x85\x1d\x10\xc3\x6a\x5b\x98\x40\x23\x2f\x23\x9a\x28\xd4\xe1\x1d\x6e\xda\xfe\xf9\xb0\xed\x3a\xad\x53\x29\x4c\x80\xe3\x1a\x7e\xc7\xd9\xb5\x05\x66\x50\x4a\xd6\xe0\x95\x86\x82\xdb\x1f\x3a\xdc\x24\x42\xe3\xa5\x6b\x8d\x8b\x68\x48\x39\x47\xf9\xc3\xbb\x9f\xdf\x9a\xec\xaf\x9b\x0c\x66\xa0\x45\x0c\x26\xb7\x2d\x73\x98\x96\x52\x22\xd7\xef\x2a\x5b\x4b\xc9\x2e\xf6\x0b\x6b\x50\x31\xc2\x8c\x21\xcd\x90\x7a\x68\x3c\x30\xd9\x19\x0a\x5e\xd5\x73\xd7\x50\xd7\x48\x0c\x0b\x69\x95\x7f\x8f\x73\x52\x32\x3d\x18\x76\xe7\xb6\x89\x22\x4c\x3a\x99\x5d\xa1\xba\x33\xde\xb7\xb8\x9b\xdc\xb9\xd9\x17\x85\xdb\x2e\x83\x98\x7d\xc0\x54\x3b\xab\x8e\xf9\x78\x7f\x4b\xf5\xbd\x17\x83\xd7\xab\x61\x6f\xd4\x21\x2b\x04\xa3\xe9\xc6\x8b\x3b\xdc\x0e\x7f\x54\x13\x8c\xba\xd7\xaa\xb4\x5a\x95\x17\x3b\x06\xef\xd0\x54\xfd\xdc\x90\xdc\x78\x12\x49\xe6\xdd\xee\x10\xe0\x7c\x8e\xa9\x36\x5a\x08\x63\x62\xed\x39\xb7\x0f\x1d\x2b\x0d\x74\x5e\x6c\x11\xec\x9c\x6b\x9a\xa3\xd2\x24\x2f\xbc\x18\x06\x26\xc9\xbe\x27\x1a\x07\xc3\x61\xa8\xc5\x8f\xd7\xbf\x5e\xdb\x01\x3e\x18\xf6\xb2\xd3\xa4\x9d\x42\x9e\x0d\xfa\x73\x1e\x53\x3d\x1c\xf6\xf6\x81\x39\x61\xed\x2f\x1e\x6d\x62\x98\xe4\x45\x29\x85\xec\x26\x85\x39\x19\x76\xb6\x64\xae\x04\xc3\xd0\x5e\x0c\xfc\xa6\x10\xe0\xd2\x1c\xd4\x59\x6b\x99\xf6\xab\xd8\x0e\x32\x57\x49\xae\x16\xbb\x1b\x50\x93\x07\xd6\xab\x82\x48\x85\x86\xcc\x16\xea\xf0\xa2\x67\x0d\x13\x8b\xad\xc3\x9d\x86\xf1\xe2\x45\x75\x1a\x8a\x35\x47\xd9\xed\x09\x6e\x5f\xec\xd6\x61\xc2\x68\xbd\xb9\x9a\x01\x8a\x74\x85\x99\x37\xf5\xe1\xa8\x13\x70\xdf\xec\x08\x53\xbb\xe6\x4e\xeb\x81\x18\x6f\x77\x43\x38\xaa\x1d\x08\x6d\x9d\xd8\xad\xd1\x50\x3f\x25\xe3\x4a\x8a\x7c\x2b\x20\x21\xb0\x94\x38\xb7\xdb\x43\x6d\x88\x69\x6a\x1e\x18\xf7\xcd\x86\x3e\xf1\x1c\x1d\xd6\x37\xa3\xc4\x9b\xee\x3b\x4d\x22\x32\x7d\x4c\x7d\xc4\xa8\xe5\xd9\x8f\x45\x9b\x6b\x80\x4c\xe1\xff\x32\x7a\xef\xc4\x7f\x86\x5d\x5b\xef\x66\xf3\xdf\xc5\xb0\x77\xfb\xdf\xc0\xb2\xbb\xcb\x1e\x74\x76\xd9\x83\x24\xaa\x96\x58\xb3\xd5\xda\xff\xb0\xf9\x57\x00\x00\x00\xff\xff\x60\xf8\x9e\x2b\x07\x1a\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/index.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
