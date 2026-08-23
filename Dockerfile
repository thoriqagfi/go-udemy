# Build Stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Final Stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY app.env .

EXPOSE 8080

CMD ["/app/main"]