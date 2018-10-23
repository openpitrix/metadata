# Copyright 2018 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

# https://github.com/protocolbuffers/protobuf/releases
# go install github.com/golang/protobuf/protoc-gen-go@v1.2

PWD:=$(shell pwd)
GOPATH:=$(shell go env GOPATH)

METADATA_TYPES_PBFILES=$(sort $(wildcard metadata/types/*.proto))
METADATA_METADATA_PBFILES=$(sort $(wildcard metadata/metadata/*.proto))
METADATA_DRONE_PBFILES=$(sort $(wildcard metadata/drone/*.proto))
METADATA_FRONTGATE_PBFILES=$(sort $(wildcard metadata/frontgate/*.proto))
METADATA_PILOT_PBFILES=$(sort $(wildcard metadata/pilot/*.proto))

generate: Makefile
	mkdir -p ../pkg/pb
	rm -rf ../pkg/pb/*

	make metadata_types
	make metadata_drone
	make metadata_frontgate
	make metadata_pilot

	@echo "ok"

metadata_types: ${METADATA_TYPES_PBFILES}
	mkdir -p $(GOPATH)/src/openpitrix.io/metadata/pkg/pb/types
	rm -rf   $(GOPATH)/src/openpitrix.io/metadata/pkg/pb/types/*
	protoc $(PROTOC_FLAGS) --go_out=$(GOPATH)/src ${METADATA_TYPES_PBFILES}
	ls $(PWD)/../build/replace
	go run $(PWD)/../build/replace/main.go -old=,omitempty -new= -dir=$(GOPATH)/src/openpitrix.io/metadata/pkg/pb/types
	@echo "ok"

metadata_drone: ${METADATA_DRONE_PBFILES}
	mkdir -p $(GOPATH)/src/openpitrix.io/metadata/pkg/pb/drone
	rm -rf   $(GOPATH)/src/openpitrix.io/metadata/pkg/pb/drone/*
	protoc $(PROTOC_FLAGS) --go_out=plugins=grpc:$(GOPATH)/src ${METADATA_DRONE_PBFILES}
	@echo "ok"

metadata_frontgate: ${METADATA_FRONTGATE_PBFILES}
	mkdir -p $(GOPATH)/src/openpitrix.io/metadata/pkg/pb/frontgate
	rm -rf   $(GOPATH)/src/openpitrix.io/metadata/pkg/pb/frontgate/*
	protoc $(PROTOC_FLAGS) --stdrpc_out=$(GOPATH)/src ${METADATA_FRONTGATE_PBFILES}
	@echo "ok"

metadata_pilot: ${METADATA_PILOT_PBFILES}
	mkdir -p $(GOPATH)/src/openpitrix.io/metadata/pkg/pb/pilot
	rm -rf   $(GOPATH)/src/openpitrix.io/metadata/pkg/pb/pilot/*
	protoc $(PROTOC_FLAGS) --go_out=plugins=grpc:$(GOPATH)/src ${METADATA_PILOT_PBFILES}
	@echo "ok"

tools:
	go get github.com/golang/protobuf/protoc-gen-go@v1.2

clean:
	rm -rf $(GOPATH)/src/openpitrix.io/metadata/pkg/pb
	@echo "ok"
