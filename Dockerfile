FROM golang:1.19-alpine

WORKDIR /app

RUN apk add build-base ffmpeg opus

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

CMD go run main.go