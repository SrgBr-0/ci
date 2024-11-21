FROM golang:1.23

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
