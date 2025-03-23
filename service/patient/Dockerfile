FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire codebase
COPY . .

# Build the application - now correctly targeting the cmd/patient directory
RUN CGO_ENABLED=0 GOOS=linux go build -o /patient-service ./cmd/patient

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /patient-service .
COPY .env ./

# Run the service
CMD ["./patient-service"]
