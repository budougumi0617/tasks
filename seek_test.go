package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const (
	testdatadir = "./testdata"
)

func TestWalkFunc(t *testing.T) {
	var tests = []struct {
		root     string
		expected []*Task
	}{
		{
			testdatadir,
			[]*Task{
				&Task{FileName: "LcdWrapper.cs"},
				&Task{FileName: "TestData.xaml"},
			},
		},
	}

	for _, test := range tests {
		fmt.Println("Root is " + test.root)
		// Set stub
		s := &seek{parse: func(filename string, r io.Reader, l Language, k []string) []*Task {
			return []*Task{&Task{FileName: filename}}
		}}
		filepath.Walk(test.root, s.walkFunc)
		if !reflect.DeepEqual(s.tasks, test.expected) {
			t.Errorf("\nResult = %v\nExpected %v", s.tasks, test.expected)
		}
	}
}

func Test_newSeek(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{[]string{"DUMMY", "DUMMY1"}, []string{"DUMMY", "DUMMY1"}},
		{nil, []string{"TODO", "FIXME", "UNDONE"}},
		{[]string{}, []string{"TODO", "FIXME", "UNDONE"}},
	}
	for _, test := range tests {
		got := newSeek(test.input)
		want := &seek{pattern: test.want, parse: parse}
		if !reflect.DeepEqual(got.pattern, want.pattern) {
			t.Errorf("\nResult = %v\nExpected %v", got, want)
		}
	}
}

func Test_getCode(t *testing.T) {
	tests := []struct {
		name string
		want io.Reader
	}{
		{"NoExistFile", nil},
	}
	for _, test := range tests {
		if got := getCode(test.name); !reflect.DeepEqual(got, test.want) {
			t.Errorf("getCode() = %v, want %v", got, test.want)
		}
	}
}

func Test_getType(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want Language
	}{
		{"Symple cs file", args{"symple.cs"}, csharp},
		{"Symple xml file", args{"symple.xml"}, xaml},
		{"Config file", args{".vimrc"}, config},
		{"Not include dor in file", args{"COMMIT_EDITMSG"}, invalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getType(tt.args.name); got != tt.want {
				t.Errorf("getType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_seek_walkFunc(t *testing.T) {
	type fields struct {
		tasks   []*Task
		pattern []string
		parse   func(s string, r io.Reader, l Language, k []string) []*Task
	}
	type args struct {
		path string
		info os.FileInfo
		err  error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &seek{
				tasks:   tt.fields.tasks,
				pattern: tt.fields.pattern,
				parse:   tt.fields.parse,
			}
			if err := s.walkFunc(tt.args.path, tt.args.info, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("seek.walkFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
