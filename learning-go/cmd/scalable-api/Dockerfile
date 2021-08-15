FROM golang:latest

RUN mkdir -p /scalable-api

WORKDIR /scalable-api

COPY . /go/scr/bulletinApi

RUN go install bulletinApi

CMD /scalable-api

EXPOSE 8080