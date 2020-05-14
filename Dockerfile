FROM golang:1.13 AS builder

LABEL maintainer="Fszta <antoinefer@hotmail.com>"

WORKDIR /cherrydock-api

# Install go dependencies
COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go mod download

COPY ./api .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/cherrydock

FROM node:10
WORKDIR /usr/src/app

# Get go api binary exec
COPY --from=builder /cherrydock-api/cherrydock .
COPY --from=builder /cherrydock-api/config.yml .

COPY package*.json ./
RUN npm install
COPY . .
RUN npm run-script build
RUN npm install -g serve

# UI port
EXPOSE 5000

# Api port, only for dev to access api from outside
EXPOSE 8001

CMD ["/bin/bash","start.sh"]