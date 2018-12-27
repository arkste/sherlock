FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
ENV GO111MODULE=on
ENV CGO_ENABLED=0
COPY . $GOPATH/src/arkste/sherlock/
WORKDIR $GOPATH/src/arkste/sherlock/
RUN go get -d -v
RUN go build -o /go/bin/sherlock

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/sherlock /go/bin/sherlock
ENTRYPOINT ["/go/bin/sherlock"]