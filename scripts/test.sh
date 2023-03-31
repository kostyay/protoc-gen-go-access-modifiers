#!/usr/bin/env bash

set -e

OUTPUT_DIR=example

export PATH=./bin:${PATH}
protoc --go_out=${OUTPUT_DIR} --go_opt=paths=source_relative \
    --go-private_out=${OUTPUT_DIR} --go-private_opt=paths=source_relative \
    -I protos \
    -I ${OUTPUT_DIR} \
    example.proto
go test -v ./...