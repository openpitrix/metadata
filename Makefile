# Copyright 2018 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

default:

dev:
	git describe --tags --always > ./_version
	git describe --exact-match || echo latest > ./_version

docker:
	docker build -t openpitrix/metadata-dev -f ./Dockerfile .
	docker images openpitrix/metadata-dev

generate:
	cd api && make

tools:
	# 1. install protoc from https://github.com/protocolbuffers/protobuf/releases
	# 2. install Go1.11+

	go get github.com/golang/protobuf/protoc-gen-go@v1.2
	go install openpitrix.io/metadata/build/protoc-gen-openpitrix-metadata-frontgate

clean:
	cd api && make clean
