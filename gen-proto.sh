#!/bin/sh


# go get github.com/gogo/protobuf/protoc-gen-gogofast
# go get github.com/gogo/protobuf/protoc-gen-gogofaster
# go get github.com/gogo/protobuf/protoc-gen-gogoslick

echo "GOPATH: ($GOPATH)"

protoc -I=. \
    -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2 \
    --proto_path=proto \
    --gogofaster_out=plugins=grpc:paxoskv \
    paxoskv.proto
