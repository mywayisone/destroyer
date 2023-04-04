# Use an official Golang runtime as a parent image
FROM golang:1.16-alpine

# Set the working directory to /app
WORKDIR /destroyer

# Copy the current directory contents into the container at /app
COPY . /destroyer/

# Copy go.mod and go.sum files to the workspace
COPY go.mod /destroyer/
COPY go.sum /destroyer/

# Download and install any required dependencies
RUN apk update && apk add git
RUN go mod download

# Build the Go application
RUN go build -o destroyer

# Expose port 9090 for the application
EXPOSE 9090

# Run the binary when the container starts
CMD ["./destroyer"]