VERSION 0.7
FROM golang:1.20.0-alpine3.17
WORKDIR app

all:
    BUILD +lint
    BUILD +test
    BUILD +build

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    # Output these back in case go mod download changes them.
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

lint:
    FROM +deps
    RUN go install github.com/mgechev/revive@latest
    COPY . ./
    RUN revive -formatter friendly

test:
    FROM +deps
    RUN go install gotest.tools/gotestsum@latest
    COPY . ./
    RUN gotestsum --format pkgname

build:
    FROM +deps
    COPY . ./
    RUN go build .