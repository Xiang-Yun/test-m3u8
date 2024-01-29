# base go image
FROM golang:alpine3.19 as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o test-m3u8 ./cmd/web

RUN chmod +x /app/test-m3u8


# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/test-m3u8 /app

CMD ["/app/test-m3u8"]
