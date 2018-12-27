FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
ENV GO111MODULE=on
COPY . $GOPATH/src/arkste/sherlock/
WORKDIR $GOPATH/src/arkste/sherlock/
RUN go get -d -v
RUN go build -o /go/bin/sherlock

FROM scratch
COPY --from=builder /go/bin/sherlock /go/bin/sherlock
ENTRYPOINT ["/go/bin/sherlock"]