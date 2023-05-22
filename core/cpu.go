package system

import (
	"os/exec"
	"strings"
)

func Cpu() string {
	cmd := exec.Command("wmic", "cpu", "get", "Name")
	output, err := cmd.Output()
	if err == nil {
		lines := strings.Split(string(output), "\n")
		if len(lines) >= 2 {
			cpu := strings.TrimSpace(lines[1])
			return cpu
		}
	}
	return ""
}
