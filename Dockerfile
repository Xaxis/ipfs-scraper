# Use the official Golang image as a base
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o scraper ./cmd/scraper
RUN ls -la ./cmd/scraper

# Make the binary executable
RUN chmod +x scraper

# List files in /app for debugging
RUN ls -la

# Simple echo command to test CMD functionality
CMD echo "Scraper binary exists in /app and Docker CMD is working"
CMD pwd
CMD ls -la
CMD ["./scraper"]