FROM golang:1.21-bullseye

RUN go install github.com/cosmtrek/air@latest
RUN groupadd -r app -g 1000 && useradd -u 1000 -r -g app -m -d /home/app -s /usr/bin/bash -c "App user" app
WORKDIR /app/core

# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

COPY . .

EXPOSE 80
# USER app
CMD ["sh", "-c", "go mod download && go mod verify && air -c .air.toml"]
