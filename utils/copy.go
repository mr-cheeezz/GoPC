package utils

import "github.com/atotto/clipboard"

func ClipBoard(specs string) {
	err := clipboard.WriteAll(specs)
	if err != nil {
		panic(err)
	}
}

func sepcs() string {
	specs := "t"
	return specs
}
