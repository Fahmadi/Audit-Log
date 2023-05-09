.PHONY: all install test buf prepare build run

install:
	@go mod tidy
	# https://docs.buf.build/tour/generate-go-code
	@GOPROXY=https://goproxy.cn
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/google/wire/cmd/wire@latest

test: install
	@go test ./...
	@echo "✅ tests done!"

buf:
	@export PATH="$$PATH:$$(go env GOPATH)/bin" && buf generate api/proto
	@echo "✅ buf done!"

prepare: install buf

wire:
	@export PATH="$$PATH:$$(go env GOPATH)/bin" && wire ./cmd/core/app

build: prepare test
	@go build -o ./dist/crs ./cmd/core

run:
	@go run ./cmd/core
