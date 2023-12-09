# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.20-alpine as build

WORKDIR /go/src/github.com/tosone/testgo/

COPY . .

ARG TARGETOS TARGETARCH

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -v && ls && cp testgo /tmp/

FROM alpine:3.8

WORKDIR /app

COPY --from=build /tmp/testgo /app

CMD /app/testgo

# docker buildx build --platform linux/amd64,linux/arm64 -t tosone/testgo:latest --push .
