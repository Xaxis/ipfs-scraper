# Use the official Golang image as a base
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod file to download dependencies
COPY go.mod tools.go ./

# Download dependencies
RUN go mod download

# Tidy up the module dependencies and generate go.sum
RUN go mod tidy

# Install dependencies, including those specified in tools.go
RUN go install github.com/cosmtrek/air@latest

# Build the application
#RUN go build -o scraper ./cmd/scraper
#RUN go build -o api ./cmd/api

# Copy the source code into the container
COPY . .

# The command to run air live reloader
CMD ["air"]

# Command to run the executable
#CMD ["./api"]
