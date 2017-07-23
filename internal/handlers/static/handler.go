//go:generate gogenstatic

package static

import (
	"net/http"

	"github.com/mxmCherry/httpfsx/internal/handlers/static/static"
)

func New() http.Handler {
	return static.Handler()
}
