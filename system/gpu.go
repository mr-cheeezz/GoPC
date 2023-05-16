package system

import (
	"os/exec"
	"strings"
)

func Gpu() string {
	cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
	output, err := cmd.Output()
	if err != nil {
		return "" // or you can return a default value or an error message
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) > 1 {
		for _, line := range lines[1:] {
			if strings.TrimSpace(line) != "" {
				return strings.TrimSpace(line)
			}
		}
	}
	return "" // return empty string if no GPU name is found
}
