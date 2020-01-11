BINARY := baelfire
VERSION := $(shell git describe --abbrev=0 --tags)

.PHONY: vendor
vendor:
	go mod vendor

build:
	go build -mod vendor -ldflags="-X main.version=${VERSION} -s -w" -o ${BINARY} .

run: build
	./${BINARY}

release:
	docker build -t impsy/baelfire:${VERSION} .
	docker push impsy/baelfire:${VERSION}
