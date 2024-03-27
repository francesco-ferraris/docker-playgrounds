IMAGE_NAME=docker-playground
TAG=base

.PHONY: build
build:
	@docker build -t $(IMAGE_NAME):$(TAG) .

.PHONY: run
run:
	@docker run --rm -p 8080:8080 $(IMAGE_NAME):$(TAG)

.PHONY: build-and-run
build-and-run: build run

.PHONY: clean
clean:
	@docker image rm $(IMAGE_NAME):$(TAG)

.DEFAULT_GOAL := build-and-run
