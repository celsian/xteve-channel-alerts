# Stage 1: Build the application
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git and build dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Create necessary directories
RUN mkdir -p file/tmp log

# Build the application with explicit architecture-agnostic settings
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o xteve-channel-alerts .

# Stage 2: Create the runtime image
FROM alpine:latest

# Install cron and tzdata for timezone support
RUN apk add --no-cache ca-certificates tzdata dcron

# Create app user for security
RUN addgroup -S app && adduser -S app -G app

# Create necessary directories
RUN mkdir -p /app/file/tmp /app/log

# Copy the binary from the builder stage
COPY --from=builder /app/xteve-channel-alerts /app/

# Set working directory
WORKDIR /app

# Set permissions
RUN chown -R app:app /app

# Create the run script
RUN echo '#!/bin/sh' > /app/run.sh && \
    echo 'cd /app && ./xteve-channel-alerts App' >> /app/run.sh && \
    chmod +x /app/run.sh

# Create entrypoint script to set up cron schedule from env var
# Modified to run crond in background and show application logs in foreground
RUN echo '#!/bin/sh' > /app/entrypoint.sh && \
    echo 'CRON_SCHEDULE=${CRON_SCHEDULE:-"0 4 * * *"}' >> /app/entrypoint.sh && \
    echo 'echo "$CRON_SCHEDULE /app/run.sh >> /app/log/cron.log 2>&1" | crontab -' >> /app/entrypoint.sh && \
    echo 'echo "Starting crond with schedule: $CRON_SCHEDULE"' >> /app/entrypoint.sh && \
    echo 'crond -l 8' >> /app/entrypoint.sh && \
    echo 'touch /app/log/app.log' >> /app/entrypoint.sh && \
    echo 'echo "Displaying application logs (app.log):"' >> /app/entrypoint.sh && \
    echo 'tail -f /app/log/app.log' >> /app/entrypoint.sh && \
    chmod +x /app/entrypoint.sh

# Set environment variables
ENV XTEVE_URL=""
ENV DISCORD_WEBHOOK_URL=""
ENV CRON_SCHEDULE="0 4 * * *"

# Create patch for godotenv to not fail if .env file is missing
RUN echo '# This is a dummy .env file to prevent godotenv from failing' > /app/.env

# Expose volumes for persistence
VOLUME ["/app/file/tmp", "/app/log"]

# Set entrypoint
ENTRYPOINT ["/app/entrypoint.sh"]
