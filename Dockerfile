# Copyright 2018 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

# -----------------------------------------------------------------------------
# builder
# -----------------------------------------------------------------------------

FROM golang:1.11-alpine3.7 as builder

# intall tools
RUN apk add --no-cache git upx

# install /usr/bin/nsenter
RUN apk add --no-cache util-linux

WORKDIR /build-dir
COPY . .

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOBIN=/build-dir

RUN echo module drone > /build-dir/go.mod
RUN git describe --tags --always > /build-dir/version
RUN git describe --exact-match 2>/dev/null || git log -1 --format="%H" > /build-dir/version

RUN go get -tags netgo openpitrix.io/metadata/cmd/drone@$(cat /build-dir/version)
RUN go get -tags netgo openpitrix.io/metadata/cmd/frontgate@$(cat /build-dir/version)

RUN find /build-dir -type f -exec upx {} \;

RUN echo version: $(cat /build-dir/version)

# -----------------------------------------------------------------------------
# for metadata image
# -----------------------------------------------------------------------------

FROM alpine:3.7

COPY --from=builder /usr/bin/nsenter     /usr/bin/

COPY --from=builder /build-dir/drone     /usr/local/bin/
COPY --from=builder /build-dir/frontgate /usr/local/bin/

RUN apk add --no-cache supervisor
COPY build/supervisord/supervisord.conf             /etc/
COPY build/supervisord/start-supervisord.sh         /usr/local/bin/
COPY build/supervisord/frontgate/start-frontgate.sh /usr/local/bin/
COPY build/supervisord/drone/start-drone.sh         /usr/local/bin/

RUN mkdir -p /etc/supervisor.d
COPY build/supervisord/frontgate/frontgate.ini /etc/supervisor.d
COPY build/supervisord/drone/drone.ini         /etc/supervisor.d

ENTRYPOINT ["start-supervisord.sh"]

# -----------------------------------------------------------------------------
# END
# -----------------------------------------------------------------------------
