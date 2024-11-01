FROM golang:1.23-alpine3.20 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test -v ./server/test/...

RUN go build -o main

FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/main .

CMD [ "/app/main" ]
