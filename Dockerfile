FROM golang:1.23-alpine AS builder

# Set the working directory
WORKDIR /workspace

# Copy module definition files
COPY go.mod go.sum ./

# Download and Cache deps
RUN go mod tidy

# Copy the go source
COPY cmd/ ./cmd/
COPY api/ ./api/
COPY pkg/ ./pkg/

# Build Go Program
RUN go build -o /workspace/bin/server ./cmd/main.go

FROM alpine:latest
WORKDIR /root
COPY --from=builder /workspace/gRPCtestingServer .

EXPOSE 8081

CMD ["./server", "-port=8081"]
