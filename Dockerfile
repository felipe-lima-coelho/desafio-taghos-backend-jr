FROM golang:1.23.2-alpine
# Working directory
WORKDIR /app
# Install air
RUN go install github.com/air-verse/air@latest
# Install golang-migrate for postgres
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# Copy project files
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Copy and guarantees permissions script entrypoint
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh
# Expose application port
EXPOSE 8080
# Use script entrypoint to start the application
ENTRYPOINT ["/app/entrypoint.sh"]