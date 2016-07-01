package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("printf: not enough arguments")
		os.Exit(1)
	}

	fmt.Print(strings.Join(args, " "))
}
