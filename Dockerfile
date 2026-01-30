FROM golang:1.25.3 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/main.go

FROM alpine:latest
WORKDIR /app
RUN mkdir db web
COPY --from=builder /app/server .
ENTRYPOINT [ "./server" ]
