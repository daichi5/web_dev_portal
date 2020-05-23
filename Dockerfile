FROM golang:latest

WORKDIR /app

ENTRYPOINT [ "go", "run", "./cmd/web_dev_portal/main.go" ]