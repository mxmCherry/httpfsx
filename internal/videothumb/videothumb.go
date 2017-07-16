package videothumb

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"time"
)

var (
	durAnchor    = []byte("Duration: ")
	durAnchorLen = len(durAnchor)
)

const durationFormat = "15:04:05.999999999"

var zeroTime time.Time

func init() {
	var err error
	zeroTime, err = time.Parse(time.RFC3339, "0000-01-01T00:00:00Z")
	if err != nil {
		panic(err.Error()) // impossible
	}
}

func Duration(name string) (time.Duration, error) {
	out, err := exec.Command("avprobe", name).CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("videothumb: failed to exec avprobe for %s: %s", name, err.Error())
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

	return parseDuration(string(out))
}

/*

	err := videothumb.Thumbnail(
		os.Stdout,
		"path/to/video.mp4",
		50*time.Second,
		100,
		100,
	)

*/

func Thumbnail(w io.Writer, name string, offset time.Duration, maxWidth, maxHeight uint) error {
	maxWidthStr := strconv.FormatUint(uint64(maxWidth), 10)
	maxHeightStr := strconv.FormatUint(uint64(maxHeight), 10)

	cmd := exec.Command(
		"avconv",
		"-i", name,
		"-ss", formatDuration(offset),
		"-vframes", "1",
		"-vf", "scale=w="+maxWidthStr+":h="+maxHeightStr+":force_original_aspect_ratio=decrease",
		"-y",
		"pipe:thumb.jpg",
	)
	cmd.Stdout = w
	return cmd.Run()
}

// ----------------------------------------------------------------------------

func parseDuration(s string) (time.Duration, error) {
	t, err := time.Parse(durationFormat, s)
	if err != nil {
		return 0, err
	}
	return t.Sub(zeroTime), nil
}

func formatDuration(d time.Duration) string {
	return zeroTime.Add(d).Format(durationFormat)
}
