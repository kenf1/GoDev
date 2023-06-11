FROM golang:1.20.5-alpine3.18

RUN apk add --no-cache bash
RUN go install -v golang.org/x/tools/gopls@latest
SHELL ["/bin/bash", "-c"]