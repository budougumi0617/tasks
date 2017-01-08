// Copyright © 2017 budougumi0617. All rights reserved.

package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

func getCode(name string) io.Reader {
	f, err := os.Open(name)
	if err != nil {
		return nil
	}
	return f
}

func getType(name string) Language {
	ext := path.Ext(name)
	var l Language
	switch ext[1:] { // Return value of path.Ext includes "dot".
	case "cs":
		l = csharp
	case "xaml", "xml":
		l = xaml
	default:
		l = 0
	}
	return l
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	fmt.Printf("Seek %s, %s\n", path, info.Name())
	tasks := parse(info.Name(), getCode(path), getType(info.Name()), []string{"TODO"})
	fmt.Println(tasks) //TODO
	return nil
}
