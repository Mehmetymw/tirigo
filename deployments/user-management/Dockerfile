# Start from the official Go image
FROM golang:1.22.3

# Set environment variables for the build
ENV GOOS=linux
ENV GOARCH=amd64

# Set the current working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/user-management

RUN chmod +x main

EXPOSE 8080

CMD ["./main"]
