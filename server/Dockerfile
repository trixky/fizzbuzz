FROM golang:alpine

RUN apk update
RUN apk add git

RUN mkdir /go/src/go-with-compose
COPY . /go/src/go-with-compose
WORKDIR /go/src/go-with-compose

RUN go get "github.com/go-redis/redis"
RUN go get "gorm.io/driver/postgres"
RUN go get "gorm.io/gorm"
RUN go get "github.com/satori/go.uuid"

ENTRYPOINT go run /go/src/go-with-compose/*.go

EXPOSE 8080