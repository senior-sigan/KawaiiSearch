.PHONY: proto
proto: proto_prepare
	protoc --proto_path=protobuf \
			--cpp_out=src/protobuf \
			--grpc_out src/protobuf \
			--plugin=protoc-gen-grpc=`which grpc_cpp_plugin` \
			protobuf/embedder.proto

.PHONY: proto_prepare
proto_prepare:
	rm -rf "src/protobuf"
	mkdir -p "src/protobuf"
