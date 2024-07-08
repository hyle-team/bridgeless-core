FROM golang:1.21.2-bullseye AS build-env

WORKDIR /go/src/github.com/evmos/evmos

RUN apt-get update -y
RUN apt-get install git -y

# Copy source code
COPY . .

# Ensure dependencies are properly handled
RUN sed -i '/toolchain/d' go.mod

# Download Go modules and verify
RUN go mod tidy && go mod verify

# Build the project
RUN make build


FROM golang:1.21.2-bullseye

RUN apt-get update -y
RUN apt-get install ca-certificates jq -y

WORKDIR /root

COPY --from=build-env /go/src/github.com/evmos/evmos/build/evmosd /usr/bin/evmosd

EXPOSE 26656 26657 1317 9090 8545 8546

CMD ["evmosd"]
