FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o exoplanet-service ./cmd

# Expose the port the service will run on
EXPOSE 8080

# Command to run the executable
CMD ["./exoplanet-service"]
