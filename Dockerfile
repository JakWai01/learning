FROM golang:1.12.0-alpine3.9
RUN apk add git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main ./cmd/kubernetes-tcp-server/main.go
CMD ["/app/main"]