VERSION=`git describe --always`
DATE=`date`
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.date=$(DATE)'
ENV ?= development
PLATFORM := telegram

# Setup
setup:
	go mod download

fmt:
	go fmt $$(go list ./... | grep -v vendor)

lint:
	golint $$(go list ./... | grep -v vendor)

build_telegram:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/telegram cmd/telegram/*.go

build: build_telegram

migrate:
	goose -dir db sqlite3 ./tba.db up

run:
	go run $$(ls -1 cmd/$(PLATFORM)/*.go | grep -v _test.go)
