FROM golang:1.20.5-alpine3.18

COPY ./pdf-cli ./pdf-cli
RUN go install -v golang.org/x/tools/gopls@latest