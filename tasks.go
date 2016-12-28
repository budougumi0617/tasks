// Copyright © 2016 budougumi0617 All rights reserved.

// The Tasks creates Task List Comments from files.
package main

import "os"

var (
	Version  string
	Revision string
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
