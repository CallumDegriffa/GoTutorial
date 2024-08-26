FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o GoTutorial .

FROM debian:latest

WORKDIR /app

COPY --from=builder /app/GoTutorial .

EXPOSE 8080

CMD ["./GoTutorial"]