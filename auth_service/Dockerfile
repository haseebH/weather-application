
FROM golang:1.22-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download

# Copy the source code into the container
COPY . .

RUN go build -o rbac .

FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/rbac .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./rbac"]
