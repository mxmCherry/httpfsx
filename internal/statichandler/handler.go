//go:generate statik -src=./public

package statichandler

import (
	"net/http"

	"github.com/rakyll/statik/fs"

	_ "github.com/mxmCherry/httpfsx/internal/statichandler/statik"
)

func New() http.Handler {
	statikFS, err := fs.New()
	if err != nil {
		panic(err.Error())
	}
	return http.FileServer(statikFS)
}
