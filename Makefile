
.PHONY: protobuf-gen
protobuf-gen:
	@protoc --proto_path=pkg --go_out=pkg --go_opt=paths=source_relative pb/portfolio.proto