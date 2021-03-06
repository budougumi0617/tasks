// Copyright © 2017 budougumi0617. All rights reserved.

package main

// Language is type of Programing language.
type Language int

// List of Support languages.
const (
	csharp Language = iota + 1
	xaml
)

var commentformat = map[Language][]string{
	csharp: []string{"///", "//", "/*"},
	xaml:   []string{"<!--"},
}
