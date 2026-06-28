# Build stage
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod from root (assuming we build from monorepo root context)
COPY go.mod ./
# If go.sum exists, copy it as well. We use a wildcard to avoid failure if it doesn't exist
COPY go.sum* ./
RUN go mod download

# Copy the rest of the monorepo source code
COPY . .

# Build the Go API statically
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/api ./apps/api/cmd/api

# Final stage: Distroless image
FROM gcr.io/distroless/static-debian11:latest

# Copy the binary from the builder stage
COPY --from=builder /go/bin/api /api

# Expose the port the API runs on
EXPOSE 8080

# Run as a non-root user for security
USER nonroot:nonroot

# Entrypoint to run the binary
ENTRYPOINT ["/api"]
