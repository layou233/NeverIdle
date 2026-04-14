FROM golang:alpine AS builder

ARG ARCH=arm64
WORKDIR /app
ENV GOTOOLCHAIN=auto
COPY . .
RUN go build -o NeverIdle main.go

FROM alpine:latest

# 安装 iproute2 (提供 ss 命令) 和 tzdata (提供时区支持)
RUN apk add --no-cache iproute2 tzdata

COPY --from=builder /app/NeverIdle /usr/local/bin/NeverIdle

ENTRYPOINT ["NeverIdle"]
CMD ["-c", "2h", "-m", "2", "-n", "4h"]
