package system

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type SpecGetterFunc func() string

func MonitorNum() string {
	cmd := exec.Command("wmic", "desktopmonitor", "get", "DeviceID")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	lines := strings.Split(string(output), "\n")
	numMonitors := len(lines) - 2
	return strconv.Itoa(numMonitors)
}

func GetMonitorResolution(i int) string {
	cmd := exec.Command("wmic", "path", "Win32_VideoController", "get", "CurrentHorizontalResolution,CurrentVerticalResolution")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 && !strings.Contains(line, "CurrentHorizontalResolution") {
			resolution := strings.Fields(line)
			if len(resolution) == 2 {
				return fmt.Sprintf("%sx%s", resolution[0], resolution[1])
			}
		}
	}
	return ""
}

func GetMonitorRefreshRate(i int) string {
	cmd := exec.Command("wmic", "path", "Win32_VideoController", "get", "CurrentRefreshRate")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 && !strings.Contains(line, "CurrentRefreshRate") {
			return line
		}
	}
	return ""
}
