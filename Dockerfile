FROM golang:1.16 AS build
WORKDIR /go/src/go-server
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get -d -v ./...
RUN go install -v ./...

ENV PORT 8080
EXPOSE $PORT

CMD ["go","run","*.go"]
