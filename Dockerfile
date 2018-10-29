# Copyright 2018 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

FROM golang:1.11 as builder

WORKDIR /build-dir
COPY ./.git ./.git

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOBIN=/build-dir/bin

RUN echo module drone > /build-dir/go.mod
RUN git describe --tags --always > /build-dir/version

RUN go get -ldflags '-w -s' -tags netgo openpitrix.io/metadata/cmd/drone@$(cat /build-dir/version)
RUN go get -ldflags '-w -s' -tags netgo openpitrix.io/metadata/cmd/frontgate@$(cat /build-dir/version)

FROM alpine:3.7

COPY --from=builder /metadata_bin/drone     /usr/local/bin/
COPY --from=builder /metadata_bin/frontgate /usr/local/bin/

CMD ["sh"]

