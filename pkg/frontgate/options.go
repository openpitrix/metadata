// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package frontgate

import (
	"bytes"
	"encoding/json"
	"fmt"

	etcdv3_backend "openpitrix.io/libconfd/etcdv3/backend"
	"openpitrix.io/logger"
	"openpitrix.io/metadata/pkg/pb/frontgate"
	"openpitrix.io/metadata/pkg/pb/pilot"
	"openpitrix.io/metadata/pkg/pb/types"
)

type Options func(opt *pbtypes.FrontgateConfig)

func NewDefaultConfigString() string {
	var (
		id      = "frontgate-001"
		nodeId  = "frontgate-001-node-01"
		localIp = "localhost"
	)
	p := &pbtypes.FrontgateConfig{
		Id:         id,
		NodeId:     nodeId,
		Host:       localIp,
		ListenPort: pbfrontgate.Default_Default_FrontgateServicePort,
		PilotHost:  localIp,

		PilotPort: pbpilot.Default_Default_PilotServicePort,
		NodeList: []*pbtypes.FrontgateEndpoint{
			{
				FrontgateId: id,
				NodeIp:      localIp,
				NodePort:    pbfrontgate.Default_Default_FrontgateServicePort,
			},
		},

		EtcdConfig: &pbtypes.EtcdConfig{
			User:     "",
			Password: "",
			NodeList: []*pbtypes.EtcdEndpoint{
				{
					Host: localIp,
					Port: 2379,
				},
			},
		},

		ConfdConfig: &pbtypes.ConfdConfig{
			ProcessorConfig: &pbtypes.ConfdProcessorConfig{
				Confdir:       "${PWD}/confd",
				Interval:      10,
				Noop:          false,
				Prefix:        "",
				SyncOnly:      false,
				LogLevel:      "DEBUG",
				Onetime:       false,
				Watch:         true,
				KeepStageFile: false,
			},
			BackendConfig: &pbtypes.ConfdBackendConfig{
				Type:         etcdv3_backend.Etcdv3BackendType,
				Host:         []string{"localhost:2379"},
				User:         "",
				Password:     "",
				ClientCaKeys: "",
				ClientCert:   "",
				ClientKey:    "",
			},
		},

		LogLevel: logger.DebugLevel.String(),
	}

	data, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		panic(err) // unreachable
	}

	data = bytes.Replace(data, []byte("\n"), []byte("\r\n"), -1)
	return string(data)
}

func NewDefaultConfig() *pbtypes.FrontgateConfig {
	s := NewDefaultConfigString()

	p := new(pbtypes.FrontgateConfig)
	if err := json.Unmarshal([]byte(s), p); err != nil {
		logger.Errorf(nil, "%+v", err)
		panic(err) // unreachable
	}
	return p
}

func WithFrontgateId(id string) func(opt *pbtypes.FrontgateConfig) {
	return func(opt *pbtypes.FrontgateConfig) {
		opt.Id = id
	}
}

func WithListenPort(port int) func(opt *pbtypes.FrontgateConfig) {
	return func(opt *pbtypes.FrontgateConfig) {
		opt.ListenPort = int32(port)
	}
}

func WithPilotService(host string, port int) func(opt *pbtypes.FrontgateConfig) {
	return func(opt *pbtypes.FrontgateConfig) {
		opt.PilotHost = host
		opt.PilotPort = int32(port)
	}
}
func WithFrontgateNodeList(node ...*pbtypes.FrontgateEndpoint) func(opt *pbtypes.FrontgateConfig) {
	return func(opt *pbtypes.FrontgateConfig) {
		opt.NodeList = append([]*pbtypes.FrontgateEndpoint{}, node...)
	}
}

func pkgGetEtcdEndpointsFromConfig(cfg *pbtypes.EtcdConfig) (endpoints []string) {
	for _, node := range cfg.GetNodeList() {
		endpoints = append(endpoints,
			fmt.Sprintf("%s:%d", node.GetHost(), node.GetPort()),
		)
	}
	return endpoints
}
