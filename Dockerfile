FROM golang:1.16 AS build
WORKDIR /go/src/go-server
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

ENV PORT 3000
EXPOSE $PORT

CMD ["go","run","main.go"]
