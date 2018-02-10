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
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		audioFlag    bool
		skipMetaFlag bool
		dir          string

		versionFlag bool
	)

	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.BoolVar(&audioFlag, "a", false, "download audio mp3 only")
	flags.BoolVar(&skipMetaFlag, "s", false, "skip the metadata edit step")
	flags.StringVar(&dir, "d", "", "specify the download directory")
	flags.BoolVar(&versionFlag, "v", false, "print version information")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if versionFlag {
		fmt.Println(Version)
		return ExitCodeOK
	}

	if audioFlag {
		fmt.Println("audio only detected")
	}
	if skipMetaFlag {
		fmt.Println("skip metadata detected")
	}
	if dir != "" {
		fmt.Printf("download directory: %v", dir)
	}

	return ExitCodeOK
}
