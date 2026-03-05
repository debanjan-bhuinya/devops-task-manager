# -------- Stage 1: Build --------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy full source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server


# -------- Stage 2: Runtime --------
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
