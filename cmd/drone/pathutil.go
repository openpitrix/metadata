// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"
)

func pathutil_MakeAbsPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	if dir, err := os.Getwd(); err == nil {
		path = filepath.Join(dir, path)
	}
	return path
}
