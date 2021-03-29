genapi:
	protoc --proto_path=grpc -I=/usr/include -I=protos --go_out=pkg/gen --go_opt=paths=source_relative --go-grpc_out=pkg/gen --go-grpc_opt=paths=source_relative --dart_out=grpc:lib grpc/*.proto protos/*.proto
gengoogle:
	protoc -I/usr/include --dart_out=grpc:lib /usr/include/google/protobuf/*.proto