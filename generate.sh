#!/bin/bash

# Install protoc-gen-go if not already installedgo install google.golang.org/protobuf/cmd/protoc-gen-go@latest


# Create api directory if it doesn't exist
mkdir -p api

# Generate Go code from proto files
protoc --go_out=. --go_opt=paths=source_relative api/user.proto 