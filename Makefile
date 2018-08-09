NAME := jwaven
VERSION := 0.4.0
RELEASE_DIR := release

setup:
	go get golang.org/x/vgo
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help

deps: setup

lint: deps
	golint -set_exit_status ./...

test: deps
	vgo test ./...

build: deps
	vgo build -o ${RELEASE_DIR}/${GOOS}_${GOARCH}/${NAME}${SUFFIX} cmd/jwaven/main.go

build-darwin-amd64:
	@$(MAKE) build GOOS=darwin GOARCH=amd64

build-linux-amd64:
	@$(MAKE) build GOOS=linux GOARCH=amd64

build-linux-386:
	@$(MAKE) build GOOS=linux GOARCH=386

build-windows-amd64:
	@$(MAKE) build GOOS=windows GOARCH=amd64 SUFFIX=.exe

build-windows-386:
	@$(MAKE) build GOOS=windows GOARCH=386 SUFFIX=.exe

build-all: build-darwin-amd64 build-linux-amd64 build-linux-386 build-windows-amd64 build-windows-386

