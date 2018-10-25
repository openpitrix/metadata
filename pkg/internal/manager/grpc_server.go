// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package manager

import (
	"context"
	"fmt"
	"net"
	"runtime/debug"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"openpitrix.io/logger"
	"openpitrix.io/metadata/pkg/version"
)

type checkerT func(ctx context.Context, req interface{}) error
type builderT func(ctx context.Context, req interface{}) interface{}

var (
	defaultChecker checkerT
	defaultBuilder builderT
)

type GrpcServer struct {
	ServiceName    string
	Port           int
	showErrorCause bool
	checker        checkerT
	builder        builderT
}

type RegisterCallback func(*grpc.Server)

func NewGrpcServer(serviceName string, port int) *GrpcServer {
	return &GrpcServer{
		ServiceName:    serviceName,
		Port:           port,
		showErrorCause: false,
		checker:        defaultChecker,
		builder:        defaultBuilder,
	}
}

func (g *GrpcServer) ShowErrorCause(b bool) *GrpcServer {
	g.showErrorCause = b
	return g
}

func (g *GrpcServer) WithChecker(c checkerT) *GrpcServer {
	g.checker = c
	return g
}

func (g *GrpcServer) WithBuilder(b builderT) *GrpcServer {
	g.builder = b
	return g
}

func (g *GrpcServer) Serve(callback RegisterCallback, opt ...grpc.ServerOption) {
	if ver := version.GetVersion(); ver != nil {
		logger.Infof(nil, "Release OpVersion: %s\n", ver.AppModVersion)
		logger.Infof(nil, "GoVersion: %s\n", ver.GoVersion)
	} else {
		logger.Infof(nil, "Release OpVersion: unknown\n")
	}

	logger.Infof(nil, "Service [%s] start listen at port [%d]", g.ServiceName, g.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Port))
	if err != nil {
		err = errors.WithStack(err)
		logger.Criticalf(nil, "failed to listen: %+v", err)
	}

	builtinOptions := []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             10 * time.Second,
			PermitWithoutStream: true,
		}),

		grpc_middleware.WithUnaryServerChain(
			grpc_validator.UnaryServerInterceptor(),
			//g.unaryServerLogInterceptor(),
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				if g.checker != nil {
					err = g.checker(ctx, req)
					if err != nil {
						return
					}
				}

				return handler(ctx, req)
			},
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				if g.builder != nil {
					req = g.builder(ctx, req)
				}
				return handler(ctx, req)
			},
			grpc_recovery.UnaryServerInterceptor(
				grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
					logger.Criticalf(nil, "GRPC server recovery with error: %+v", p)
					logger.Criticalf(nil, string(debug.Stack()))
					if e, ok := p.(error); ok {
						return e
					}
					return nil
				}),
			),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(
				grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
					logger.Criticalf(nil, "GRPC server recovery with error: %+v", p)
					logger.Criticalf(nil, string(debug.Stack()))
					if e, ok := p.(error); ok {
						return e
					}
					return nil
				}),
			),
		),
	}

	grpcServer := grpc.NewServer(append(opt, builtinOptions...)...)
	reflection.Register(grpcServer)
	callback(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		err = errors.WithStack(err)
		logger.Criticalf(nil, "%+v", err)
	}
}
