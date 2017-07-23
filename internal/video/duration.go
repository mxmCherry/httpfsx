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
		panic("video init: " + err.Error()) // impossible
	}
}

func parseDuration(s string) (time.Duration, error) {
	t, err := time.Parse(durationFormat, s)
	if err != nil {
		return 0, fmt.Errorf("videotools: failed to parse duration %s: %s", s, err.Error())
	}
	return t.Sub(zeroTime), nil
}

func formatDuration(d time.Duration) string {
	return zeroTime.Add(d).Format(durationFormat)
}

func videoDuration(file string) (time.Duration, error) {
	out, err := exec.Command("avprobe", file).CombinedOutput()
	if err != nil {
		return 0, err
	}

	i := bytes.Index(out, durAnchor)
	if i < 0 {
		return 0, fmt.Errorf("video: failed to parse '%s' anchor in avprobe output", durAnchor)
	}
	out = out[i+durAnchorLen:]

	i = bytes.IndexAny(out, ",\r\n")
	if i < 0 {
		return 0, fmt.Errorf("video: failed to parse closing comma after '%s...' in avprobe output", durAnchor)
	}
	out = out[0:i]

	return parseDuration(string(out))
}
