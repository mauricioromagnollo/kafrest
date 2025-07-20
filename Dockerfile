# ===================================
# Build API to Binary File
# ===================================
FROM golang:1.24.5 AS builder

# Set the root directory
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Get Dependencies
RUN go mod download && go mod tidy

# Copy all files from the current directory to the container
COPY . .

# Build the API binary
RUN CGO_ENABLED=0 GO111MODULES=on go build -ldflags="-w -s" -a -o api cmd/api/main.go

# ===================================
# Run API Binary File
# ===================================
FROM scratch

# Set the root directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/api /app

# Set the environment variable for the port
ENV PORT 8080
EXPOSE $PORT

# Run the API binary
ENTRYPOINT ["/app/api"]