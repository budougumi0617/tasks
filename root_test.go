// Copyright © 2017 budougumi0617. All rights reserved.

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestRoot_Projects(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			"Get sub directory.",
			fields{name: "tempdir"},
			[]string{"test"},
		},
		{
			"Get sub directories.",
			fields{name: "tempdir"},
			[]string{"test", "test2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arange
			rootDir, _ := ioutil.TempDir("", tt.fields.name)
			r := &Root{
				name: rootDir,
			}

			ioutil.TempFile(rootDir, "tempfile")
			for _, subdir := range tt.want {
				fpath := filepath.Join(rootDir, subdir)
				os.Mkdir(fpath, 0644)
			}

			// Act
			if got := r.Projects(); !reflect.DeepEqual(got, tt.want) {
				// Assert
				t.Errorf("Root.Projects() = %v, want %v", got, tt.want)
			}
		})
	}
}
