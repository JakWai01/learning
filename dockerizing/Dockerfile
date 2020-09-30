FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD dockerizing.go /app
WORKDIR /app
RUN go build -o dockerizing .
CMD ["/app/dockerizing"]