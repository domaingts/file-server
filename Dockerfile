FROM golang:1.25.7-alpine3.22 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o file-server ./


FROM alpine:3.22 AS dist

WORKDIR /app

COPY --from=builder /app/file-server .

ENTRYPOINT [ "./file-server" ]
