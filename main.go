package main

import (
	"fmt"

	"github.com/mr-cheeezz/GoPC/handler"
)

const repo = "https://github.com/mr-cheeezz/GoPC"

func main() {
	fmt.Println("")
	handlers()
}

func handlers() {
	handler.PrintSpecs()
	handler.CopySpecs()
}
