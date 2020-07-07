FROM golang:latest

WORKDIR /

COPY . .

CMD ["go", "run", "server.go"]