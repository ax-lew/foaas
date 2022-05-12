GORUN ?= go run
GOFMT ?= gofmt
GOGET ?= go get
GOBUILD ?= go build
IMAGE_TAG ?= $(shell git rev-parse --short=7 HEAD)

deps:
	$(GOGET) ./...

run:
	$(GORUN) ./main.go serve

fmt:
	$(GOFMT) -w .


docker-build:
	docker build --build-arg VERSION=${IMAGE_TAG} -t fooas .

docker-run:
	docker run -it --rm -p 8080:8080 fooas

