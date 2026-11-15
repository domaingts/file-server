FROM golang:1.27.4-alpine3.24 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o file-server ./


FROM alpine:3.24 AS dist

WORKDIR /app

COPY --from=builder /app/file-server .

ENTRYPOINT [ "./file-server" ]
