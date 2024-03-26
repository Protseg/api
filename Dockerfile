# Use the official Golang image as the base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go binary
RUN go build -o api-linux

# Expose port 3003 for the container
EXPOSE 3003

# Set the entrypoint to run the binary
ENTRYPOINT ["./api-linux"]