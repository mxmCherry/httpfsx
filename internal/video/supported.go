package video

import "os/exec"

var supported bool

func init() {
	supported = exec.Command("avprobe", "-version").Run() == nil &&
		exec.Command("avconv", "-version").Run() == nil
}
