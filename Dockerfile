FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /workspace

# Copy module definition files first
COPY go.mod go.sum ./

# Download dependencies (cached if go.mod/go.sum doesn't change)
RUN go mod tidy

# Copy the rest of the source code
COPY cmd/ ./cmd/
COPY api/ ./api/
COPY pkg/ ./pkg/

RUN go mod tidy

# Build the Go program
RUN go build -o /workspace/bin/server ./cmd/main.go

# Runtime image
FROM alpine:latest
WORKDIR /root
COPY --from=builder /workspace/bin/server ./server

# Expose the required port
EXPOSE 8081

# Run the built server binary
CMD ["./server", "-port=8081"]
