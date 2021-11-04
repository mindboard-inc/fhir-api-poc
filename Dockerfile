# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /bin/start-api api/main.go


EXPOSE 1500

ENTRYPOINT ["/bin/start-api"]

