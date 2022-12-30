# Build Stage
FROM  golang:1.19-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY wait-for.sh .
EXPOSE 3000
CMD ["/app/main"]