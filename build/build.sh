#!/usr/bin/env bash

echo "starting ..."

#rm -f ../pkg/rpc/pb/*.pb.go

protoc --go_out=plugins=grpc:../pkg/rpc --proto_path=../pkg/rpc/protos ../pkg/rpc/protos/*.proto
RESULT=$?
if [ "${RESULT}" != "0" ]; then
	exit
fi

RESULT=`protoc --go_out=plugins=grpc:../pkg/rpc --proto_path=../pkg/rpc/protos ../pkg/rpc/protos/models/*.proto`
RESULT=$?
if [ "${RESULT}" != "0" ]; then
	exit
fi


# generate rpc.json
./proto-json.sh --quiet

echo "ok"