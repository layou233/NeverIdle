@echo off

@REM Linux arm64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -trimpath -ldflags="-s -w -buildid=" -o NeverIdle-linux-arm64 main.go

@REM Linux amd64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -trimpath -ldflags="-s -w -buildid=" -o NeverIdle-linux-amd64 main.go

@REM Windows amd64
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -trimpath -ldflags="-s -w -buildid=" -o NeverIdle-windows-amd64.exe main.go

@REM @REM macOS amd64
@REM SET CGO_ENABLED=0
@REM SET GOOS=darwin
@REM SET GOARCH=amd64
@REM go build -trimpath -ldflags="-s -w -buildid=" -o NeverIdle-darwin-amd64 main.go