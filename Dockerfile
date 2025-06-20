# Use official Go image
FROM golang:1.21-alpine

# Set working directory to 
WORKDIR /app

# Copy go mod  files first to cache depencies
COPY go.mod ./
RUN go mod tidy


# Copy the entire project
COPY . .

# Build the app
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the executable
CMD ["./main"]