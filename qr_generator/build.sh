#!/bin/bash

#build qr1 into standalone app
GOOS=linux GOARCH=arm64 go build qr1.go

echo "Build complete"
