package main

import (
	"time"

	"github.com/mr-cheeezz/GoPC/handler"
)

func main() {
	handlers()
	time.Sleep(3 * time.Second)
}

func handlers() {
	handler.PrintSpecs()
}
