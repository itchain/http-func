FROM golang:latest

WORKDIR /go/src/http-func

COPY ./main.go  /go/src/http-func

RUN go build -o httpfunc .

COPY ./hellofunc .

ENTRYPOINT [ "http-func" ]