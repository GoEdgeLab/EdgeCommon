#!/usr/bin/env bash

rm -f ../internal/rpc/pb/*
protoc --go_out=plugins=grpc:../pkg/rpc --proto_path=../pkg/rpc/protos ../pkg/rpc/protos/*.proto
