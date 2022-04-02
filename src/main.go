package main

import (
	"go-wordle/src/cli"
	"os"
)

func main() {
	cli.SetupAndPlayGame(os.Args, os.Stdin, os.Stdout)
}
