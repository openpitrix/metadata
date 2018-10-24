// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package manager

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"openpitrix.io/logger"
)

type GrpcServer struct {
	ServiceName string
	Port        int
}

type RegisterCallback func(*grpc.Server)

func NewGrpcServer(serviceName string, port int) *GrpcServer {
	return &GrpcServer{
		ServiceName: serviceName,
		Port:        port,
	}
}

func (g *GrpcServer) Serve(callback RegisterCallback, opt ...grpc.ServerOption) {
	logger.Infof(nil, "Service [%s] start listen at port [%d]", g.ServiceName, g.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Port))
	if err != nil {
		err = errors.WithStack(err)
		logger.Criticalf(nil, "failed to listen: %+v", err)
	}

	grpcServer := grpc.NewServer(opt...)
	reflection.Register(grpcServer)
	callback(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		err = errors.WithStack(err)
		logger.Criticalf(nil, "%+v", err)
	}
}
