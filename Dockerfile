FROM golang:latest
RUN go env -w GO111MODULE=auto
WORKDIR /app
COPY . .
CMD go run /app/http_server.go
