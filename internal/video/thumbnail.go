package video

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"time"
)

type ThumbnailOptions struct {
	MaxWidth, MaxHeight uint
	// ImageQuality        float64 // TODO
	Offset float64
}

func Thumbnail(ctx context.Context, w io.Writer, file string, opt *ThumbnailOptions) error {
	if opt.Offset > 1 || opt.Offset < -1 {
		return fmt.Errorf("video: offset must be within -1..1 range")
	}
	if opt.Offset < 0 {
		opt.Offset = 1 - opt.Offset
	}

	dur, err := videoDuration(ctx, file)
	if err != nil {
		return err
	}
	offsetDur := time.Duration(float64(dur) * opt.Offset)

	maxWidth := strconv.FormatUint(uint64(opt.MaxWidth), 10)
	maxHeight := strconv.FormatUint(uint64(opt.MaxHeight), 10)

	cmd := exec.CommandContext(
		ctx,
		"avconv",
		"-i", file,
		"-ss", formatDuration(offsetDur),
		"-vframes", "1",
		"-vf", "scale=w="+maxWidth+":h="+maxHeight+":force_original_aspect_ratio=decrease",
		"-y",
		"pipe:thumb.jpg",
	)
	cmd.Stdout = w
	return cmd.Run()
}
