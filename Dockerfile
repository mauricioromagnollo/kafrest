FROM golang:1.20.6 as builder
WORKDIR /app
COPY . .
RUN GOOS=linux go build -ldflags="-w -s" -o bin/api cmd/api/main.go

FROM scratch
COPY --from=builder /app/bin/api .
EXPOSE 8080
CMD [ "./api" ]