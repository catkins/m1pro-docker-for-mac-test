FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN go get -d -v .
RUN go build -v -o bench main.go 

CMD ["./bench"]