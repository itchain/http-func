FROM golang:latest

WORKDIR /go/src/http-func

COPY ./main.go  /go/src/http-func

RUN go build -o httpfunc .

FROM golang:latest

COPY --from=0 /go/src/http-func/httpfunc /usr/local/bin

CMD [ "httpfunc" ]