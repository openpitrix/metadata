# Copyright 2017 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

FROM openpitrix/openpitrix-builder as builder

WORKDIR /go/src/openpitrix.io/metadata/
COPY . .

RUN apk add --no-cache util-linux

RUN mkdir -p /metadata_bin
RUN go generate openpitrix.io/metadata/pkg/version && \
	CGO_ENABLED=0 GOOS=linux GOBIN=/metadata_bin go install -ldflags '-w -s' -tags netgo openpitrix.io/metadata/cmd/...

RUN find /metadata_bin -type f -exec upx {} \;

FROM alpine:3.7
RUN apk add --update ca-certificates && update-ca-certificates

COPY --from=builder /usr/bin/nsenter        /usr/bin/

COPY --from=builder /metadata_bin/drone     /usr/local/bin/
COPY --from=builder /metadata_bin/frontgate /usr/local/bin/

RUN apk add --no-cache supervisor
COPY build/supervisord/supervisord.conf /etc/
COPY build/supervisord/start-supervisord.sh /usr/local/bin/
COPY build/supervisord/frontgate/start-frontgate.sh /usr/local/bin/
COPY build/supervisord/drone/start-drone.sh /usr/local/bin/

RUN mkdir -p /etc/supervisor.d
COPY build/supervisord/frontgate/frontgate.ini /etc/supervisor.d
COPY build/supervisord/drone/drone.ini /etc/supervisor.d

ENTRYPOINT ["start-supervisord.sh"]
