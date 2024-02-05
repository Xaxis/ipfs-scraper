# Use the official Golang image as a base
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download
RUN go get -d -v github.com/lib/pq@v1.10.9
RUN go get -d -v github.com/gorilla/mux@v1.7.4

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o scraper ./cmd/scraper
RUN go build -o server ./cmd/server

# Make the binary executable
RUN chmod +x scraper
RUN chmod +x server

# List files in /app for debugging
RUN ls -la

# Copy start script and make it executable
COPY ./scripts/start.sh ./start.sh
RUN chmod +x start.sh

# Run the start script as the default command
CMD ["./start.sh"]