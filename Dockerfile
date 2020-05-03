FROM golang:latest

LABEL maintainer="Fszta <antoinefer@hotmail.com>"

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main .
EXPOSE 8001

CMD ["./main"]
