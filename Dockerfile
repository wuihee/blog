# Use the official Go image as the base.
FROM golang:1.23-alpine AS builder

# Set the working directory in the container.
WORKDIR /app

# Copy the Go module files and download dependencies.
COPY go.mod ./
RUN go mod download

# Copy the entire project into the container.
COPY . .

# Build the Go application.
RUN go build -o blog .

# Use a minimal base image for the final container.

FROM alpine:latest

# Install necessary libraries
RUN apk --no-cache add ca-certificates

# Set the working directory in the container.
WORKDIR /root/

# Copy the built Go application from the builder.
COPY --from=builder /app/blog .

# Copy static files and templates.
COPY static ./static
COPY templates ./templates

# Expose the port the app runs on.
EXPOSE 8080

# Run the application.
CMD ["./blog"]
