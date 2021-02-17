VERSION=`git describe --always`
DATE=`date`
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.date=$(DATE)'
CGO_CFLAGS="-I/usr/include/rocksdb"
CGO_LDFLAGS="-L/usr/lib/x86_64-linux-gnu/librocksdb.so.6 -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd"
ENV ?= development
PLATFORM := telegram

# Setup
setup:
	go mod download

build_telegram:
	CGO_FLAGS=$(CGO_CFLAGS) CGO_LDFLAGS=$(CGO_LDFLAGS) GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/telegram cmd/telegram/*.go

build: build_telegram

run:
	go run $$(ls -1 cmd/$(PLATFORM)/*.go | grep -v _test.go)
