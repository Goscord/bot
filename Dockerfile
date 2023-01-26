FROM golang:1.18-alpine

WORKDIR /app

RUN sudo apk add ffmpeg

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

CMD go run main.go