package video

import (
	"fmt"
	"time"
)

const durationFormat = "15:04:05.999999999"

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
