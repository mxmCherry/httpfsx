package video

import (
	"io"
	"os/exec"
	"strconv"
	"time"

	"github.com/mxmCherry/httpfsx/internal/video"

	_ "image/gif"
	_ "image/png"
)

// TODO: make offset a ratio (percentage) of total video time? Like, 0.5 == totalTime/2
// TODO: make something configurable to return empty thumbnail or so if there's no avconv installed.
func Thumbnail(w io.Writer, width, height uint, _jpegQuality int, offset time.Duration, filename string) error {
	widthStr := strconv.FormatUint(uint64(width), 10)
	heightStr := strconv.FormatUint(uint64(height), 10)

	cmd := exec.Command(
		"avconv",
		"-i", filename,
		"-ss", video.FormatDuration(offset),
		"-vframes", "1",
		"-vf", "scale=w="+widthStr+":h="+heightStr+":force_original_aspect_ratio=decrease",
		"-y",
		"pipe:thumb.jpg",
	)
	cmd.Stdout = w
	return cmd.Run()
}
