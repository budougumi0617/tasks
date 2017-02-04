// Copyright © 2017 budougumi0617. All rights reserved.

package main

// Language is type of Programing language.
type Language int

// List of Support languages.
const (
	invalid Language = iota
	csharp
	xaml
	config
)

var commentformat = map[Language][]string{
	invalid: []string{},
	csharp:  []string{"///", "//", "/*"},
	xaml:    []string{"<!--"},
	config:  []string{},
}
