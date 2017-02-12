//go:generate gogenstatic

package statichandler

import (
	"net/http"

	"github.com/mxmCherry/httpfsx/internal/statichandler/static"
)

func New() http.Handler {
	return static.Handler()
}
