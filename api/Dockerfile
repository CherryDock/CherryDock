FROM golang:latest

LABEL maintainer="Fszta <antoinefer@hotmail.com>"

WORKDIR /app/docker-api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/docker-api .
EXPOSE 8001

CMD ["./out/docker-api"]