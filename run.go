// Copyright © 2017 budougumi0617. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"os"
)

var cmdRun = &Command{
	UsageLine: "run outputs results.",
}

func init() {
	cmdRun.Run = runRun
}

func runRun(cmd *Command, args []string) {
	fmt.Fprintln(stdout, "Called runRun.")
	flags := flag.NewFlagSet("run", flag.ContinueOnError)
	var (
		iFlag string
		oFlag string
	)

	// Set flags.
	flags.StringVar(&iFlag, "dir", "./", "Target root directory")
	flags.StringVar(&oFlag, "outdir", "./", "Target output directory")
	if err := flags.Parse(args); err != nil {
		os.Exit(ExitCodeParseFlagError)
	}

	r := Root{iFlag}
	projs := r.Projects()
	s := newSeek(nil)
	for _, p := range projs {

		s.Walk(p)
		fmt.Println(s.tasks)
	}
}
