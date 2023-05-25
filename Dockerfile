# Builder
FROM golang:1.19-alpine as builder
RUN apk update && apk upgrade && \
    apk --update add gcc git make curl && \
    source env_var_template

RUN mkdir -p /usr/github.com/isaias-dgr/stori_balance

WORKDIR /usr/github.com/isaias-dgr/stori_balance
COPY . /usr/github.com/isaias-dgr/stori_balance
RUN make 

