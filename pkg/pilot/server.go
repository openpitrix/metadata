// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package pilot

import (
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"openpitrix.io/logger"
	"openpitrix.io/metadata/pkg/internal/manager"
	"openpitrix.io/metadata/pkg/pb/pilot"
	"openpitrix.io/metadata/pkg/pb/types"
	"openpitrix.io/metadata/pkg/pilot/pilotutil"
)

type Server struct {
	cfg           *pbtypes.PilotConfig
	pbTlsCfg      *pbtypes.PilotTLSConfig
	fgClientMgr   *FrontgateClientManager
	taskStatusMgr *TaskStatusManager
}

func Serve(cfg *pbtypes.PilotConfig, pbTlsCfg *pbtypes.PilotTLSConfig, opts ...Options) {
	if cfg != nil {
		cfg = proto.Clone(cfg).(*pbtypes.PilotConfig)
	} else {
		cfg = NewDefaultConfig()
	}

	for _, fn := range opts {
		fn(cfg)
	}

	p := &Server{
		cfg:           cfg,
		pbTlsCfg:      proto.Clone(pbTlsCfg).(*pbtypes.PilotTLSConfig),
		fgClientMgr:   NewFrontgateClientManager(),
		taskStatusMgr: NewTaskStatusManager(),
	}

	go func() {
		for {
			p.fgClientMgr.CheckAllClient()
			time.Sleep(time.Second * 10)
		}
	}()

	// internal service
	go manager.NewGrpcServer("pilot-service", int(p.cfg.ListenPort)).Serve(
		func(server *grpc.Server) {
			pbpilot.RegisterPilotServiceServer(server, p)
		},
	)

	// tls for public service
	if pbTlsCfg != nil {
		tlsCfg, err := pilotutil.NewServerTLSConfigFromPbConfig(pbTlsCfg)
		if err != nil {
			logger.Criticalf(nil, "%+v", err)
			os.Exit(1)
		}

		manager.NewGrpcServer("pilot-service-for-frontgate", int(p.cfg.TlsListenPort)).Serve(
			func(server *grpc.Server) {
				pbpilot.RegisterPilotServiceForFrontgateServer(server, p)
			},
			grpc.Creds(credentials.NewTLS(tlsCfg)),
		)
	} else {
		manager.NewGrpcServer("pilot-service-for-frontgate", int(p.cfg.TlsListenPort)).Serve(
			func(server *grpc.Server) {
				pbpilot.RegisterPilotServiceForFrontgateServer(server, p)
			},
		)
	}
}
