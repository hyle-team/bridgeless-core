FROM golang:1.20-alpine as buildbase

RUN apk add build-base git



WORKDIR /go/src/github.com/hyle-team/bridgeless-core

ENV GO111MODULE="on"
ENV CGO_ENABLED=1
ENV GOOS="linux"
ENV GOPRIVATE=github.com/*
ENV GONOSUMDB=github.com/*
ENV GONOPROXY=github.com/*

COPY ./go.mod ./go.sum ./
# Read the CI_ACCESS_TOKEN from the .env file
RUN --mount=type=secret,id=_env,dst=./.env

RUN git config --global url."https://${CI_ACCESS_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN go mod download

COPY . .

RUN go mod vendor

RUN go build -o /usr/local/bin/bridgeless-core github.com/hyle-team/bridgeless-core/cmd/bridgeless-cored



###

FROM alpine:3.9

RUN apk add --no-cache ca-certificates

COPY --from=buildbase /usr/local/bin/bridgeless-core /usr/local/bin/bridgeless-core

ENTRYPOINT ["bridgeless-core"]