#!/usr/bin/env bash

#rm -f ../pkg/rpc/pb/*.pb.go
protoc --go_out=plugins=grpc:../pkg/rpc --proto_path=../pkg/rpc/protos ../pkg/rpc/protos/*.proto
