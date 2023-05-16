package handler

import (
	"fmt"
	"strings"
)

func PrintSpecs() {
	header("Hardware")

}

func print(str string) {
	fmt.Println(fmt.Sprintf(str))
}

func header(title string) {
	totalLength := 30
	dashes := (totalLength - len(title)) / 2

	header := strings.Repeat("-", dashes) + "[" + title + "]" + strings.Repeat("-", dashes)

	fmt.Println(header)
}
