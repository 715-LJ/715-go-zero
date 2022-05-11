
GOPATH:=$(shell go env GOPATH)

.PHONY: userProto
userProto:
	goctl rpc protoc user.proto --go_out=rpc --go-grpc_out=rpc --zrpc_out=.

.PHONY: userApi
userApi:
	goctl api go -api api/user.api -dir ./api

.PHONY: userModel
userModel:
	goctl model pg datasource --url="postgres://sase:admin123456@10.47.119.170:5432/sase?sslmode=disable" --table="user" --dir='./model'

.PHONY: menuModel
menuModel:
	goctl model pg datasource --url="postgres://sase:admin123456@10.47.119.170:5432/sase?sslmode=disable" --table="menu" --dir='./model'
