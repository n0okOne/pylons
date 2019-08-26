# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Michael Sofaer <m@pylons.tech>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build ./cmd/pylonsd/

RUN ./pylonsd init --chain-id pylonschain

RUN go build ./cmd/pylonscli/

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
RUN /app/pylonsd start &
CMD ["/app/pylonscli", "rest-server", "--chain-id", "pylonschain", "--trust-node", "--cors", "*", "--laddr", "tcp://0.0.0.0:8080"]
