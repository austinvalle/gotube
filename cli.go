package main

import (
	"flag"
	"fmt"
	"io"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		audio    bool
		skipMeta bool
		dir      string

		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&audio, "audio", false, "download audio mp3 only")
	flags.BoolVar(&audio, "a", false, "download audio mp3 only(Short)")
	flags.BoolVar(&skipMeta, "skip-meta", false, "skip the metadata edit step")
	flags.BoolVar(&skipMeta, "s", false, "skip the metadata edit step(Short)")
	flags.StringVar(&dir, "dir", "", "specify the download directory")
	flags.StringVar(&dir, "d", "", "specify the download directory(Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	_ = audio

	_ = skipMeta

	_ = dir

	return ExitCodeOK
}
