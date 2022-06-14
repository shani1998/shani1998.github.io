# Docker Image Config
IMAGE_REGISTRY 		   ?= shani1998
IMAGE_NAME             := shani1998.github.io
DOCKER_GENERATOR_IMAGE := $(IMAGE_REGISTRY)/$(IMAGE_NAME):v0.0.1


.PHONY: vendor
vendor:
	@export GO111MODULE=on; go mod tidy; go mod vendor; unset GO111MODULE


.PHONY: docker-build
docker-build:
	@docker build -f $(DOCKERFILE_PATH)/Dockerfile -t $(DOCKER_GENERATOR_IMAGE) .

.PHONY: docker-push
docker-push:
	@docker push $(DOCKER_GENERATOR_IMAGE)

.PHONY: lint
lint:  ## Check go lint
	@docker run --rm \
		-v $(PWD):/go/src/$(PROJECT_ROOT) \
		-w /go/src/$(PROJECT_ROOT) \
		golangci/golangci-lint:v1.41.0-alpine golangci-lint run --config .golangci.yaml
