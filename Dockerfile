FROM golang:1.20-alpine as buildbase

RUN apk add build-base git

ARG CI_JOB_TOKEN

WORKDIR /go/src/github.com/hyle-team/bridgeless-core

ENV GO111MODULE="on"
ENV CGO_ENABLED=1
ENV GOOS="linux"
ENV GOPRIVATE=github.com/*
ENV GONOSUMDB=github.com/*
ENV GONOPROXY=github.com/*

RUN echo "machine github.com login gitlab-ci-token password $CI_JOB_TOKEN" > ~/.netrc
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go mod vendor

RUN go build -o /usr/local/bin/bridgeless-core github.com/hyle-team/bridgeless-core/cmd/bridgeless-cored



###

FROM alpine:3.9

RUN apk add --no-cache ca-certificates

COPY --from=buildbase /usr/local/bin/bridgeless-core /usr/local/bin/bridgeless-core

ENTRYPOINT ["bridgeless-core"]