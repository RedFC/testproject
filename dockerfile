# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Build the Go application
RUN go build -o testproject .

# Expose the port on which the Go application will run
EXPOSE 8080

# Command to run the executable
CMD ["./testproject"]
