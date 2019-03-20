package main

//go:generate protoc -I./protocol -I/usr/local/include -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./protocol --grpc-gateway_out=logtostderr=true,grpc_api_configuration=./protocol/count_service.yaml:./protocol --swagger_out=logtostderr=true,grpc_api_configuration=./protocol/count_service.yaml:./protocol --csharp_out=./protocol --grpc_out=./protocol --plugin=protoc-gen-grpc=/root/.nuget/packages/grpc.tools/1.19.0/tools/linux_x64/grpc_csharp_plugin count.proto
