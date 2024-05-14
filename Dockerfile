FROM golang AS builder
COPY . /NeverIdle
WORKDIR /NeverIdle
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o NeverIdle \
    && chmod +x NeverIdle

FROM scratch
COPY --from=builder /NeverIdle/NeverIdle /app/NeverIdle
CMD ["/app/NeverIdle"]
