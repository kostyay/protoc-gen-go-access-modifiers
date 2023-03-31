#!/usr/bin/env bash

if ! command -v protoc-gen-go &> /dev/null; then
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
fi
