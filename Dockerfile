FROM golang:alpine3.13

RUN apk add curl gcc make musl-dev                                                                                    \
    && go get -u github.com/pressly/goose/cmd/goose                                                                   \
    && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin \
    && addgroup -S appgroup -g 1000                                                                                   \
    && adduser -D -S appuser -G appgroup                                                                              \
    && chown -R appuser /go

USER appuser
WORKDIR /go/src
