FROM cargo30.dev.caicloud.xyz/library/golang:1.10.1-alpine3.7 as build

WORKDIR /go/src/github.com/tosone/testgo/

COPY . .

RUN go build && ls && cp testgo /tmp/

FROM cargo30.dev.caicloud.xyz/library/alpine:3.7

WORKDIR /app

COPY --from=build /tmp/testgo /app

CMD /app/testgo
