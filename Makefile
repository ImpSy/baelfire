BINARY := baelfire
VERSION := $(shell git describe --always --long)

.PHONY: vendor
vendor:
	go mod vendor

build:
	go build -mod vendor -ldflags="-X main.version=${VERSION} -s -w" -o ${BINARY} .

run: build
	./${BINARY}
