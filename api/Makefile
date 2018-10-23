# Copyright 2018 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

PWD:=$(shell pwd)
GOPATH:=$(shell go env GOPATH)

METADATA_TYPES_PBFILES=$(sort $(wildcard metadata/types/*.proto))
METADATA_METADATA_PBFILES=$(sort $(wildcard metadata/metadata/*.proto))
METADATA_DRONE_PBFILES=$(sort $(wildcard metadata/drone/*.proto))
METADATA_FRONTGATE_PBFILES=$(sort $(wildcard metadata/frontgate/*.proto))
METADATA_PILOT_PBFILES=$(sort $(wildcard metadata/pilot/*.proto))

generate: Makefile
	@mkdir -p ../pkg/pb
	@rm -rf ../pkg/pb/*

	protoc $(PROTOC_FLAGS) --go_out=$(GOPATH)/src ${METADATA_TYPES_PBFILES}
	go run $(PWD)/../build/replace -old=,omitempty -new= -dir=$(GOPATH)/src/openpitrix.io/metadata/pkg/pb/types

	protoc $(PROTOC_FLAGS) --go_out=plugins=grpc:$(GOPATH)/src ${METADATA_DRONE_PBFILES}
	protoc $(PROTOC_FLAGS) --openpitrix-metadata-frontgate_out=$(GOPATH)/src ${METADATA_FRONTGATE_PBFILES}
	protoc $(PROTOC_FLAGS) --go_out=plugins=grpc:$(GOPATH)/src ${METADATA_PILOT_PBFILES}

clean:
	rm -rf $(GOPATH)/src/openpitrix.io/metadata/pkg/pb
