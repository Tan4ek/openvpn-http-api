FROM golang:1.17.8-alpine as builder

WORKDIR /opt/openvpn-http-api

COPY go.mod go.sum ./
RUN go mod download

COPY config config
COPY ovpn   ovpn
COPY server server
COPY *.go   ./

RUN go build app.go

FROM kylemanna/openvpn

RUN apk add --update supervisor && \
  rm -rf /tmp/* /var/tmp/* /var/cache/apk/* /var/cache/distfiles/*

COPY supervisord.conf /etc/supervisord.conf

WORKDIR /opt/openvpn-http-api
COPY --from=builder /opt/openvpn-http-api/app .

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
