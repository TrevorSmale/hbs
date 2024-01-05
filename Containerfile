# Base image
FROM alpine:latest

# Install necessary packages
RUN apk update && \
    apk add go nodejs npm vim bash

# Expose port 8080
EXPOSE 8080

# Set working directory
WORKDIR /app

# Copy project files to the container
COPY . /app

# Additional configurations or commands (e.g., installing Terminal) can be added here

# Command to run when the container starts
CMD ["sh"]
