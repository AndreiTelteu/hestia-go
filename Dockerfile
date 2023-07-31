FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go install github.com/cosmtrek/air@latest

EXPOSE 80

CMD ["air", "-c", ".air.toml"]
