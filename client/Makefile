.PHONY: build

VERSION=v0.1
GOOS=linux
GOARCH=mipsle
BINARY=reporter
BUILD_ARGS=-trimpath
BUILD_FLAGS=-s -w -X main.version=$(VERSION)

build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) GOMIPS=softfloat \
	 go build -ldflags="$(BUILD_FLAGS)" $(BUILD_ARGS) -o $(BINARY)
