
.PHONY: vendor
vendor:
	@export GO111MODULE=on; go mod tidy; go mod vendor;unset GO111MODULE

.PHONY: protobuf-gen
protobuf-gen:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  pkg/pb/portfolio.proto