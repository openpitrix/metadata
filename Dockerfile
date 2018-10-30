# Copyright 2018 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

FROM golang:1.11 as builder

WORKDIR /build-dir
COPY ./.git ./.git

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOBIN=/build-dir

RUN echo module drone > /build-dir/go.mod
RUN git describe --tags --always > /build-dir/version
RUN git describe --exact-match 2>/dev/null || echo latest > /build-dir/version

RUN go get -ldflags '-w -s' -tags netgo openpitrix.io/metadata/cmd/drone@$(cat /build-dir/version)
RUN go get -ldflags '-w -s' -tags netgo openpitrix.io/metadata/cmd/frontgate@$(cat /build-dir/version)

RUN echo version: $(cat /build-dir/version)

FROM alpine:3.7

COPY --from=builder /build-dir/drone     /usr/local/bin/
COPY --from=builder /build-dir/frontgate /usr/local/bin/

CMD ["sh"]

