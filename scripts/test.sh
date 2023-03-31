#!/usr/bin/env bash

set -e

export PATH=./bin:${PATH}
protoc --go_out=. --go_opt=paths=source_relative \
    --go-private_out=. --go-private_opt=paths=source_relative \
    example/example.proto
go test -v ./...