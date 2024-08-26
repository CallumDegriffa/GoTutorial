# Stage 1: Build the Go application
FROM golang:1.23 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o GoTutorial .

# Stage 2: Run the Go application
FROM debian:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/GoTutorial .

# Expose port 8080 (adjust according to your application's needs)
EXPOSE 8080

# Command to run the executable
CMD ["./GoTutorial"]