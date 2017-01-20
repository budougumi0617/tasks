// Copyright © 2017 budougumi0617. All rights reserved.

package main

import "io/ioutil"

// Root has Tasks are exists in subdirectories.
type Root struct {
	name string
}

// Projects returns project directories.
func (r *Root) Projects() []string {
	var p []string
	files, err := ioutil.ReadDir(r.name)
	if err != nil {
		return p
	}
	for _, f := range files {
		if f.IsDir() {
			p = append(p, f.Name())
		}
	}
	return p
}
