#!/bin/sh

cd grpcsample

git --version 2>&1 >/dev/null
GIT_IS_AVAILABLE=$?
if [ $GIT_IS_AVAILABLE -eq 0 ]; then 
git init
else
echo "git is not installled.Install git from the following url https://git-scm.com/book/en/v2/Getting-Started-Installing-Git"
fi

## generate protos for all models
cd ..
protoc --version 2>&1 >/dev/null
PROTOC_IS_AVAILABLE=$?
if [ $PROTOC_IS_AVAILABLE -eq 0 ]; then
     protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpcsample/protos/product.proto
else
echo "protoc; protocol buffer compiler is not stalled.Install from the following url https://github.com/protocolbuffers/protobuf/releases"
fi

cd grpcsample

 go mod init grpcsample

 go mod tidy

 go build ./...