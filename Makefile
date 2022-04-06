.PHONY: build release wasm-common fmt test lint clean help run
# Inspired by: https://github.com/azer/go-makefile-example/blob/master/Makefile

BUILD := $(shell git rev-parse --short HEAD 2>/dev/null)
PROJECTNAME := $(shell basename "$(PWD)")
GOROOT := $(shell go env GOROOT)

# -X add string value definition of the form importpath.name=value
RELEASE := -ldflags "-s -w -X project.name=anotherwebasembly"
DISTDIR := ./dist

## serve: runs a devserver so the UI is accessible at http://localhost:8080/
serve: wasm-common
	$(DISTDIR)/devserver -dir $(DISTDIR)

## build: build all the things
build: wasm-common
	@echo "  >  BUILD WASM app"
	@env GOARCH=wasm GOOS=js go build -o "$(DISTDIR)/lib.wasm" ./wasm
	@echo "  DONE! run devserver in the dist directory"

wasm-common:
	@echo "  >  BUILD WASM common"
	@mkdir -p $(DISTDIR)
	@cp -rf web/ $(DISTDIR)
	@go build -o "$(DISTDIR)/devserver" $(RELEASE) cmd/devserver/main.go
	@cp -f "$(GOROOT)/misc/wasm/wasm_exec.js" $(DISTDIR)

## format: format code using go fmt
fmt:
	@gofmt -w .

## test: run unit tests
test:
	@env PATH=$(PATH):$(GOROOT)/misc/wasm GOARCH=wasm GOOS=js go test -cover -v ./wasm

## lint: static analyze source
lint:
	@env GOARCH=wasm GOOS=js golangci-lint run ./...

## clean: removes build files
clean:
	@go clean
	@rm -rf $(DISTDIR)/*

all: help
help: Makefile
	@echo "Make commands available in "$(PROJECTNAME)":"
	@# Prints all the ## comments from this file with nice formating
	@# make targets without a ## comment wont be printed
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ make/'
