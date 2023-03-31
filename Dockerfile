FROM golang:1.17-alpine3.16 AS builder
RUN apk update && apk add --no-cache git && apk add gcc libc-dev

WORKDIR $GOPATH/src/package-service

ENV GOSUMDB=off

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode=external' -o /go/bin/package-service pkg/main.go

FROM alpine:3.12
RUN apk add --no-cache tzdata ca-certificates libc6-compat

COPY --from=builder /go/bin/package-service /go/bin/bountie/task-center-service
COPY --from=builder /go/src/package-service/.env /go/src/package-service/.env

ENTRYPOINT ["/go/bin/bountie/task-center-service"]