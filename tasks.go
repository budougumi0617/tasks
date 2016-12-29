// Copyright © 2016 budougumi0617 All rights reserved.

package main

import "os"

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
