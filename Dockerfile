# Start from the latest golang base image
FROM golang:1.24-alpine

# Install necessary build tools
RUN apk add --no-cache git curl

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Install Air for hot reloading
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# Install Delve for debugging
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Expose port 8080 to the outside world
EXPOSE 8080 2345

# Command to run the executable
CMD ["air", "-c", ".air.toml"]
