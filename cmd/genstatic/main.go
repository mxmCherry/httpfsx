package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
}

func run() error {
	const dirPerm = 0755
	const filePerm = 0644

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	src := flag.String("src", filepath.Join(wd, "public"), "public root dir")
	dst := flag.String("dst", filepath.Join(wd, "static"), "destination dir to write static.go file")
	genHandler := flag.Bool("gen-handler", true, "generate HTTP Handler func")
	flag.Parse()

	if err = os.MkdirAll(*dst, dirPerm); err != nil {
		return err
	}

	outFile, err := os.OpenFile(
		filepath.Join(*dst, "static.go"),
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		filePerm,
	)
	if err != nil {
		return err
	}
	if err = outFile.Truncate(0); err != nil {
		return err
	}

	return generate(outFile, *src, *genHandler)
}

func generate(out io.Writer, src string, genHandler bool) error {
	const hexDict = "0123456789ABCDEF"
	const hexBytesPerLine = 13 // results in 78 chars per line (+ newline)

	src = filepath.Clean(src)

	buf := make([]byte, hexBytesPerLine) // read buffer

	hexLineBufSize := 3 + hexBytesPerLine*(4+1+1) // 3 tabs for each line + "0x00" + "," + " " or "\n" for each byte
	hexLineBuf := make([]byte, hexLineBufSize)

	jsonBuf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(jsonBuf)

	hexLineBuf[0] = '\t'
	hexLineBuf[1] = '\t'
	hexLineBuf[2] = '\t'
	for i, lastI := 3, hexLineBufSize-1; i <= lastI; {

		hexLineBuf[i] = '0'
		i++

		hexLineBuf[i] = 'x'
		i += 1 + 2 // i++ and skip next 2 hex chars

		hexLineBuf[i] = ','
		i += 2 // i++ and skip next space or newline character (no point of assiging it here, as it may be overriden on hex line generation)
	}

	_ = encoder.Encode(src)
	quotedRoot := jsonBuf.Bytes()
	jsonBuf.Reset()

	_, _ = io.WriteString(out, "// Static is a generated package,\n")
	_, _ = io.WriteString(out, "// that embeds files from ")
	_, _ = out.Write(quotedRoot[0 : len(quotedRoot)-1])
	_, _ = io.WriteString(out, " directory.\n")
	_, _ = io.WriteString(out, "// WARNING!!! Don't edit this file manually!!!\n")
	_, _ = io.WriteString(out, "package static\n\n")

	if genHandler {
		_, _ = io.WriteString(out, "import (\n")
		_, _ = io.WriteString(out, "\t\"bytes\"\n")
		_, _ = io.WriteString(out, "\t\"net/http\"\n")
		_, _ = io.WriteString(out, "\t\"path\"\n")
		_, _ = io.WriteString(out, "\t\"sync\"\n")
		_, _ = io.WriteString(out, "\t\"time\"\n")
		_, _ = io.WriteString(out, ")\n\n")
	} else {
		_, _ = io.WriteString(out, "import \"time\"\n")
	}

	_, _ = io.WriteString(out, "// File is a file record.\n")
	_, _ = io.WriteString(out, "type File struct {\n")
	_, _ = io.WriteString(out, "\tContents []byte\n")
	_, _ = io.WriteString(out, "}\n\n")
	_, _ = io.WriteString(out, "// ModTime holds package modification (generation) time.\n")
	_, _ = fmt.Fprintf(out, "var ModTime = time.Unix(%d, 0)\n\n", time.Now().Unix())
	_, _ = io.WriteString(out, "// Files hold path-to-contents file mapping.\n")
	_, _ = io.WriteString(out, "var Files map[string]File\n\n")
	_, _ = io.WriteString(out, "func init() {\n")
	_, _ = io.WriteString(out, "\tFiles = map[string]File{}\n")

	err := filepath.Walk(src, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filePath != "." && strings.HasPrefix(filePath, ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if info.IsDir() {
			return nil
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}

		rel := path.Clean(strings.TrimPrefix(filePath, src))

		_, _ = io.WriteString(out, "\tFiles[")

		_ = encoder.Encode(rel)
		id := jsonBuf.Bytes()
		jsonBuf.Reset()

		_, _ = out.Write(id[0 : len(id)-1])
		_, _ = io.WriteString(out, "] = File{\n")
		_, _ = io.WriteString(out, "\t\tContents: []byte{\n")

		for {

			n, err := file.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}

			j := 3 // skip 3 tabs
			for i, lastN := 0, n-1; i <= lastN; i++ {
				b := buf[i]

				j += 2 // skip "0x"

				hexLineBuf[j] = hexDict[(b&0xF0)>>4]
				j++

				hexLineBuf[j] = hexDict[b&0x0F]
				j++

				hexLineBuf[j] = ','
				j++

				if i == lastN {
					hexLineBuf[j] = '\n'
				} else {
					hexLineBuf[j] = ' '
				}
				j++

			}

			_, _ = out.Write(hexLineBuf[0:j])
		}

		_, _ = io.WriteString(out, "\t\t},\n") // close Content: []byte{}
		_, _ = io.WriteString(out, "\t}\n")    // close File{}

		return nil
	})
	if err != nil {
		return err
	}

	_, _ = io.WriteString(out, "}\n")

	if genHandler {
		_, _ = io.WriteString(out, handlerCode)
	}

	return nil
}

const handlerCode = `
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
`
