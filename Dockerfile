## Build
FROM golang:1.21.10-alpine3.18 AS buildenv

LABEL version="1.0"
LABEL author="mrbelka12000"

RUN apk add --no-cache build-base

ADD go.mod go.sum /

RUN go mod download

WORKDIR /app

ADD . .

ENV GO111MODULE=on
ENV CGO_ENABLED=1

RUN  go build -o main cmd/main.go

## Deploy
FROM alpine

WORKDIR /

COPY --from=buildenv  /app/ /

CMD ["/main"]