FROM golang:alpine

RUN apk update
RUN apk add git


# create the project folder
RUN mkdir /go/src/go-with-compose
# copy the source
COPY . /go/src/go-with-compose
WORKDIR /go/src/go-with-compose

# install dependencies
RUN go get "github.com/go-redis/redis"
RUN go get "gorm.io/driver/postgres"
RUN go get "gorm.io/gorm"
RUN go get "github.com/satori/go.uuid"
RUN go get "github.com/julienschmidt/httprouter"

ENTRYPOINT go run main.go

EXPOSE 8080