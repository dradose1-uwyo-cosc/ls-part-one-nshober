package main

import (
	"ls-part-one-nshober/functions"
	"os"
)

func main() {
	args := os.Args[1:]
	useColor := functions.IsTerminal(os.Stdout)
	functions.SimpleLS(os.Stdout, args, useColor)
}
