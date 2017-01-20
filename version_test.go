// Copyright © 2016 budougumi0617. All rights reserved.

package main

import "testing"

func TestFullversion(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Get correct name",
			want: Version + "-" + Revision,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fullversion(); got != tt.want {
				t.Errorf("fullversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
