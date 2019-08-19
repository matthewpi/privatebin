CGO    = 0
GOOS   = linux
GOARCH = amd64

SRC_PATH = "go.matthewp.io/privatebin"
OUTPUT_FILE = privatebin

PKG_LIST := $(shell go list ${SRC_PATH}/... | grep -v /vendor/)

all: clean build

test:
	@go test -short ${PKG_LIST}

race:
	@go test -race -short ${PKG_LIST}

mem_san:
	@go test -msan -short ${PKG_LIST}

lint:
	@golint -set_exit_status ${PKG_LIST}

clean:
	@go clean
	@rm $(OUTPUT_FILE) -f

build:
	@CGO_ENABLED=$(CGO) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTPUT_FILE) main.go
