package system

import (
	"fmt"
	"os/exec"
	"strings"
)

func Memory() string {
	cmd1 := exec.Command("wmic", "OS", "get", "FreePhysicalMemory")
	cmd2 := exec.Command("wmic", "OS", "get", "TotalVisibleMemorySize")

	output1, err := cmd1.Output()
	if err != nil {
		return ""
	}

	output2, err := cmd2.Output()
	if err != nil {
		return ""
	}

	freeMemory := parseWmicOutput(string(output1))
	totalMemory := parseWmicOutput(string(output2))

	usedMemory := totalMemory - freeMemory
	percentFree := (freeMemory * 100) / totalMemory
	percentUsed := 100 - percentFree

	return fmt.Sprintf("%d/%d MB [%d%%]", usedMemory/1024, totalMemory/1024, percentUsed)
}

func parseWmicOutput(output string) uint64 {
	lines := strings.Split(output, "\n")
	if len(lines) >= 2 {
		value := strings.TrimSpace(lines[1])
		var mem uint64
		_, err := fmt.Sscan(value, &mem)
		if err == nil {
			return mem
		}
	}
	return 0
}
