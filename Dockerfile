# Use the official Go image as a base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /build

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o /nextgen .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["/build/nextgen"]
