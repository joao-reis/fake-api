.PHONY: build ship help run

PROJECT_NAME := fakeapi
GOFLAGS  = -v
LDFLAGS += -s -w
MKDIR_P = mkdir -p

GOPATH = $(shell go env GOPATH)

#### DO NOT CHANGE THOSE FOLDERS.
FOLDER_BIN = bin
MAIN_FOLDER = cmd

ship: clean mod build

help:
	@awk 'BEGIN {FS = ":.*?##"; printf "Usage: make <target>\n"} /^[a-zA-Z_-]+:.*?##/ {printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Create directories
$(FOLDER_BIN):
	$(MKDIR_P) $@

clean:
	rm -rf "$(FOLDER_BIN)"
	@find . -name ".DS_Store" -print0 | xargs -0 rm -f
	go clean -i ./...

mod:
	go mod tidy
	go mod vendor

build: export CGO_ENABLED=0
build: LDFLAGS += -extldflags "-static"
build: $(FOLDER_BIN)
	go build -o $(FOLDER_BIN)/$(PROJECT_NAME)_cli $(GOFLAGS) -ldflags '$(LDFLAGS)' ./$(MAIN_FOLDER)
	go build -o $(FOLDER_BIN)/$(PROJECT_NAME)_api $(GOFLAGS) -ldflags '$(LDFLAGS)' ./$(MAIN_FOLDER)/api

run: ship
	./$(FOLDER_BIN)/$(PROJECT_NAME)_cli
	./$(FOLDER_BIN)/$(PROJECT_NAME)_api
