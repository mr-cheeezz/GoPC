package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/mr-cheeezz/GoPC/content"
)

func PrintSpecs() {
	header("Hardware")
	miniHeader("PC")
	for part, name := range content.Specs {
		color.Blue("%s: %s", part, name())
	}
	miniHeader("Physical")
	for part, funcs := range content.Physical {
		for i, name := range funcs {
			color.Blue("%s %d: %s", part, i+1, name())
		}
	}
	space()
	header("Software")
	for app, name := range content.SoftW {
		color.Blue("%s: %s", app, name())
	}
	space()
}

func CopySpecs() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to copy all the specs to the clipboard? [Y/n]: ")
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(strings.TrimSpace(text))

	switch text {
	case "", "y":
		var specs string
		for name, getter := range content.Specs {
			spec := getter()
			specs += fmt.Sprintf("%s: %s\n", name, spec)
		}

		if err := clipboard.WriteAll(specs); err != nil {
			color.Red("Failed to copy specs to clipboard")
			panic(err)
		} else {
			color.Green("Copied specs to clipboard.")
			time.Sleep(1 * time.Second)
		}
	default:
		color.Yellow("Not copying the specs to the clipboard.")
		time.Sleep(1 * time.Second)
	}
}

func miniHeader(title string) {
	totalLength := 40
	dashes := (totalLength - len(title)) / 2

	if len(title)%2 != 0 {
		dashes--
	}

	miniHeader := strings.Repeat(" ", dashes) + "[" + title + "]" + strings.Repeat(" ", dashes)

	color.Cyan(miniHeader)
}

func header(title string) {
	totalLength := 40
	dashes := (totalLength - len(title)) / 2

	header := strings.Repeat("-", dashes) + "[" + title + "]" + strings.Repeat("-", dashes)

	color.Blue(header)
}

func Share(repo string) {
	color.Magenta("Send this to your friends | %s", repo)
}

func space() {
	fmt.Println("")
}
