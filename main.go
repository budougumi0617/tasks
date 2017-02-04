// Copyright © 2016 budougumi0617 All rights reserved.

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var stdin io.Reader = os.Stdout  // modified during testing
var stdout io.Writer = os.Stdout // modified during testing
var stderr io.Writer = os.Stderr // modified during testing

const (
	// ExitCodeOK has no error.
	ExitCodeOK = 0
	// ExitCodeParseFlagError has any error in flags.
	ExitCodeParseFlagError = 1
	// ExitCodeInputDirectoryError shows input flag is incorrect.
	ExitCodeInputDirectoryError = 2
)

// Commands lists the available commands and help topics.
var commands = []*Command{
	cmdRun,
	// cmdVersion,
}

func main() {
	args := os.Args[1:]
	var (
		vFlag bool
		iFlag string
		oFlag string
	)

	flags := flag.NewFlagSet("tasks", flag.ContinueOnError)
	flags.SetOutput(stderr)
	flags.BoolVar(&vFlag, "v", false, "Print tasks version")
	flags.StringVar(&iFlag, "dir", "./", "Target root directory")
	flags.StringVar(&oFlag, "outdir", "./", "Target output directory")
	if err := flags.Parse(args[1:]); err != nil {
		os.Exit(ExitCodeParseFlagError)
	}
	if vFlag {
		fmt.Fprintf(stdout, "Tasks version is %s\n", Fullversion())
		os.Exit(ExitCodeOK)
	}
	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Runnable() {
			cmd.Run(cmd, args[1:])
			os.Exit(ExitCodeOK)
		}
	}

	// TODO Add help
}
