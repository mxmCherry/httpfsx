package mime

import (
	"net/http"
	"os"
)

func Detect(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	buf := make([]byte, 512)
	n, err := f.Read(buf)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(buf[0:n]), nil
}
