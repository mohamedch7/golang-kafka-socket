# Use Go base image
FROM golang:1.20-alpine

WORKDIR /app

# Copy Go modules files
COPY go.mod ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go socket application
RUN go build -o socket_server .

# Start the socket app
CMD ["./socket_server"]
