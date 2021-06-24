#/bin/bash
protoc --proto_path=./src --go_out=pb --go-grpc_out=./pb ./src/*.proto