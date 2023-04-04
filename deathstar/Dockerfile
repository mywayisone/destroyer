# Use an official Golang runtime as a parent image
FROM golang:1.16-alpine

# Set the working directory to /app
WORKDIR /deathstar

# Copy the current directory contents into the container at /app
COPY . /deathstar/

# Download and install any required dependencies
RUN apk update && apk add git
RUN go mod download

# Build the Go application
RUN go build -o deathstar

# Expose port 8080 for the application
EXPOSE 8080

# Run the binary when the container starts
CMD ["./deathstar"]
