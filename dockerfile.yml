FROM golang:1.20-alpine AS build-env

SHELL ["/bin/sh", "-c"]

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get -d ./... ;\
    CGO_ENABLED=0 GOOS=linux go build -o main . ;
    
RUN go build -o main .

CMD ["./main"]