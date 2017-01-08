// Copyright © 2016 budougumi0617. All rights reserved.

// Package version manages package information.
package version

// These values are overridden by Makefile.
var (
	// Version shows this program version.
	Version = "0.0.0"
	// Revision shows Git revision.
	Revision = "-"
)

// Fullversion shows version and revision
func Fullversion() string {
	return Version + "-" + Revision
}
