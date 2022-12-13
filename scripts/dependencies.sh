#!/bin/sh
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.20.3/protoc-3.20.3-linux-x86_64.zip &&
unzip protoc-3.20.3-linux-x86_64.zip    &&
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  &&
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest