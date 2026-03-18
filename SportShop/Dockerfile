FROM golang:1.25-alpine

WORKDIR /app

COPY . .

RUN go build -o main ./backend/cmd/main.go

EXPOSE 8081

CMD ["./main"]