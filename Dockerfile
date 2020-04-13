FROM golang:1.13 

ARG config

ENV APP_CONFIG=$config

WORKDIR /go

ADD . .

RUN go get ./... \
    && go install github.com/istore221/custo-api

ENTRYPOINT /go/bin/custo-api -config=$APP_CONFIG

EXPOSE 8080