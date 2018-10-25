// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package version

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"rsc.io/goversion/version"
)

type Version struct {
	AppPath        string // /path/to/exe
	AppModPath     string // app go import path
	AppModVersion  string // app version, e.g. v0.1.0
	GoVersion      string // Go version (runtime.Version in the program)
	ModuleInfo     string // program's module information
	BoringCrypto   bool   // program uses BoringCrypto
	StandardCrypto bool   // program uses standard crypto (replaced by BoringCrypto)
	FIPSOnly       bool   // program imports "crypto/tls/fipsonly"
}

func GetVersion() *Version {
	if pkgVersion != nil {
		var q = *pkgVersion
		return &q
	}
	return nil
}

func GetVersionString() string {
	if pkgVersion != nil {
		return pkgVersion.AppModVersion
	}
	return ""
}

var pkgVersion = func() *Version {
	apppath, err := getAppPath()
	if err != nil {
		return nil
	}
	v, _ := ReadVersion(apppath)
	return v
}()

func ReadVersion(apppath string) (*Version, error) {
	exeVersion, err := version.ReadExe(apppath)
	if err != nil {
		return nil, err
	}

	modPath, modVersion := parseVersionInfo(exeVersion)

	v := &Version{
		AppPath:        apppath,
		AppModPath:     modPath,
		AppModVersion:  modVersion,
		GoVersion:      exeVersion.Release,
		ModuleInfo:     exeVersion.ModuleInfo,
		BoringCrypto:   exeVersion.BoringCrypto,
		StandardCrypto: exeVersion.StandardCrypto,
		FIPSOnly:       exeVersion.FIPSOnly,
	}

	return v, nil
}

func parseVersionInfo(v version.Version) (modPath, modVersion string) {
	for _, line := range strings.Split(strings.TrimSpace(v.ModuleInfo), "\n") {
		row := strings.Split(line, "\t")
		if len(row) >= 2 && row[0] == "path" {
			modPath = row[1]
		}
		if len(row) >= 3 && row[0] == "mod" {
			modVersion = row[2]
		}
	}
	return
}

/*

	AppPath        string // /path/to/exe
	AppModPath     string // app go import path
	AppVersion     string // app version, e.g. v0.1.0
*/

func getAppPath() (string, error) {
	prog := os.Args[0]
	p, err := filepath.Abs(prog)
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(p)
	if err == nil {
		if !fi.Mode().IsDir() {
			return p, nil
		}
		err = fmt.Errorf("winsvc.GetAppPath: %s is directory", p)
	}
	if filepath.Ext(p) == "" {
		p += ".exe"
		fi, err := os.Stat(p)
		if err == nil {
			if !fi.Mode().IsDir() {
				return p, nil
			}
			err = fmt.Errorf("winsvc.GetAppPath: %s is directory", p)
		}
	}
	return "", err
}
