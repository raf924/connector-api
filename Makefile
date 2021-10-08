genapi: genrpc
genmessages:
	capnp -I$(GOPATH)/pkg/mod/zombiezen.com/go/capnproto2@v2.18.2+incompatible/std -Icapnp compile --src-prefix=capnp -ojava:java/src/main/java/tech/raf924/connectorapi -ogo:pkg/connector capnp/packet.capnp
genrpc: genmessages
	capnp -I$(GOPATH)/pkg/mod/zombiezen.com/go/capnproto2@v2.18.2+incompatible/std -Icapnp compile --src-prefix=capnp -ogo:pkg/connector capnp/connector.capnp
build: genapi
	java\gradlew build -q -p java
#gendart:
#	protoc --proto_path=grpc -I=/usr/include -I=protos --dart_out=grpc:lib grpc/*.proto protos/*.proto
#genjava:
#	capnp -Icapnp compile --src-prefix=capnp -ojava:src/main/java/tech/raf924/connectorapi capnp/user.capnp capnp/packet.capnp capnp/command.capnp