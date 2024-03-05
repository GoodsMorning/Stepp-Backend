# Use an official Golang runtime as a parent image
FROM golang:1.17 AS builder

# Set the current working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go app
RUN go build -o stepp-backend ./src/main.go

# Start a new stage from scratch
FROM gcr.io/distroless/base-debian10

# Set the working directory inside the container
WORKDIR /

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/app/stepp-backend .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./stepp-backend"]
