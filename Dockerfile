FROM golang:1.23.2-alpine
# Working directory
WORKDIR /app
# English: Install air
RUN go install github.com/air-verse/air@latest
# Copy project files
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Expose application port
EXPOSE 8080
# Use air to start the application
CMD ["air", "-c", ".air.toml"]