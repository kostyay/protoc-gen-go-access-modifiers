#!/usr/bin/env bash

set -e

OUTPUT_DIR=example

export PATH=./bin:${PATH}
protoc --go_out=${OUTPUT_DIR} --go_opt=paths=source_relative \
    --go-grpc_out=${OUTPUT_DIR} --go-grpc_opt=paths=source_relative \
    --go-access-modifiers_out=${OUTPUT_DIR} --go-access-modifiers_opt=paths=source_relative \
    -I protos \
    -I ${OUTPUT_DIR} \
    example.proto
go test -v ./...