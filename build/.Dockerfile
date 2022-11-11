FROM golang:1.18 AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /build .
CMD ["./app"]