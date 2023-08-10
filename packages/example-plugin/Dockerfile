FROM golang:1.21-bullseye
RUN apt-get update && apt-get install -y build-essential

RUN groupadd -r app -g 1000 && useradd -u 1000 -r -g app -m -d /home/app -s /usr/bin/bash -c "App user" app
WORKDIR /app/packages/example-plugin

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go install github.com/cosmtrek/air@latest && chown -R app: /app

EXPOSE 80
USER app
CMD ["air", "-c", ".air.toml"]
