// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package frontgateutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"openpitrix.io/logger"
	"openpitrix.io/metadata/pkg/pb/frontgate"
	"openpitrix.io/metadata/pkg/pb/types"
)

func MustLoadFrontgateConfig(path string) *pbtypes.FrontgateConfig {
	p, err := LoadFrontgateConfig(path)
	if err != nil {
		logger.Criticalf(nil, "%+v", err)
		os.Exit(1)
	}
	return p
}

func LoadFrontgateConfig(path string) (*pbtypes.FrontgateConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	p := new(pbtypes.FrontgateConfig)
	if err := json.Unmarshal(data, p); err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	if p.ConfdConfig == nil {
		p.ConfdConfig = &pbtypes.ConfdConfig{}
	}

	if p.ConfdConfig.ProcessorConfig == nil {
		p.ConfdConfig.ProcessorConfig = &pbtypes.ConfdProcessorConfig{}
	}
	if p.ConfdConfig.BackendConfig == nil {
		p.ConfdConfig.BackendConfig = &pbtypes.ConfdBackendConfig{}
	}

	p.ConfdConfig.ProcessorConfig.Confdir = strExtractingEnvValue(
		p.ConfdConfig.ProcessorConfig.Confdir,
	)

	return p, nil
}

func DialFrontgateService(host string, port int) (
	client *pbfrontgate.FrontgateServiceClient,
	err error,
) {
	c, err := pbfrontgate.DialFrontgateService(
		"tcp", fmt.Sprintf("%s:%d", host, port),
	)
	if err != nil {
		logger.Warnf(nil, "%+v", err)
		return nil, err
	}

	return c, nil
}

func strExtractingEnvValue(s string) string {
	if !strings.ContainsAny(s, "${}") {
		return s
	}

	env := os.Environ()
	if runtime.GOOS == "windows" {
		if os.Getenv("HOME") == "" {
			home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
			if home == "" {
				home = os.Getenv("USERPROFILE")
			}

			env = append(env, "HOME="+home)
		}

		if os.Getenv("PWD") == "" {
			pwd, _ := os.Getwd()
			env = append(env, "PWD="+pwd)
		}
	}

	for _, e := range env {
		if i := strings.Index(e, "="); i >= 0 {
			s = strings.Replace(s,
				fmt.Sprintf("${%s}", strings.TrimSpace(e[:i])),
				strings.TrimSpace(e[i+1:]),
				-1,
			)
		}
	}
	return s
}
