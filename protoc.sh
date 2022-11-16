#!/bin/bash

pushd internal/rpc > /dev/null 2>&1 || exit > /dev/null 2>&1

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/pb.proto

popd > /dev/null 2>&1 || exit > /dev/null 2>&1
