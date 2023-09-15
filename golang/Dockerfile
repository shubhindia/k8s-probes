FROM golang:1.20.8-alpine3.18

WORKDIR /app

COPY . /app

RUN go build -o main .

CMD ["/app/main"]
