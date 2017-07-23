package video

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

const durationFormat = "15:04:05.999999999"

var (
	durAnchor    = []byte("Duration: ")
	durAnchorLen = len(durAnchor)
)

var zeroTime time.Time

func init() {
	var err error
	zeroTime, err = time.Parse(time.RFC3339, "0000-01-01T00:00:00Z")
	if err != nil {
		panic("videotools init: " + err.Error()) // impossible
	}
}

func ParseDuration(s string) (time.Duration, error) {
	t, err := time.Parse(durationFormat, s)
	if err != nil {
		return 0, fmt.Errorf("videotools: failed to parse duration %s: %s", s, err.Error())
	}
	return t.Sub(zeroTime), nil
}

func FormatDuration(d time.Duration) string {
	return zeroTime.Add(d).Format(durationFormat)
}

func Duration(filename string) (time.Duration, error) {
	out, err := exec.Command("avprobe", filename).CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("videothumb: failed to exec avprobe for %s: %s", filename, err.Error())
	}

	i := bytes.Index(out, durAnchor)
	if i < 0 {
		return 0, fmt.Errorf("videothumb: failed to parse '%s' anchor in avprobe output", durAnchor)
	}
	out = out[i+durAnchorLen:]

	i = bytes.IndexAny(out, ",\r\n")
	if i < 0 {
		return 0, fmt.Errorf("videothumb: failed to parse closing comma after '%s...' in avprobe output", durAnchor)
	}
	out = out[0:i]

	return ParseDuration(string(out))
}
