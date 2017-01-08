// Copyright © 2017 budougumi0617. All rights reserved.

package main

import (
	"bufio"
	"io"
	"strings"
)

// parse collects task comments from io.Reader.
func parse(r io.Reader, l Language, k []string) []*Task {
	tasks := []*Task{}

	lines := bufio.NewScanner(r)
	for lines.Scan() {
		buf := strings.TrimSpace(lines.Text())
		if isComment(buf, l) && hasKey(buf, k) {
			tasks = append(tasks, &Task{Desc: buf})
		}
	}
	return tasks
}

// Check a s is comment.
func isComment(s string, l Language) bool {
	cfs := commentformat[l]
	for _, cf := range cfs {
		if strings.Index(s, cf) == 0 {
			return true
		}
	}
	return false
}

// Check a s includes any key.
func hasKey(s string, keys []string) bool {
	for _, k := range keys {
		if strings.Contains(strings.ToUpper(s), strings.ToUpper(k)) {
			return true
		}
	}
	return false
}
