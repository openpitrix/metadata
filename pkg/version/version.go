// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

//go:generate go run gen_helper.go
//go:generate go fmt

package version

import (
	"fmt"
	"io"
)

var (
	ShortVersion   = "dev"
	GitSha1Version = "git-sha1"
	BuildDate      = "2017-01-01"
)

func PrintVersionInfo() {
	fmt.Printf("Release OpVersion: %s\n", ShortVersion)
	fmt.Printf("Git Commit Hash: %s\n", GitSha1Version)
	fmt.Printf("Build Time: %s\n", BuildDate)
}

func FprintVersionInfo(w io.Writer) {
	fmt.Fprintf(w, "Release OpVersion: %s\n", ShortVersion)
	fmt.Fprintf(w, "Git Commit Hash: %s\n", GitSha1Version)
	fmt.Fprintf(w, "Build Time: %s\n", BuildDate)
}

func GetVersionString() string {
	return fmt.Sprintf("%s; git: %s; build time: %s", ShortVersion, GitSha1Version, BuildDate)
}
