#!/usr/bin/env bash

. /etc/profile

# arm64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o neverIdle.linux-arm64 main.go
# linux x64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o neverIdle.linux-amd64 main.go
# windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o neverIdle.exe main.go
# mac
#CGO_ENABLED=0 GOOS=linux GOARCH=darwin go build -o neverIdle.mac-amd64 main.go
