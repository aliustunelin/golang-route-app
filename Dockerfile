# Builder aşaması
FROM golang:1.21.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o location

EXPOSE 8080

CMD ["./location"]