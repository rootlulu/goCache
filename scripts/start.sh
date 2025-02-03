#!/bin/bash

export cryptKey="sjguejglsjejguej"
if [ -z "$1" ]; then
    echo "No argument supplied"
    exit 1
fi

if [[ "$1" = "dev" ]]; then
    go run cmd/server/main.go -c /root/projs/gos/cache/config/config.yaml
fi

if [[ "$1" == "prod" ]]; then
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64
    go build -a -ldflags="-s -w" -o ./build/app ./cmd/server/main.go
    ./build/app -c /root/projs/gos/cache/config/config.yaml
fi
