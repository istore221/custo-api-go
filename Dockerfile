FROM golang:1.13 

ARG config

ENV APP_CONFIG=$config

ENV GOPATH=/go:/app

WORKDIR /app

ADD . .

RUN go get ./... \
    && go get github.com/stretchr/testify \
    && go install github.com/istore221/custo-api

ENTRYPOINT /app/bin/custo-api -config=$APP_CONFIG

EXPOSE 8080