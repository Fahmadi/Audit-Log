ARG BASE_BUILD_IMAGE=hub.indraproject.ir/hubproxy/library/
ARG BASE_RUNTIME_IMAGE=hub.indraproject.ir/hubproxy/library/

FROM ${BASE_BUILD_IMAGE}golang:1.19.1-buster AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go env -w GOPROXY=https://goproxy.mirrors.alibaba.ir/
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo ./cmd/core

FROM ${BASE_RUNTIME_IMAGE}alpine:latest
RUN apk update; apk add curl
WORKDIR /app
COPY swagger swagger
COPY --from=builder /app/core /app/core
CMD ["/app/core"]
