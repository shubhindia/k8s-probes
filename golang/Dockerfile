FROM golang:1.20.8

WORKDIR /app

COPY . /app

RUN go build -o main .

CMD ["/app/main"]
