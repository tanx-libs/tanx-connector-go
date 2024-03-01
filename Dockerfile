# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Golang library
RUN go build -o mylibrary

# Set the entry point for the container
ENTRYPOINT ["./mylibrary"]