package httptool

import (
	"net/http"
	"time"
)

// NotModified handles If-Modified-Since requests and serves them.
// If request was served, it returns true.
func NotModified(w http.ResponseWriter, r *http.Request, lastMod time.Time) bool {
	if ifMod := r.Header.Get("If-Modified-Since"); ifMod != "" {
		t, err := time.ParseInLocation(http.TimeFormat, ifMod, time.Local)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return true
		}

		// UNIX seconds comparison is done because in real life, os.Stat may return time with nanoseconds,
		// but If-Modified-Since comes in with 1 second precision:
		if lastMod.Unix() <= t.Unix() {
			w.WriteHeader(http.StatusNotModified)
			return true
		}
	}
	w.Header().Set("Last-Modified", lastMod.Format(http.TimeFormat))
	return false
}
