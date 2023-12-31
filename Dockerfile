FROM golang:1.20.6 AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -ldflags="-w -s" -o api cmd/api/main.go

FROM golang:1.20.6 AS final
WORKDIR /
COPY --from=builder /app/api .
EXPOSE 8080
ENTRYPOINT ["/api"]