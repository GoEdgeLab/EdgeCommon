#!/usr/bin/env bash

echo "starting ..."

function assert() {
	RESULT=$?
	if [ "${RESULT}" != "0" ]; then
		exit
	fi
}

#rm -f ../pkg/rpc/pb/*.pb.go
protoc --go_out=../pkg/rpc --proto_path=../pkg/rpc/protos  ../pkg/rpc/protos/*.proto
assert

protoc --go-grpc_out=../pkg/rpc --go-grpc_opt=require_unimplemented_servers=false --proto_path=../pkg/rpc/protos  ../pkg/rpc/protos/*.proto
assert

protoc --go_out=../pkg/rpc --proto_path=../pkg/rpc/protos ../pkg/rpc/protos/models/*.proto
RESULT=$?
assert


# generate rpc.json
./proto-json.sh --quiet
assert

echo "ok"