# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the main application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o seat-arrangements .

# Build the migration tool
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migrate ./cmd/migrate

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk --no-cache add ca-certificates tzdata

# Copy the binaries from the builder stage
COPY --from=builder /app/seat-arrangements /app/
COPY --from=builder /app/migrate /app/

# Copy the config directory
COPY --from=builder /app/config /app/config

# Copy the migrations directory
COPY --from=builder /app/db/migrations /app/migrations

# Copy the migration script and make it executable
COPY run-migrations.sh /app/
RUN chmod +x /app/run-migrations.sh

# Expose the port
EXPOSE 8080

# Set the entrypoint
ENTRYPOINT ["/app/seat-arrangements"]

# Set the default command
CMD ["serve"]