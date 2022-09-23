FROM golang:1.18

RUN apt-get update &&  apt-get add git

RUN go get github.com/cosmtrek/air@v1.29.0

WORKDIR /app

CMD ["air", "-c", ".air.toml"]