// Copyright © 2017 budougumi0617. All rights reserved.

package main

import (
	"strconv"
)

// Task defines a task.
type Task struct {
	// FileName is name of founded file.
	FileName string
	// Num is number of a task.
	Num uint
	// Desc is comments of a task.
	Desc string
}

func (t *Task) String() string {
	return "\nTask[FileName: \"" + t.FileName + "\", Num: " + strconv.FormatUint(uint64(t.Num), 10) + ", Desc: \"" + t.Desc + "\"]"
}
