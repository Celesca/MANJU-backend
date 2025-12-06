FROM golang:1.24

WORKDIR /app

# Fetch dependencies early (cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Default command for development; runs the app directly
CMD ["sh", "-c", "go run main.go"]
