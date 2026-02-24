# ===================================
# Build API to Binary File
# ===================================
FROM public.ecr.aws/docker/library/golang:1.26.0 AS builder

#avoid root
ENV USER=appuser
ENV UID=1000

#avoid root
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# Set the root directory
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod* go.sum* ./

# Get Dependencies
RUN go mod download && go mod tidy

# Copy all files from the current directory to the container
COPY . .

# Build the API binary
RUN CGO_ENABLED=0 GO111MODULES=on go build -ldflags="-w -s" -a -o kafrest-api cmd/api

# ===================================
# Run API Binary File
# ===================================
FROM scratch

# Set the root directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/kafrest-api /app

# Set the environment variable for the port
ENV PORT 8080
EXPOSE $PORT

#avoid rootless
USER appuser:appuser

# Run the API binary
ENTRYPOINT ["/app/kafrest-api"]
