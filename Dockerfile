FROM alpine:latest as builder

ARG ARCH=arm64

RUN wget -O /NeverIdle "https://github.com/layou233/NeverIdle/releases/latest/download/NeverIdle-linux-${ARCH}" \
         && chmod +x /NeverIdle

FROM scratch

COPY --from=builder /NeverIdle /NeverIdle
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/NeverIdle"]
CMD ["-c","2h","-m","2","-n","4h"]