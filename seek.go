// Copyright © 2017 budougumi0617. All rights reserved.

package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

// newSeek is instead of constructor.
func newSeek(p []string) *seek {
	if p == nil || len(p) == 0 {
		p = []string{"TODO", "FIXME", "UNDONE"}
	}
	return &seek{pattern: p, parse: parse}
}

// Walk collects task from rootPath.
func (s *seek) Walk(rootPath string) {
	fmt.Println("Walk in " + rootPath)
	filepath.Walk(rootPath, s.walkFunc)
}

func getCode(name string) io.Reader {
	f, err := os.Open(name)
	if err != nil {
		return nil
	}
	return f
}

func getType(name string) Language {
	if name[0] == '.' { // config file like a .vimrc etc...
		return config
	}
	ext := path.Ext(name)

	if len(ext) == 0 {
		return invalid
	}

	var l Language
	switch ext[1:] { // Return value of path.Ext includes "dot".
	case "cs":
		l = csharp
	case "xaml", "xml":
		l = xaml
	default:
		l = invalid
	}
	return l
}

type seek struct {
	tasks   []*Task
	pattern []string
	parse   func(s string, r io.Reader, l Language, k []string) []*Task
}

// type is path/filepath/WalkFunc
func (s *seek) walkFunc(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	if l := getType(info.Name()); l != 0 {
		s.tasks = append(s.tasks, s.parse(info.Name(), getCode(path), l, s.pattern)...)
		fmt.Printf("result = %v\n", s.tasks)
	}
	return nil
}
