// Copyright © 2016 budougumi0617. All rights reserved.

package testdata

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Command is based on go/src/cmd/go/main.go.
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string)

	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	// TODO task comment.
	UsageLine string

	// Short is the short description shown in the 'go help' output.
	Short string

	/* TODO task comment */

	// Long is the long message shown in the 'go help <this-command>' output.
	Long string

	// Output streams.
	outStream, errStream io.Writer

	// Input steam.
	inStream io.Reader

	// subCommands lists the available commands and help topics.
	// The order here is the order in which they are printed by 'go help'.
	subCommands []*Command
}

// Name returns the command's name: the first word in the usage line.
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// Usage shows how to use a command.
func (c *Command) Usage() {
	fmt.Fprintf(c.errStream, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(c.errStream, "%s\n", strings.TrimSpace(c.Long))
	os.Exit(2)
}

// Runnable reports whether the command can be run; otherwise
// it is a documentation pseudo-command such as importpath.
func (c *Command) Runnable() bool {
	return c.Run != nil
}
