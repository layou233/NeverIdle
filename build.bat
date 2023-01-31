@echo off

@REM linux arm
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -o neverIdle.linux-arm64 main.go

@REM linux amd64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o neverIdle.linux-amd64 main.go

@REM Win
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o neverIdle.exe main.go

@REM @REM MAC
@REM SET CGO_ENABLED=0
@REM SET GOOS=darwin
@REM SET GOARCH=amd64
@REM go build -o neverIdle.mac main.go