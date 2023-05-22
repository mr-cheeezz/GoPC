package system

import (
	"fmt"
	"os/exec"
	"strings"
)

type Win32_BaseBoard struct {
	Manufacturer string
	Product      string
}

func MotherBoard() string {
	cmd := exec.Command("wmic", "baseboard", "get", "Product")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) >= 2 {
		motherboard := strings.TrimSpace(lines[1])
		return motherboard
	} else {
		return ""
	}
}
