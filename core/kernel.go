package system

import (
	"os/exec"
	"strings"
)

func KernelVer() string {
	cmd := exec.Command("wmic", "os", "get", "Version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) >= 2 {
		kernelVersion := strings.TrimSpace(lines[1])
		return kernelVersion
	} else {
		return ""
	}
}
