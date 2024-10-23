FROM golang:1.16

WORKDIR /app
COPY . .

# Install Go modules
RUN go mod download

# Build the Go application
RUN go build -o consumer

CMD ["./consumer"]
