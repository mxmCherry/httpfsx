// Static is a generated package,
// that embeds files from "/home/mxmcherry/go/src/github.com/mxmCherry/httpfsx/internal/statichandler/public" directory.
// WARNING!!! Don't edit this file manually!!!
package static

import (
	"bytes"
	"net/http"
	"path"
	"sync"
	"time"
)

// File is a file record.
type File struct {
	Contents []byte
}

// ModTime holds package modification (generation) time.
var ModTime = time.Unix(1486849440, 0)

// Files hold path-to-contents file mapping.
var Files map[string]File

func init() {
	Files = map[string]File{}
	Files["/script.js"] = File{
		Contents: []byte{
			0x28, 0x66, 0x75, 0x6E, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x28, 0x29, 0x20, 0x7B,
			0x0A, 0x09, 0x27, 0x75, 0x73, 0x65, 0x20, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
			0x27, 0x0A, 0x0A, 0x09, 0x2F, 0x2F, 0x20, 0x72, 0x6F, 0x6F, 0x74, 0x20, 0x65,
			0x6C, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x3A, 0x0A, 0x09, 0x76, 0x61, 0x72, 0x20,
			0x68, 0x74, 0x74, 0x70, 0x66, 0x73, 0x78, 0x20, 0x3D, 0x20, 0x64, 0x6F, 0x63,
			0x75, 0x6D, 0x65, 0x6E, 0x74, 0x2E, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65,
			0x6C, 0x65, 0x63, 0x74, 0x6F, 0x72, 0x28, 0x27, 0x2E, 0x68, 0x74, 0x74, 0x70,
			0x66, 0x73, 0x78, 0x27, 0x29, 0x0A, 0x0A, 0x09, 0x2F, 0x2F, 0x20, 0x66, 0x69,
			0x6C, 0x65, 0x2D, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6D, 0x20, 0x69, 0x74, 0x65,
			0x6D, 0x20, 0x6E, 0x6F, 0x64, 0x65, 0x73, 0x3A, 0x0A, 0x09, 0x76, 0x61, 0x72,
			0x20, 0x69, 0x74, 0x65, 0x6D, 0x73, 0x20, 0x3D, 0x20, 0x68, 0x74, 0x74, 0x70,
			0x66, 0x73, 0x78, 0x2E, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x6C, 0x65,
			0x63, 0x74, 0x6F, 0x72, 0x41, 0x6C, 0x6C, 0x28, 0x27, 0x2E, 0x68, 0x74, 0x74,
			0x70, 0x66, 0x73, 0x78, 0x20, 0x2E, 0x6C, 0x69, 0x73, 0x74, 0x20, 0x2E, 0x69,
			0x74, 0x65, 0x6D, 0x27, 0x29, 0x0A, 0x0A, 0x09, 0x2F, 0x2F, 0x20, 0x70, 0x61,
			0x74, 0x68, 0x73, 0x2C, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x61, 0x72, 0x65,
			0x20, 0x6C, 0x69, 0x73, 0x74, 0x65, 0x64, 0x20, 0x6F, 0x6E, 0x20, 0x63, 0x75,
			0x72, 0x72, 0x65, 0x6E, 0x74, 0x20, 0x70, 0x61, 0x67, 0x65, 0x3A, 0x0A, 0x09,
			0x76, 0x61, 0x72, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6E, 0x67, 0x50,
			0x61, 0x74, 0x68, 0x73, 0x20, 0x3D, 0x20, 0x5B, 0x5D, 0x0A, 0x0A, 0x09, 0x66,
			0x6F, 0x72, 0x28, 0x20, 0x76, 0x61, 0x72, 0x20, 0x69, 0x20, 0x3D, 0x20, 0x30,
			0x3B, 0x20, 0x69, 0x20, 0x3C, 0x20, 0x69, 0x74, 0x65, 0x6D, 0x73, 0x2E, 0x6C,
			0x65, 0x6E, 0x67, 0x74, 0x68, 0x3B, 0x20, 0x69, 0x2B, 0x2B, 0x20, 0x29, 0x20,
			0x7B, 0x0A, 0x0A, 0x09, 0x09, 0x2F, 0x2F, 0x20, 0x69, 0x74, 0x65, 0x6D, 0x20,
			0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x20, 0x6E, 0x6F, 0x64, 0x65, 0x3A,
			0x0A, 0x09, 0x09, 0x76, 0x61, 0x72, 0x20, 0x69, 0x74, 0x65, 0x6D, 0x20, 0x3D,
			0x20, 0x69, 0x74, 0x65, 0x6D, 0x73, 0x5B, 0x69, 0x5D, 0x0A, 0x0A, 0x09, 0x09,
			0x76, 0x61, 0x72, 0x20, 0x73, 0x74, 0x61, 0x72, 0x20, 0x3D, 0x20, 0x69, 0x74,
			0x65, 0x6D, 0x2E, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x6C, 0x65, 0x63,
			0x74, 0x6F, 0x72, 0x28, 0x27, 0x2E, 0x73, 0x74, 0x61, 0x72, 0x27, 0x29, 0x20,
			0x2F, 0x2F, 0x20, 0x73, 0x74, 0x61, 0x72, 0x72, 0x69, 0x6E, 0x67, 0x20, 0x65,
			0x6C, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x0A, 0x09, 0x09, 0x76, 0x61, 0x72, 0x20,
			0x6C, 0x69, 0x6E, 0x6B, 0x20, 0x3D, 0x20, 0x69, 0x74, 0x65, 0x6D, 0x2E, 0x71,
			0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x6C, 0x65, 0x63, 0x74, 0x6F, 0x72, 0x28,
			0x27, 0x2E, 0x6C, 0x69, 0x6E, 0x6B, 0x27, 0x29, 0x20, 0x2F, 0x2F, 0x20, 0x6C,
			0x69, 0x6E, 0x6B, 0x20, 0x65, 0x6C, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x20, 0x28,
			0x66, 0x6F, 0x72, 0x20, 0x69, 0x74, 0x65, 0x6D, 0x27, 0x73, 0x20, 0x70, 0x61,
			0x74, 0x68, 0x20, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x29,
			0x0A, 0x0A, 0x09, 0x09, 0x2F, 0x2F, 0x20, 0x69, 0x74, 0x65, 0x6D, 0x27, 0x73,
			0x20, 0x70, 0x61, 0x74, 0x68, 0x3A, 0x0A, 0x09, 0x09, 0x76, 0x61, 0x72, 0x20,
			0x70, 0x61, 0x74, 0x68, 0x20, 0x3D, 0x20, 0x6C, 0x69, 0x6E, 0x6B, 0x2E, 0x67,
			0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x28, 0x27,
			0x68, 0x72, 0x65, 0x66, 0x27, 0x29, 0x0A, 0x0A, 0x09, 0x09, 0x2F, 0x2F, 0x20,
			0x77, 0x68, 0x61, 0x74, 0x20, 0x6B, 0x65, 0x79, 0x20, 0x69, 0x73, 0x20, 0x75,
			0x73, 0x65, 0x64, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x73, 0x74, 0x6F, 0x72, 0x69,
			0x6E, 0x67, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6E, 0x74, 0x20, 0x69, 0x74,
			0x65, 0x6D, 0x27, 0x73, 0x20, 0x73, 0x74, 0x61, 0x72, 0x72, 0x65, 0x64, 0x20,
			0x73, 0x74, 0x61, 0x74, 0x65, 0x3A, 0x0A, 0x09, 0x09, 0x76, 0x61, 0x72, 0x20,
			0x73, 0x74, 0x61, 0x72, 0x4B, 0x65, 0x79, 0x20, 0x3D, 0x20, 0x27, 0x68, 0x74,
			0x74, 0x70, 0x66, 0x73, 0x78, 0x3A, 0x73, 0x74, 0x61, 0x72, 0x3A, 0x27, 0x20,
			0x2B, 0x20, 0x70, 0x61, 0x74, 0x68, 0x0A, 0x0A, 0x09, 0x09, 0x2F, 0x2F, 0x20,
			0x72, 0x65, 0x6D, 0x65, 0x6D, 0x62, 0x65, 0x72, 0x20, 0x22, 0x73, 0x74, 0x61,
			0x72, 0x72, 0x69, 0x6E, 0x67, 0x22, 0x20, 0x6B, 0x65, 0x79, 0x20, 0x74, 0x6F,
			0x20, 0x73, 0x69, 0x6D, 0x70, 0x6C, 0x69, 0x66, 0x79, 0x20, 0x22, 0x74, 0x6F,
			0x67, 0x67, 0x6C, 0x65, 0x20, 0x73, 0x74, 0x61, 0x72, 0x72, 0x69, 0x6E, 0x67,
			0x22, 0x20, 0x63, 0x6C, 0x69, 0x63, 0x6B, 0x20, 0x68, 0x61, 0x6E, 0x64, 0x6C,
			0x65, 0x72, 0x3A, 0x0A, 0x09, 0x09, 0x73, 0x74, 0x61, 0x72, 0x2E, 0x73, 0x65,
			0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x28, 0x27, 0x64,
			0x61, 0x74, 0x61, 0x2D, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73, 0x78, 0x2D, 0x73,
			0x74, 0x61, 0x72, 0x2D, 0x6B, 0x65, 0x79, 0x27, 0x2C, 0x20, 0x73, 0x74, 0x61,
			0x72, 0x4B, 0x65, 0x79, 0x29, 0x0A, 0x0A, 0x09, 0x09, 0x2F, 0x2F, 0x20, 0x63,
			0x68, 0x61, 0x6E, 0x67, 0x65, 0x20, 0x73, 0x74, 0x61, 0x72, 0x27, 0x73, 0x20,
			0x76, 0x69, 0x65, 0x77, 0x2C, 0x20, 0x69, 0x66, 0x20, 0x69, 0x74, 0x65, 0x6D,
			0x20, 0x69, 0x73, 0x20, 0x73, 0x74, 0x61, 0x72, 0x72, 0x65, 0x64, 0x3A, 0x0A,
			0x09, 0x09, 0x69, 0x66, 0x28, 0x20, 0x6C, 0x6F, 0x63, 0x61, 0x6C, 0x53, 0x74,
			0x6F, 0x72, 0x61, 0x67, 0x65, 0x2E, 0x67, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6D,
			0x28, 0x73, 0x74, 0x61, 0x72, 0x4B, 0x65, 0x79, 0x29, 0x20, 0x29, 0x20, 0x7B,
			0x0A, 0x09, 0x09, 0x09, 0x73, 0x74, 0x61, 0x72, 0x2E, 0x63, 0x6C, 0x61, 0x73,
			0x73, 0x4C, 0x69, 0x73, 0x74, 0x2E, 0x61, 0x64, 0x64, 0x28, 0x27, 0x61, 0x63,
			0x74, 0x69, 0x76, 0x65, 0x27, 0x29, 0x0A, 0x09, 0x09, 0x7D, 0x0A, 0x0A, 0x09,
			0x09, 0x2F, 0x2F, 0x20, 0x72, 0x65, 0x6D, 0x65, 0x6D, 0x62, 0x65, 0x72, 0x20,
			0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x74, 0x65, 0x6D, 0x27, 0x73, 0x20, 0x70,
			0x61, 0x74, 0x68, 0x20, 0x74, 0x6F, 0x20, 0x72, 0x65, 0x6D, 0x6F, 0x76, 0x65,
			0x20, 0x64, 0x65, 0x6C, 0x65, 0x74, 0x65, 0x64, 0x20, 0x69, 0x74, 0x65, 0x6D,
			0x73, 0x20, 0x66, 0x72, 0x6F, 0x6D, 0x20, 0x6C, 0x6F, 0x63, 0x61, 0x6C, 0x53,
			0x74, 0x6F, 0x72, 0x61, 0x67, 0x65, 0x20, 0x64, 0x6F, 0x77, 0x6E, 0x20, 0x74,
			0x68, 0x65, 0x20, 0x63, 0x6F, 0x64, 0x65, 0x3A, 0x0A, 0x09, 0x09, 0x65, 0x78,
			0x69, 0x73, 0x74, 0x69, 0x6E, 0x67, 0x50, 0x61, 0x74, 0x68, 0x73, 0x2E, 0x70,
			0x75, 0x73, 0x68, 0x28, 0x70, 0x61, 0x74, 0x68, 0x2E, 0x72, 0x65, 0x70, 0x6C,
			0x61, 0x63, 0x65, 0x28, 0x2F, 0x5C, 0x2F, 0x7B, 0x32, 0x2C, 0x7D, 0x7C, 0x5C,
			0x2F, 0x24, 0x2F, 0x67, 0x2C, 0x20, 0x27, 0x27, 0x29, 0x29, 0x0A, 0x0A, 0x09,
			0x7D, 0x0A, 0x0A, 0x09, 0x2F, 0x2F, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6E,
			0x74, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x28, 0x6C, 0x6F,
			0x63, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x29, 0x20, 0x70, 0x61, 0x74, 0x68, 0x3A,
			0x0A, 0x09, 0x76, 0x61, 0x72, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6E, 0x74,
			0x50, 0x61, 0x74, 0x68, 0x20, 0x3D, 0x20, 0x6C, 0x6F, 0x63, 0x61, 0x74, 0x69,
			0x6F, 0x6E, 0x2E, 0x70, 0x61, 0x74, 0x68, 0x6E, 0x61, 0x6D, 0x65, 0x2E, 0x72,
			0x65, 0x70, 0x6C, 0x61, 0x63, 0x65, 0x28, 0x2F, 0x5C, 0x2F, 0x7B, 0x32, 0x2C,
			0x7D, 0x7C, 0x5C, 0x2F, 0x24, 0x2F, 0x67, 0x2C, 0x20, 0x27, 0x27, 0x29, 0x0A,
			0x0A, 0x09, 0x2F, 0x2F, 0x20, 0x74, 0x72, 0x61, 0x76, 0x65, 0x72, 0x73, 0x69,
			0x6E, 0x67, 0x20, 0x6C, 0x6F, 0x63, 0x61, 0x6C, 0x53, 0x74, 0x6F, 0x72, 0x61,
			0x67, 0x65, 0x20, 0x69, 0x74, 0x65, 0x6D, 0x73, 0x20, 0x74, 0x6F, 0x20, 0x63,
			0x6C, 0x65, 0x61, 0x6E, 0x20, 0x75, 0x70, 0x20, 0x64, 0x65, 0x6C, 0x65, 0x74,
			0x65, 0x64, 0x20, 0x6F, 0x6E, 0x65, 0x73, 0x3A, 0x0A, 0x09, 0x66, 0x6F, 0x72,
			0x28, 0x20, 0x76, 0x61, 0x72, 0x20, 0x6B, 0x65, 0x79, 0x20, 0x69, 0x6E, 0x20,
			0x6C, 0x6F, 0x63, 0x61, 0x6C, 0x53, 0x74, 0x6F, 0x72, 0x61, 0x67, 0x65, 0x20,
			0x29, 0x20, 0x7B, 0x0A, 0x0A, 0x09, 0x09, 0x2F, 0x2F, 0x20, 0x69, 0x67, 0x6E,
			0x6F, 0x72, 0x69, 0x6E, 0x67, 0x20, 0x61, 0x6E, 0x79, 0x20, 0x66, 0x6F, 0x72,
			0x65, 0x69, 0x67, 0x6E, 0x20, 0x6B, 0x65, 0x79, 0x73, 0x3A, 0x0A, 0x09, 0x09,
			0x69, 0x66, 0x28, 0x20, 0x6B, 0x65, 0x79, 0x2E, 0x69, 0x6E, 0x64, 0x65, 0x78,
			0x4F, 0x66, 0x28, 0x27, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73, 0x78, 0x3A, 0x27,
			0x29, 0x20, 0x3D, 0x3D, 0x20, 0x2D, 0x31, 0x20, 0x29, 0x20, 0x7B, 0x0A, 0x09,
			0x09, 0x09, 0x63, 0x6F, 0x6E, 0x74, 0x69, 0x6E, 0x75, 0x65, 0x0A, 0x09, 0x09,
			0x7D, 0x0A, 0x0A, 0x09, 0x09, 0x2F, 0x2F, 0x20, 0x65, 0x78, 0x74, 0x72, 0x61,
			0x63, 0x74, 0x69, 0x6E, 0x67, 0x20, 0x73, 0x74, 0x6F, 0x72, 0x65, 0x64, 0x20,
			0x69, 0x74, 0x65, 0x6D, 0x20, 0x70, 0x61, 0x74, 0x68, 0x20, 0x66, 0x72, 0x6F,
			0x6D, 0x20, 0x6B, 0x65, 0x79, 0x3A, 0x0A, 0x09, 0x09, 0x76, 0x61, 0x72, 0x20,
			0x73, 0x74, 0x6F, 0x72, 0x65, 0x64, 0x50, 0x61, 0x74, 0x68, 0x20, 0x3D, 0x20,
			0x6B, 0x65, 0x79, 0x2E, 0x72, 0x65, 0x70, 0x6C, 0x61, 0x63, 0x65, 0x28, 0x2F,
			0x68, 0x74, 0x74, 0x70, 0x66, 0x73, 0x78, 0x3A, 0x5B, 0x5E, 0x3A, 0x5D, 0x2B,
			0x3F, 0x3A, 0x2F, 0x2C, 0x20, 0x27, 0x27, 0x29, 0x0A, 0x0A, 0x09, 0x09, 0x2F,
			0x2F, 0x20, 0x67, 0x6F, 0x74, 0x20, 0x69, 0x74, 0x65, 0x6D, 0x20, 0x66, 0x72,
			0x6F, 0x6D, 0x20, 0x6F, 0x74, 0x68, 0x65, 0x72, 0x20, 0x70, 0x61, 0x74, 0x68,
			0x2C, 0x20, 0x63, 0x61, 0x6E, 0x6E, 0x6F, 0x74, 0x20, 0x74, 0x6F, 0x75, 0x63,
			0x68, 0x20, 0x69, 0x74, 0x3A, 0x0A, 0x09, 0x09, 0x69, 0x66, 0x28, 0x20, 0x73,
			0x74, 0x6F, 0x72, 0x65, 0x64, 0x50, 0x61, 0x74, 0x68, 0x2E, 0x69, 0x6E, 0x64,
			0x65, 0x78, 0x4F, 0x66, 0x28, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6E, 0x74, 0x50,
			0x61, 0x74, 0x68, 0x29, 0x20, 0x21, 0x3D, 0x20, 0x30, 0x20, 0x29, 0x20, 0x7B,
			0x0A, 0x09, 0x09, 0x09, 0x63, 0x6F, 0x6E, 0x74, 0x69, 0x6E, 0x75, 0x65, 0x0A,
			0x09, 0x09, 0x7D, 0x0A, 0x0A, 0x09, 0x09, 0x2F, 0x2F, 0x20, 0x64, 0x6F, 0x65,
			0x73, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6E, 0x74, 0x20, 0x6C, 0x6F, 0x63,
			0x61, 0x6C, 0x53, 0x74, 0x6F, 0x72, 0x61, 0x67, 0x65, 0x20, 0x69, 0x74, 0x65,
			0x6D, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x20, 0x28, 0x6E, 0x6F, 0x74,
			0x20, 0x64, 0x65, 0x6C, 0x65, 0x74, 0x65, 0x64, 0x29, 0x3F, 0x0A, 0x09, 0x09,
			0x76, 0x61, 0x72, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x20, 0x3D, 0x20,
			0x66, 0x61, 0x6C, 0x73, 0x65, 0x0A, 0x0A, 0x09, 0x09, 0x2F, 0x2F, 0x20, 0x63,
			0x68, 0x65, 0x63, 0x6B, 0x69, 0x6E, 0x67, 0x2C, 0x20, 0x69, 0x66, 0x20, 0x63,
			0x75, 0x72, 0x72, 0x65, 0x6E, 0x74, 0x20, 0x6C, 0x6F, 0x63, 0x61, 0x6C, 0x53,
			0x74, 0x6F, 0x72, 0x61, 0x67, 0x65, 0x20, 0x69, 0x74, 0x65, 0x6D, 0x20, 0x69,
			0x73, 0x20, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6E, 0x74, 0x20, 0x6F, 0x6E, 0x20,
			0x63, 0x75, 0x72, 0x72, 0x65, 0x6E, 0x74, 0x20, 0x6C, 0x6F, 0x63, 0x61, 0x74,
			0x69, 0x6F, 0x6E, 0x20, 0x28, 0x70, 0x61, 0x67, 0x65, 0x29, 0x3A, 0x0A, 0x09,
			0x09, 0x66, 0x6F, 0x72, 0x28, 0x20, 0x76, 0x61, 0x72, 0x20, 0x69, 0x20, 0x3D,
			0x20, 0x30, 0x3B, 0x20, 0x69, 0x20, 0x3C, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74,
			0x69, 0x6E, 0x67, 0x50, 0x61, 0x74, 0x68, 0x73, 0x2E, 0x6C, 0x65, 0x6E, 0x67,
			0x74, 0x68, 0x3B, 0x20, 0x69, 0x2B, 0x2B, 0x20, 0x29, 0x20, 0x7B, 0x0A, 0x09,
			0x09, 0x09, 0x76, 0x61, 0x72, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6E,
			0x67, 0x50, 0x61, 0x74, 0x68, 0x20, 0x3D, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74,
			0x69, 0x6E, 0x67, 0x50, 0x61, 0x74, 0x68, 0x73, 0x5B, 0x69, 0x5D, 0x0A, 0x09,
			0x09, 0x09, 0x69, 0x66, 0x28, 0x20, 0x73, 0x74, 0x6F, 0x72, 0x65, 0x64, 0x50,
			0x61, 0x74, 0x68, 0x2E, 0x69, 0x6E, 0x64, 0x65, 0x78, 0x4F, 0x66, 0x28, 0x65,
			0x78, 0x69, 0x73, 0x74, 0x69, 0x6E, 0x67, 0x50, 0x61, 0x74, 0x68, 0x29, 0x20,
			0x3D, 0x3D, 0x20, 0x30, 0x20, 0x29, 0x20, 0x7B, 0x0A, 0x09, 0x09, 0x09, 0x09,
			0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x20, 0x3D, 0x20, 0x74, 0x72, 0x75, 0x65,
			0x0A, 0x09, 0x09, 0x09, 0x09, 0x62, 0x72, 0x65, 0x61, 0x6B, 0x0A, 0x09, 0x09,
			0x09, 0x7D, 0x0A, 0x09, 0x09, 0x7D, 0x0A, 0x0A, 0x09, 0x09, 0x2F, 0x2F, 0x20,
			0x72, 0x65, 0x6D, 0x6F, 0x76, 0x69, 0x6E, 0x67, 0x20, 0x64, 0x65, 0x6C, 0x65,
			0x74, 0x65, 0x64, 0x20, 0x66, 0x69, 0x6C, 0x65, 0x2D, 0x73, 0x79, 0x73, 0x74,
			0x65, 0x6D, 0x20, 0x69, 0x74, 0x65, 0x6D, 0x73, 0x20, 0x66, 0x72, 0x6F, 0x6D,
			0x20, 0x6C, 0x6F, 0x63, 0x61, 0x6C, 0x53, 0x74, 0x6F, 0x72, 0x61, 0x67, 0x65,
			0x3A, 0x0A, 0x09, 0x09, 0x69, 0x66, 0x28, 0x20, 0x21, 0x65, 0x78, 0x69, 0x73,
			0x74, 0x73, 0x20, 0x29, 0x20, 0x7B, 0x0A, 0x09, 0x09, 0x09, 0x6C, 0x6F, 0x63,
			0x61, 0x6C, 0x53, 0x74, 0x6F, 0x72, 0x61, 0x67, 0x65, 0x2E, 0x72, 0x65, 0x6D,
			0x6F, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6D, 0x28, 0x6B, 0x65, 0x79, 0x29, 0x0A,
			0x09, 0x09, 0x7D, 0x0A, 0x0A, 0x09, 0x7D, 0x0A, 0x0A, 0x09, 0x2F, 0x2F, 0x20,
			0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x69, 0x6E, 0x67, 0x20, 0x22, 0x73, 0x74,
			0x61, 0x72, 0x22, 0x20, 0x61, 0x6E, 0x64, 0x20, 0x22, 0x63, 0x6C, 0x65, 0x61,
			0x72, 0x2D, 0x73, 0x74, 0x6F, 0x72, 0x61, 0x67, 0x65, 0x22, 0x20, 0x63, 0x6C,
			0x69, 0x63, 0x6B, 0x73, 0x3A, 0x0A, 0x09, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73,
			0x78, 0x2E, 0x61, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6E, 0x74, 0x4C, 0x69, 0x73,
			0x74, 0x65, 0x6E, 0x65, 0x72, 0x28, 0x27, 0x63, 0x6C, 0x69, 0x63, 0x6B, 0x27,
			0x2C, 0x20, 0x66, 0x75, 0x6E, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x28, 0x65, 0x76,
			0x65, 0x6E, 0x74, 0x29, 0x20, 0x7B, 0x0A, 0x0A, 0x09, 0x09, 0x69, 0x66, 0x28,
			0x20, 0x65, 0x76, 0x65, 0x6E, 0x74, 0x2E, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
			0x2E, 0x63, 0x6C, 0x61, 0x73, 0x73, 0x4C, 0x69, 0x73, 0x74, 0x2E, 0x63, 0x6F,
			0x6E, 0x74, 0x61, 0x69, 0x6E, 0x73, 0x28, 0x27, 0x73, 0x74, 0x61, 0x72, 0x27,
			0x29, 0x20, 0x29, 0x20, 0x7B, 0x0A, 0x0A, 0x09, 0x09, 0x09, 0x76, 0x61, 0x72,
			0x20, 0x73, 0x74, 0x61, 0x72, 0x20, 0x3D, 0x20, 0x65, 0x76, 0x65, 0x6E, 0x74,
			0x2E, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x0A, 0x0A, 0x09, 0x09, 0x09, 0x76,
			0x61, 0x72, 0x20, 0x73, 0x74, 0x61, 0x72, 0x4B, 0x65, 0x79, 0x20, 0x3D, 0x20,
			0x73, 0x74, 0x61, 0x72, 0x2E, 0x67, 0x65, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69,
			0x62, 0x75, 0x74, 0x65, 0x28, 0x27, 0x64, 0x61, 0x74, 0x61, 0x2D, 0x68, 0x74,
			0x74, 0x70, 0x66, 0x73, 0x78, 0x2D, 0x73, 0x74, 0x61, 0x72, 0x2D, 0x6B, 0x65,
			0x79, 0x27, 0x29, 0x0A, 0x0A, 0x09, 0x09, 0x09, 0x2F, 0x2F, 0x20, 0x74, 0x6F,
			0x67, 0x67, 0x6C, 0x65, 0x20, 0x73, 0x74, 0x61, 0x72, 0x72, 0x69, 0x6E, 0x67,
			0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3A, 0x0A, 0x09, 0x09, 0x09, 0x69,
			0x66, 0x28, 0x20, 0x6C, 0x6F, 0x63, 0x61, 0x6C, 0x53, 0x74, 0x6F, 0x72, 0x61,
			0x67, 0x65, 0x2E, 0x67, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6D, 0x28, 0x73, 0x74,
			0x61, 0x72, 0x4B, 0x65, 0x79, 0x29, 0x20, 0x29, 0x20, 0x7B, 0x0A, 0x09, 0x09,
			0x09, 0x09, 0x6C, 0x6F, 0x63, 0x61, 0x6C, 0x53, 0x74, 0x6F, 0x72, 0x61, 0x67,
			0x65, 0x2E, 0x72, 0x65, 0x6D, 0x6F, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6D, 0x28,
			0x73, 0x74, 0x61, 0x72, 0x4B, 0x65, 0x79, 0x29, 0x0A, 0x09, 0x09, 0x09, 0x09,
			0x73, 0x74, 0x61, 0x72, 0x2E, 0x63, 0x6C, 0x61, 0x73, 0x73, 0x4C, 0x69, 0x73,
			0x74, 0x2E, 0x72, 0x65, 0x6D, 0x6F, 0x76, 0x65, 0x28, 0x27, 0x61, 0x63, 0x74,
			0x69, 0x76, 0x65, 0x27, 0x29, 0x0A, 0x09, 0x09, 0x09, 0x7D, 0x20, 0x65, 0x6C,
			0x73, 0x65, 0x20, 0x7B, 0x0A, 0x09, 0x09, 0x09, 0x09, 0x6C, 0x6F, 0x63, 0x61,
			0x6C, 0x53, 0x74, 0x6F, 0x72, 0x61, 0x67, 0x65, 0x2E, 0x73, 0x65, 0x74, 0x49,
			0x74, 0x65, 0x6D, 0x28, 0x73, 0x74, 0x61, 0x72, 0x4B, 0x65, 0x79, 0x2C, 0x20,
			0x27, 0x54, 0x27, 0x29, 0x0A, 0x09, 0x09, 0x09, 0x09, 0x73, 0x74, 0x61, 0x72,
			0x2E, 0x63, 0x6C, 0x61, 0x73, 0x73, 0x4C, 0x69, 0x73, 0x74, 0x2E, 0x61, 0x64,
			0x64, 0x28, 0x27, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x27, 0x29, 0x0A, 0x09,
			0x09, 0x09, 0x7D, 0x0A, 0x0A, 0x09, 0x09, 0x7D, 0x20, 0x65, 0x6C, 0x73, 0x65,
			0x20, 0x69, 0x66, 0x28, 0x20, 0x65, 0x76, 0x65, 0x6E, 0x74, 0x2E, 0x74, 0x61,
			0x72, 0x67, 0x65, 0x74, 0x2E, 0x63, 0x6C, 0x61, 0x73, 0x73, 0x4C, 0x69, 0x73,
			0x74, 0x2E, 0x63, 0x6F, 0x6E, 0x74, 0x61, 0x69, 0x6E, 0x73, 0x28, 0x27, 0x63,
			0x6C, 0x65, 0x61, 0x72, 0x2D, 0x73, 0x74, 0x6F, 0x72, 0x61, 0x67, 0x65, 0x27,
			0x29, 0x20, 0x29, 0x20, 0x7B, 0x0A, 0x0A, 0x09, 0x09, 0x09, 0x2F, 0x2F, 0x20,
			0x63, 0x6F, 0x6E, 0x66, 0x69, 0x72, 0x6D, 0x20, 0x61, 0x6E, 0x64, 0x20, 0x63,
			0x6C, 0x65, 0x61, 0x72, 0x20, 0x6C, 0x6F, 0x63, 0x61, 0x6C, 0x53, 0x74, 0x6F,
			0x72, 0x61, 0x67, 0x65, 0x3A, 0x0A, 0x09, 0x09, 0x09, 0x69, 0x66, 0x28, 0x20,
			0x63, 0x6F, 0x6E, 0x66, 0x69, 0x72, 0x6D, 0x28, 0x27, 0x43, 0x6C, 0x65, 0x61,
			0x72, 0x20, 0x73, 0x74, 0x6F, 0x72, 0x61, 0x67, 0x65, 0x3F, 0x27, 0x29, 0x20,
			0x29, 0x20, 0x7B, 0x0A, 0x0A, 0x09, 0x09, 0x09, 0x09, 0x6C, 0x6F, 0x63, 0x61,
			0x6C, 0x53, 0x74, 0x6F, 0x72, 0x61, 0x67, 0x65, 0x2E, 0x63, 0x6C, 0x65, 0x61,
			0x72, 0x28, 0x29, 0x0A, 0x0A, 0x09, 0x09, 0x09, 0x09, 0x2F, 0x2F, 0x20, 0x61,
			0x70, 0x70, 0x6C, 0x79, 0x20, 0x6C, 0x6F, 0x6F, 0x73, 0x69, 0x6E, 0x67, 0x20,
			0x73, 0x74, 0x61, 0x72, 0x73, 0x20, 0x74, 0x6F, 0x20, 0x55, 0x49, 0x3A, 0x0A,
			0x09, 0x09, 0x09, 0x09, 0x76, 0x61, 0x72, 0x20, 0x73, 0x74, 0x61, 0x72, 0x73,
			0x20, 0x3D, 0x20, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73, 0x78, 0x2E, 0x71, 0x75,
			0x65, 0x72, 0x79, 0x53, 0x65, 0x6C, 0x65, 0x63, 0x74, 0x6F, 0x72, 0x41, 0x6C,
			0x6C, 0x28, 0x27, 0x2E, 0x73, 0x74, 0x61, 0x72, 0x27, 0x29, 0x0A, 0x09, 0x09,
			0x09, 0x09, 0x66, 0x6F, 0x72, 0x28, 0x20, 0x76, 0x61, 0x72, 0x20, 0x69, 0x20,
			0x3D, 0x20, 0x30, 0x3B, 0x20, 0x69, 0x20, 0x3C, 0x20, 0x73, 0x74, 0x61, 0x72,
			0x73, 0x2E, 0x6C, 0x65, 0x6E, 0x67, 0x74, 0x68, 0x3B, 0x20, 0x69, 0x2B, 0x2B,
			0x20, 0x29, 0x20, 0x7B, 0x0A, 0x09, 0x09, 0x09, 0x09, 0x09, 0x76, 0x61, 0x72,
			0x20, 0x73, 0x74, 0x61, 0x72, 0x20, 0x3D, 0x20, 0x73, 0x74, 0x61, 0x72, 0x73,
			0x5B, 0x69, 0x5D, 0x0A, 0x09, 0x09, 0x09, 0x09, 0x09, 0x73, 0x74, 0x61, 0x72,
			0x2E, 0x63, 0x6C, 0x61, 0x73, 0x73, 0x4C, 0x69, 0x73, 0x74, 0x2E, 0x72, 0x65,
			0x6D, 0x6F, 0x76, 0x65, 0x28, 0x27, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x27,
			0x29, 0x0A, 0x09, 0x09, 0x09, 0x09, 0x7D, 0x0A, 0x0A, 0x09, 0x09, 0x09, 0x09,
			0x61, 0x6C, 0x65, 0x72, 0x74, 0x28, 0x27, 0x53, 0x74, 0x6F, 0x72, 0x61, 0x67,
			0x65, 0x20, 0x63, 0x6C, 0x65, 0x61, 0x72, 0x65, 0x64, 0x27, 0x29, 0x0A, 0x09,
			0x09, 0x09, 0x7D, 0x0A, 0x0A, 0x09, 0x09, 0x7D, 0x0A, 0x0A, 0x09, 0x7D, 0x29,
			0x0A, 0x0A, 0x7D, 0x29, 0x28, 0x29, 0x0A,
		},
	}
	Files["/style.css"] = File{
		Contents: []byte{
			0x2E, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73, 0x78, 0x20, 0x7B, 0x0A, 0x09, 0x66,
			0x6F, 0x6E, 0x74, 0x2D, 0x66, 0x61, 0x6D, 0x69, 0x6C, 0x79, 0x3A, 0x20, 0x73,
			0x61, 0x6E, 0x73, 0x2D, 0x73, 0x65, 0x72, 0x69, 0x66, 0x3B, 0x0A, 0x7D, 0x0A,
			0x0A, 0x2E, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73, 0x78, 0x20, 0x61, 0x20, 0x7B,
			0x0A, 0x09, 0x74, 0x65, 0x78, 0x74, 0x2D, 0x64, 0x65, 0x63, 0x6F, 0x72, 0x61,
			0x74, 0x69, 0x6F, 0x6E, 0x3A, 0x20, 0x6E, 0x6F, 0x6E, 0x65, 0x3B, 0x0A, 0x7D,
			0x0A, 0x0A, 0x2E, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73, 0x78, 0x20, 0x2E, 0x68,
			0x65, 0x61, 0x64, 0x65, 0x72, 0x20, 0x7B, 0x0A, 0x09, 0x6D, 0x61, 0x72, 0x67,
			0x69, 0x6E, 0x3A, 0x20, 0x30, 0x3B, 0x0A, 0x09, 0x70, 0x61, 0x64, 0x64, 0x69,
			0x6E, 0x67, 0x3A, 0x20, 0x30, 0x2E, 0x35, 0x63, 0x6D, 0x20, 0x30, 0x3B, 0x0A,
			0x09, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x73, 0x69, 0x7A, 0x65, 0x3A, 0x20, 0x31,
			0x63, 0x6D, 0x3B, 0x0A, 0x7D, 0x0A, 0x0A, 0x2E, 0x68, 0x74, 0x74, 0x70, 0x66,
			0x73, 0x78, 0x20, 0x2E, 0x6C, 0x69, 0x73, 0x74, 0x20, 0x7B, 0x0A, 0x09, 0x6D,
			0x61, 0x72, 0x67, 0x69, 0x6E, 0x3A, 0x20, 0x30, 0x3B, 0x0A, 0x09, 0x70, 0x61,
			0x64, 0x64, 0x69, 0x6E, 0x67, 0x3A, 0x20, 0x30, 0x3B, 0x0A, 0x09, 0x6C, 0x69,
			0x73, 0x74, 0x2D, 0x73, 0x74, 0x79, 0x6C, 0x65, 0x3A, 0x20, 0x6E, 0x6F, 0x6E,
			0x65, 0x3B, 0x0A, 0x7D, 0x0A, 0x0A, 0x2E, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73,
			0x78, 0x20, 0x2E, 0x6C, 0x69, 0x73, 0x74, 0x20, 0x2E, 0x69, 0x74, 0x65, 0x6D,
			0x20, 0x2E, 0x73, 0x74, 0x61, 0x72, 0x20, 0x7B, 0x0A, 0x09, 0x64, 0x69, 0x73,
			0x70, 0x6C, 0x61, 0x79, 0x3A, 0x20, 0x69, 0x6E, 0x6C, 0x69, 0x6E, 0x65, 0x2D,
			0x62, 0x6C, 0x6F, 0x63, 0x6B, 0x3B, 0x0A, 0x09, 0x77, 0x69, 0x64, 0x74, 0x68,
			0x3A, 0x20, 0x31, 0x63, 0x6D, 0x3B, 0x0A, 0x09, 0x68, 0x65, 0x69, 0x67, 0x68,
			0x74, 0x3A, 0x20, 0x31, 0x63, 0x6D, 0x3B, 0x0A, 0x09, 0x66, 0x6F, 0x6E, 0x74,
			0x2D, 0x73, 0x69, 0x7A, 0x65, 0x3A, 0x20, 0x31, 0x63, 0x6D, 0x3B, 0x0A, 0x09,
			0x6C, 0x69, 0x6E, 0x65, 0x2D, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3A, 0x20,
			0x31, 0x63, 0x6D, 0x3B, 0x0A, 0x09, 0x74, 0x65, 0x78, 0x74, 0x2D, 0x61, 0x6C,
			0x69, 0x67, 0x6E, 0x3A, 0x20, 0x63, 0x65, 0x6E, 0x74, 0x65, 0x72, 0x3B, 0x0A,
			0x09, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6C, 0x2D, 0x61, 0x6C, 0x69,
			0x67, 0x6E, 0x3A, 0x20, 0x6D, 0x69, 0x64, 0x64, 0x6C, 0x65, 0x3B, 0x0A, 0x09,
			0x63, 0x6F, 0x6C, 0x6F, 0x72, 0x3A, 0x20, 0x23, 0x38, 0x30, 0x38, 0x30, 0x38,
			0x30, 0x3B, 0x0A, 0x7D, 0x0A, 0x0A, 0x2E, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73,
			0x78, 0x20, 0x2E, 0x6C, 0x69, 0x73, 0x74, 0x20, 0x2E, 0x69, 0x74, 0x65, 0x6D,
			0x20, 0x2E, 0x73, 0x74, 0x61, 0x72, 0x2E, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
			0x20, 0x7B, 0x0A, 0x09, 0x63, 0x6F, 0x6C, 0x6F, 0x72, 0x3A, 0x20, 0x23, 0x38,
			0x42, 0x30, 0x30, 0x30, 0x30, 0x3B, 0x0A, 0x7D, 0x0A, 0x0A, 0x2E, 0x68, 0x74,
			0x74, 0x70, 0x66, 0x73, 0x78, 0x20, 0x2E, 0x6C, 0x69, 0x73, 0x74, 0x20, 0x2E,
			0x69, 0x74, 0x65, 0x6D, 0x20, 0x2E, 0x6C, 0x69, 0x6E, 0x6B, 0x20, 0x7B, 0x0A,
			0x09, 0x64, 0x69, 0x73, 0x70, 0x6C, 0x61, 0x79, 0x3A, 0x20, 0x69, 0x6E, 0x6C,
			0x69, 0x6E, 0x65, 0x2D, 0x62, 0x6C, 0x6F, 0x63, 0x6B, 0x3B, 0x0A, 0x09, 0x77,
			0x69, 0x64, 0x74, 0x68, 0x3A, 0x20, 0x63, 0x61, 0x6C, 0x63, 0x28, 0x31, 0x30,
			0x30, 0x25, 0x20, 0x2D, 0x20, 0x31, 0x63, 0x6D, 0x20, 0x2D, 0x20, 0x30, 0x2E,
			0x35, 0x63, 0x6D, 0x29, 0x3B, 0x0A, 0x09, 0x6D, 0x69, 0x6E, 0x2D, 0x68, 0x65,
			0x69, 0x67, 0x68, 0x74, 0x3A, 0x20, 0x31, 0x63, 0x6D, 0x3B, 0x0A, 0x09, 0x70,
			0x61, 0x64, 0x64, 0x69, 0x6E, 0x67, 0x3A, 0x20, 0x30, 0x3B, 0x0A, 0x09, 0x76,
			0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6C, 0x2D, 0x61, 0x6C, 0x69, 0x67, 0x6E,
			0x3A, 0x20, 0x74, 0x6F, 0x70, 0x3B, 0x0A, 0x7D, 0x0A, 0x0A, 0x2E, 0x68, 0x74,
			0x74, 0x70, 0x66, 0x73, 0x78, 0x20, 0x2E, 0x6C, 0x69, 0x73, 0x74, 0x20, 0x2E,
			0x69, 0x74, 0x65, 0x6D, 0x20, 0x2E, 0x6C, 0x69, 0x6E, 0x6B, 0x20, 0x2E, 0x6E,
			0x61, 0x6D, 0x65, 0x20, 0x7B, 0x0A, 0x09, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x73,
			0x69, 0x7A, 0x65, 0x3A, 0x20, 0x30, 0x2E, 0x35, 0x63, 0x6D, 0x3B, 0x0A, 0x09,
			0x6C, 0x69, 0x6E, 0x65, 0x2D, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3A, 0x20,
			0x30, 0x2E, 0x36, 0x63, 0x6D, 0x3B, 0x0A, 0x09, 0x63, 0x6F, 0x6C, 0x6F, 0x72,
			0x3A, 0x20, 0x23, 0x31, 0x30, 0x31, 0x30, 0x31, 0x30, 0x3B, 0x0A, 0x7D, 0x0A,
			0x0A, 0x2E, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73, 0x78, 0x20, 0x2E, 0x6C, 0x69,
			0x73, 0x74, 0x20, 0x2E, 0x69, 0x74, 0x65, 0x6D, 0x20, 0x2E, 0x6C, 0x69, 0x6E,
			0x6B, 0x20, 0x2E, 0x6E, 0x61, 0x6D, 0x65, 0x3A, 0x3A, 0x61, 0x66, 0x74, 0x65,
			0x72, 0x20, 0x7B, 0x0A, 0x09, 0x63, 0x6F, 0x6C, 0x6F, 0x72, 0x3A, 0x20, 0x23,
			0x45, 0x30, 0x45, 0x30, 0x45, 0x30, 0x3B, 0x0A, 0x7D, 0x0A, 0x0A, 0x2E, 0x68,
			0x74, 0x74, 0x70, 0x66, 0x73, 0x78, 0x20, 0x2E, 0x6C, 0x69, 0x73, 0x74, 0x20,
			0x2E, 0x69, 0x74, 0x65, 0x6D, 0x2E, 0x64, 0x69, 0x72, 0x20, 0x2E, 0x6C, 0x69,
			0x6E, 0x6B, 0x20, 0x2E, 0x6E, 0x61, 0x6D, 0x65, 0x3A, 0x3A, 0x61, 0x66, 0x74,
			0x65, 0x72, 0x20, 0x7B, 0x0A, 0x09, 0x63, 0x6F, 0x6E, 0x74, 0x65, 0x6E, 0x74,
			0x3A, 0x20, 0x22, 0x20, 0xF0, 0x9F, 0x93, 0x81, 0x22, 0x3B, 0x0A, 0x7D, 0x0A,
			0x0A, 0x2E, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73, 0x78, 0x20, 0x2E, 0x6C, 0x69,
			0x73, 0x74, 0x20, 0x2E, 0x69, 0x74, 0x65, 0x6D, 0x20, 0x2E, 0x6C, 0x69, 0x6E,
			0x6B, 0x20, 0x2E, 0x6D, 0x65, 0x74, 0x61, 0x20, 0x7B, 0x0A, 0x09, 0x64, 0x69,
			0x73, 0x70, 0x6C, 0x61, 0x79, 0x3A, 0x20, 0x62, 0x6C, 0x6F, 0x63, 0x6B, 0x3B,
			0x0A, 0x09, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x73, 0x69, 0x7A, 0x65, 0x3A, 0x20,
			0x30, 0x2E, 0x33, 0x63, 0x6D, 0x3B, 0x0A, 0x09, 0x6C, 0x69, 0x6E, 0x65, 0x2D,
			0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3A, 0x20, 0x30, 0x2E, 0x36, 0x63, 0x6D,
			0x3B, 0x0A, 0x09, 0x63, 0x6F, 0x6C, 0x6F, 0x72, 0x3A, 0x20, 0x23, 0x38, 0x30,
			0x38, 0x30, 0x38, 0x30, 0x3B, 0x0A, 0x7D, 0x0A, 0x0A, 0x2E, 0x68, 0x74, 0x74,
			0x70, 0x66, 0x73, 0x78, 0x20, 0x2E, 0x66, 0x6F, 0x6F, 0x74, 0x65, 0x72, 0x20,
			0x7B, 0x0A, 0x09, 0x66, 0x6F, 0x6E, 0x74, 0x2D, 0x73, 0x69, 0x7A, 0x65, 0x3A,
			0x20, 0x30, 0x2E, 0x33, 0x63, 0x6D, 0x3B, 0x0A, 0x09, 0x6C, 0x69, 0x6E, 0x65,
			0x2D, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3A, 0x20, 0x30, 0x2E, 0x33, 0x63,
			0x6D, 0x3B, 0x0A, 0x09, 0x74, 0x65, 0x78, 0x74, 0x2D, 0x61, 0x6C, 0x69, 0x67,
			0x6E, 0x3A, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3B, 0x0A, 0x09, 0x70, 0x61,
			0x64, 0x64, 0x69, 0x6E, 0x67, 0x2D, 0x74, 0x6F, 0x70, 0x3A, 0x20, 0x31, 0x63,
			0x6D, 0x3B, 0x0A, 0x7D, 0x0A, 0x0A, 0x2E, 0x68, 0x74, 0x74, 0x70, 0x66, 0x73,
			0x78, 0x20, 0x2E, 0x66, 0x6F, 0x6F, 0x74, 0x65, 0x72, 0x20, 0x2E, 0x63, 0x6C,
			0x65, 0x61, 0x72, 0x2D, 0x73, 0x74, 0x6F, 0x72, 0x61, 0x67, 0x65, 0x20, 0x7B,
			0x0A, 0x09, 0x63, 0x6F, 0x6C, 0x6F, 0x72, 0x3A, 0x20, 0x23, 0x38, 0x30, 0x38,
			0x30, 0x38, 0x30, 0x3B, 0x0A, 0x7D, 0x0A,
		},
	}
}

// readerPool is a pool of *bytes.Reader.
var readerPool = sync.Pool{}

// Handler is a net/http.Handler, that serves files, embedded into this package.
var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("/", r.URL.Path)

	file, ok := Files[filePath]
	if !ok {
		http.NotFound(w, r)
		return
	}

	var buf *bytes.Reader
	if v := readerPool.Get(); v == nil {
		buf = bytes.NewReader(file.Contents)
	} else {
		buf = v.(*bytes.Reader)
		buf.Reset(file.Contents)
	}

	http.ServeContent(w, r, path.Base(filePath), ModTime, buf)
})