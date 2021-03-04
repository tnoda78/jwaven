NAME := jwaven
RELEASE_DIR := release

setup:
	go get -u github.com/golang/dep/cmd/dep
	go get golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help

deps: setup
	dep ensure

lint: deps
	go list ./... | xargs -L1 golint -set_exit_status

test: deps
	go test ./...

build: deps
	go build -o ${RELEASE_DIR}/${GOOS}_${GOARCH}/${NAME}${SUFFIX} cmd/jwaven/main.go

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

