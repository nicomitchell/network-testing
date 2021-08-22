FROM golang:1.16-alpine

WORKDIR /app

COPY src ./

RUN go mod download

RUN go build main.go 

EXPOSE 8080
EXPOSE 14111-14116

CMD ["./main", "ipl","8080","14111:14116"]