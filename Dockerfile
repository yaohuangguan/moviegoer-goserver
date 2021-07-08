FROM golang:1.16 AS build
WORKDIR /go/src/go-server
# COPY . .
# RUN go env -w GOPROXY=https://goproxy.cn,direct
# RUN go get -d -v ./...
# RUN go install -v ./...

# ENV PORT 8080
# EXPOSE $PORT

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
