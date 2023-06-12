FROM golang:1.20.5-alpine3.18
RUN apk add --no-cache bash

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go install -v golang.org/x/tools/gopls@latest
RUN go install github.com/spf13/cobra-cli@latest

RUN mkdir dev
COPY . ./dev
SHELL ["/bin/bash", "-c"]