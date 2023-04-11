FROM alpine:3.17.3 as builder

ENV ARCH arm64

run wget -O /NevreIdle "https://github.com/layou233/NeverIdle/releases/latest/download/NeverIdle-linux-$ARCH" \
         && chmod +x /NevreIdle

FROM scratch

COPY --from=builder /NevreIdle /NevreIdle
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/NevreIdle"]
CMD ["-c","2h","-m","2","-n","4h"]