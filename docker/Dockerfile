# Start from the official Go image with Alpine for a minimal base
FROM golang:1.22-alpine AS build

# Set necessary environment variables required for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Install git and ca-certificates
RUN apk update && apk add --no-cache ca-certificates git

# Use Go's module proxy and set a direct backup if unavailable
ENV GOPROXY=https://proxy.golang.org,direct

# Copy go.mod and go.sum files to leverage Docker layer caching
COPY go.mod go.sum ./

# Download dependencies early to improve layer caching
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go app as a static binary to avoid dependencies on libc
RUN go build -ldflags="-s -w" -o /usr/local/bin/app cmd/main.go

# Create a non-root user and group in the build stage
RUN addgroup -S appgroup && adduser -S -G appgroup appuser

# Change ownership of the binary to the non-root user
RUN chown appuser:appgroup /usr/local/bin/app

# Change ownership of the application files to the non-root user
RUN chown -R appuser:appgroup /app

# Final minimal image
FROM scratch

# Copy ca-certificates from build stage
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the pre-built binary from the build stage
COPY --from=build /usr/local/bin/app /usr/local/bin/app

# Copy the necessary configuration and resources
COPY .env /app/.env
COPY migrations /app/migrations
COPY seeds /app/seeds

# Copy the non-root user setup from the build stage
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group

# Set the working directory to /app
WORKDIR /app

# Switch to the non-root user
USER appuser

# Expose the port your application will run on
EXPOSE 3060

# Command to run the executable
CMD ["app"]
