// Copyright © 2017 budougumi0617. All rights reserved.

// Copyright © 2017 budougumi0617. All rights reserved.

package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

const (
	testdata = "./testdata"
)

func TestWalkFunc(t *testing.T) {
	var tests = []struct {
		root     string
		expected []string
	}{
		{
			testdata,
			[]string{""},
		},
	}

	for _, test := range tests {
		fmt.Println("Root is " + test.root)
		filepath.Walk(test.root, walkFunc)
		// got := parse(test.filename, strings.NewReader(test.input), test.lang, test.key)
		// if !reflect.DeepEqual(got, test.expected) {
		// 	t.Errorf("\nResult = %v\nExpected %v", got, test.expected)
		// }
	}
}
