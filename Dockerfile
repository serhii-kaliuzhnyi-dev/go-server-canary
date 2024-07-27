# Stage 1: Build the Go application
FROM golang:1.23rc2-alpine3.19 AS builder

# Install necessary packages
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Create a minimal image
FROM scratch

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]
