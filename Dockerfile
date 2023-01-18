FROM golang:1.17-alpine

RUN mkdir /go/src/app
WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@v1.27.3
COPY . .

EXPOSE 8080
CMD "sh" "-c" "go mod tidy && air -c .air.toml"