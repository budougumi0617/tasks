package main

import (
	"io"
	"testing"
)

func TestCLI_Run(t *testing.T) {
	type fields struct {
		outStream io.Writer
		errStream io.Writer
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CLI{
				outStream: tt.fields.outStream,
				errStream: tt.fields.errStream,
			}
			if got := c.Run(tt.args.args); got != tt.want {
				t.Errorf("CLI.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
