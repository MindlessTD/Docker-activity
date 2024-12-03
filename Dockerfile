# Step 1: Use the official Go image as the builder
FROM golang:1.23 AS builder

# Step 2: Set the working directory
WORKDIR /app

# Step 3: Copy go.mod and generate go.sum
COPY go.mod ./

# Step 4: Copy the source code
COPY . .

# Step 5: Build the Go application
RUN go build -o main.go .

# Step 6: Use a minimal image for the runtime
FROM golang:1.23

# Step 7: Set the working directory
WORKDIR /app

# Step 8: Copy the binary from the builder stage
COPY --from=builder /app/main.go .

# Step 9: Expose the port
EXPOSE 8080

# Step 10: Run the application
CMD ["./main.go"]