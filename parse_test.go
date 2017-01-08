// Copyright © 2017 budougumi0617. All rights reserved.

package main

import (
	"reflect"
	"strings"
	"testing"
)

const (
	csTodoName = "todo.cs"
	csWithTodo = `/// TODO OK line.
/// todo OK line.
todo = ngLine;
/// TO DO NG line.
/**
* TODO OK line.
*/
// TODO OK line.
    `

	csHackName = "hack.cs"
	csWithHack = `/// HACK OK line.
/// hack OK line.
hack = ngLing;
/// HA CK NG line.
/**
* HACK OK line.
*/
// HACK OK line.
    `
	xamlTodoName = "todo.xaml"
	xamlWithTodo = `<!-- TODO OK line. -->
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
		filename string
		input    string
		lang     Language
		key      []string
		expected []*Task
	}{
		{ // TODO tasks in .cs file.
			csTodoName,
			csWithTodo,
			csharp,
			[]string{
				"TODO",
			},
			[]*Task{
				{FileName: csTodoName, Num: 1, Desc: "/// TODO OK line."},
				{FileName: csTodoName, Num: 2, Desc: "/// todo OK line."},
				//&Task{FileName: csTodoName, Num: 6, Desc: "* TODO OK line."}, // TODO Actually, must be collect task from multiple lines.
				{FileName: csTodoName, Num: 8, Desc: "// TODO OK line."},
			},
		},
		{ // HACK tasks in .cs file.
			csHackName,
			csWithHack,
			csharp,
			[]string{
				"TODO",
				"HACK",
			},
			[]*Task{
				{FileName: csHackName, Num: 1, Desc: "/// HACK OK line."},
				{FileName: csHackName, Num: 2, Desc: "/// hack OK line."},
				//&Task{FileName: csHackName, Num: 6, Desc: "* HACK OK line."}, // TODO Actually, must be collect task from multiple lines.
				{FileName: csHackName, Num: 8, Desc: "// HACK OK line."},
			},
		},
		{ // TODO tasks in .xaml file.
			xamlTodoName,
			xamlWithTodo,
			xaml,
			[]string{
				"TODO",
			},
			[]*Task{
				{FileName: xamlTodoName, Num: 1, Desc: "<!-- TODO OK line. -->"},
				{FileName: xamlTodoName, Num: 2, Desc: "<!-- todo OK line. -->"},
				//&Task{FileName: xamlTodoName, Num: 6, Desc: "TODO OK line."}, // TODO Actually, must be collect task from multiple lines.
			},
		},
	}

	for _, test := range tests {

		got := parse(test.filename, strings.NewReader(test.input), test.lang, test.key)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("\nResult = %v\nExpected %v", got, test.expected)
		}
	}
}
