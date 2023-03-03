VERSION 0.7
FROM golang:1.20.0-alpine3.17

RUN apk add git

WORKDIR app

all:
    BUILD +lint
    # BUILD +gen
    BUILD +test
    BUILD +build

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    # Output these back in case go mod download changes them.
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

# TODO: implement gen check
# gen:
#     FROM +deps
#     COPY . ./
#     RUN go run gen/md.go
#     RUN git diff --quiet --exit-code docs/
#     RUN echo $?
    # ARG GEN_CREATED_DIFF=$(git diff --quiet --exit-code docs/)
    # # RUN 
    # # IF [$? == 1]
    # # RUN git diff --quiet --exit-code go.mod
    # # RUN git diff --quiet --exit-code docs/
    # # RUN git diff --exit-code docs/
    # # RUN $(git diff --quiet --exit-code docs/)
    # IF [ "$GEN_CREATED_DIFF" = 1 ]
    #     RUN git diff docs/
    #     RUN false
    # END

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