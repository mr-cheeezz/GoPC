package system

import (
	"fmt"
	"os/exec"
	"strings"
)

func OSName() string {
	cmd := exec.Command("wmic", "os", "get", "Caption")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) >= 2 {
		osName := strings.TrimSpace(lines[1])
		return osName
	} else {
		return ""
	}
}
