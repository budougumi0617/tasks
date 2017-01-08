// Copyright © 2017 budougumi0617. All rights reserved.

package main

import (
	"reflect"
	"strings"
	"testing"
)

const (
	csWithTodo string = `
/// TODO OK line.
/// todo OK line.
todo = ngLine;
/// TO DO NG line.
/**
* TODO OK line.
*/
// TODO OK line.
    `

	csWithHack string = `
/// HACK OK line.
/// hack OK line.
hack = ngLing;
/// HA CK NG line.
/**
* HACK OK line.
*/
// HACK OK line.
    `

	xamlWithTodo string = `
<!-- TODO OK line. -->
<!-- todo OK line. -->
<!-- TO DO NG line. -->
<!--
TODO OK line.
-->
TODO NG line.
    `
)

func TestCorrectParse(t *testing.T) {
	var tests = []struct {
		input    string
		lang     Language
		key      []string
		expected []*Task
	}{
		{ // TODO tasks in .cs file.
			csWithTodo,
			csharp,
			[]string{
				"TODO",
			},
			[]*Task{
				&Task{FileName: "", Num: 1, Desc: "/// TODO OK line."},
				&Task{FileName: "", Num: 2, Desc: "/// todo OK line."},
				//&Task{FileName: "", Num: 6, Desc: "* TODO OK line."}, // TODO Actually, must be collect task from multiple lines.
				&Task{FileName: "", Num: 8, Desc: "// TODO OK line."},
			},
		},
		{ // TODO tasks in .xaml file.
			xamlWithTodo,
			xaml,
			[]string{
				"TODO",
			},
			[]*Task{
				&Task{FileName: "", Num: 1, Desc: "<!-- TODO OK line. -->"},
				&Task{FileName: "", Num: 2, Desc: "<!-- todo OK line. -->"},
				//&Task{FileName: "", Num: 6, Desc: "TODO OK line."}, // TODO Actually, must be collect task from multiple lines.
			},
		},
	}

	for _, test := range tests {

		got := parse(strings.NewReader(test.input), test.lang, test.key)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("\nResult = %v\nExpected %v", got, test.expected)
		}
	}
}
