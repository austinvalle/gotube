package main

import (
	"os"

	"github.com/moosebot/gotube/internal/cli"
)

func main() {
	cli := &cli.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
