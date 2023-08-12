#!/usr/bin/env bash

echo "starting ..."

#rm -f ../pkg/rpc/pb/*.pb.go
protoc --go_out=../pkg/rpc --proto_path=../pkg/rpc/protos  ../pkg/rpc/protos/*.proto
RESULT=$?
if [ "${RESULT}" != "0" ]; then
	exit
fi

protoc --go-grpc_out=../pkg/rpc --go-grpc_opt=require_unimplemented_servers=false --proto_path=../pkg/rpc/protos  ../pkg/rpc/protos/*.proto
RESULT=$?
if [ "${RESULT}" != "0" ]; then
	exit
fi

protoc --go_out=../pkg/rpc --proto_path=../pkg/rpc/protos ../pkg/rpc/protos/models/*.proto
RESULT=$?
if [ "${RESULT}" != "0" ]; then
	exit
fi


# generate rpc.json
./proto-json.sh --quiet

echo "ok"