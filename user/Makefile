
GOPATH:=$(shell go env GOPATH)

.PHONY: userApi
userApi:
	goctl api go -api api/user.api -dir ./api

.PHONY: userProto
userProto:
	goctl rpc protoc rpc/user.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc
