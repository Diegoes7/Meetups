# Use Go official image
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Install dependencies (e.g., git)
RUN apk add --no-cache git

# Copy go.mod and go.sum from root
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire project folder (all files and subfolders)
COPY . .

# Build the binary from server.go
RUN go build -o server ./server.go

# Expose the port your app listens on (adjust if different)
EXPOSE 8080

# Run the binary
CMD ["./server"]

