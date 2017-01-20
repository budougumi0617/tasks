package main

import (
	"io"
	"testing"
)

func TestCommand_Name(t *testing.T) {
	type fields struct {
		UsageLine string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"No have space", fields{"name"}, "name"},
		{"Include space", fields{"name invalidstring"}, "name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				UsageLine: tt.fields.UsageLine,
			}
			if got := c.Name(); got != tt.want {
				t.Errorf("Command.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_Usage(t *testing.T) {
	type fields struct {
		Run         func(cmd *Command, args []string)
		UsageLine   string
		Short       string
		Long        string
		outStream   io.Writer
		errStream   io.Writer
		inStream    io.Reader
		subCommands []*Command
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				Run:         tt.fields.Run,
				UsageLine:   tt.fields.UsageLine,
				Short:       tt.fields.Short,
				Long:        tt.fields.Long,
				outStream:   tt.fields.outStream,
				errStream:   tt.fields.errStream,
				inStream:    tt.fields.inStream,
				subCommands: tt.fields.subCommands,
			}
			c.Usage()
		})
	}
}

func TestCommand_Runnable(t *testing.T) {
	type fields struct {
		Run func(cmd *Command, args []string)
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Run is nil", fields{nil}, false},
		{"Run is not nil", fields{func(c *Command, a []string) {
			// Dummy function.
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				Run: tt.fields.Run,
			}
			if got := c.Runnable(); got != tt.want {
				t.Errorf("Command.Runnable() = %v, want %v", got, tt.want)
			}
		})
	}
}
