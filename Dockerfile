FROM golang:1.13 

WORKDIR /go

ADD . .

RUN go install github.com/istore221/custo-api

WORKDIR /go/bin

ENTRYPOINT /go/bin/custo-api

EXPOSE 8080