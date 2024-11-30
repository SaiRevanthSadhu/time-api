# Build Stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./src/main.go

# Runtime Stage
FROM debian:bullseye
WORKDIR /app
COPY --from=builder /app/main .
COPY config/ config/
EXPOSE 8080
CMD ["./main"]
