FROM golang:1.9.2-stretch
MAINTAINER Kohei SUZUKI <skjune12@mercari.com>

WORKDIR /app
ADD main.go /app
# ADD static /app
RUN go build -o app

EXPOSE 3000
CMD ["./app"]
