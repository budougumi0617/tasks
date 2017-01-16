// Copyright © 2016 budougumi0617. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"io"
)

// CLI defines and outputs.
type CLI struct {
	outStream, errStream io.Writer
}

const (
	// ExitCodeOK has no error.
	ExitCodeOK = 0
	// ExitCodeParseFlagError has any error in flags.
	ExitCodeParseFlagError = 1
	// ExitCodeInputDirectoryError shows input flag is incorrect.
	ExitCodeInputDirectoryError = 2
)

// Run runs the Tasks command-line interface.  Typical usage is
//     os.Exit(cli.Run(os.Stdin, os.Stdout, os.Stderr, os.Args))
// All arguments must be non-nil, and args[0] is required.
func (c *CLI) Run(args []string) int {
	var (
		vFlag bool
		iFlag string
		oFlag string
	)
	flags := flag.NewFlagSet("tasks", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&vFlag, "v", false, "Print tasks version")
	flags.StringVar(&iFlag, "dir", "./", "Target root directory")
	flags.StringVar(&oFlag, "outdir", "./", "Target output directory")
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}
	if vFlag {
		fmt.Fprintf(c.outStream, "Tasks version is %s\n", Fullversion())
		return ExitCodeOK
	}

	fmt.Fprint(c.outStream, "Do something work...\n")
	return ExitCodeOK
}
