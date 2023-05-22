# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code from the host to the container
COPY . .

# Build the Go application
RUN go build -o app

EXPOSE 8080

# Set the command to run the executable when the container starts
CMD ["./app"]
