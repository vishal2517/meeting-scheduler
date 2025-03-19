# Use Golang base image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy all files
COPY . .

# Install dependencies
RUN go mod tidy

# Build the application
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["/app/main"]
