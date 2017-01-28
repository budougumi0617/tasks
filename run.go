// Copyright © 2017 budougumi0617. All rights reserved.

package main

import "fmt"

var cmdRun = &Command{
	UsageLine: "run outputs results.",
}

func init() {
	cmdRun.Run = runRun
}

func runRun(cmd *Command, args []string) {
	fmt.Fprintln(stdout, "Called runRun.")
}
