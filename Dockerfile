FROM alpine:latest as builder

RUN apk add --no-cache ca-certificates

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY redis-template-linux .

EXPOSE 80

ENTRYPOINT [ "./redis-template-linux" ]