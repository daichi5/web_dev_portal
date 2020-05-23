FROM golang:latest

RUN go get -v github.com/oxequa/realize

EXPOSE 8080

CMD [ "realize", "start", "--run" ]
