FROM golang:latest as builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o api ./cmd/main.go

FROM ubuntu AS production

WORKDIR /app

RUN apt-get update
RUN apt-get -y install postgresql-client

COPY --from=builder /app/api ./
COPY --from=builder /app/wait-for-postgres.sh ./
COPY --from=builder /app/config.env ./

RUN