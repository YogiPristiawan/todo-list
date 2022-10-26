FROM golang:1.19.2-alpine3.16 AS build

WORKDIR /app

RUN apk add \
    build-base

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build \
    -tags timetzdata \
    -o binary \
    cmd/api/main.go

FROM alpine:3.16.0

RUN apk --no-cache add \
    ca-certificates \
    bash \
    curl

ENV APP_PORT=3030
ENV MYSQL_PORT=3306

WORKDIR /app

COPY --from=build /app ./

EXPOSE 3030

ENTRYPOINT ["./binary"]

