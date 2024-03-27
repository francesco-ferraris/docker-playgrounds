IMAGE_NAME=docker-playground
TAG=multistage
TAG_TEST=test

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
	@docker image rm $(IMAGE_NAME):$(TAG_TEST)

.PHONY: test
test:
	@docker build --target run-test-stage --progress plain --no-cache -t $(IMAGE_NAME):$(TAG_TEST) .

.DEFAULT_GOAL := build-and-run
