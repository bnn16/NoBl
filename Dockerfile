FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod file
COPY go.mod .

# Copy source code
COPY . .

# Build the application
RUN go build -o NoBl

# Create a minimal image
FROM alpine:latest

WORKDIR /app

# Copy the binary
COPY --from=builder /app/NoBl .

# Copy required files
COPY examples.nobl .
COPY test.nobl .

# Expose port (will be overridden by PORT env variable if set)
EXPOSE 8080

# Run the application
CMD ["./NoBl", "--web"]
