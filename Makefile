run:
	protoc --proto_path=./pb --proto_path=. --go-grpc_out=. ./pb/c2c.proto

	protoc --proto_path=./pb --proto_path=. --go-grpc_out=. --go_out=. ./pb/user.proto