FROM golang:latest

RUN go get github.com/oxequa/realize \
    && go get github.com/go-sql-driver/mysql \
    && go get -u github.com/jinzhu/gorm

EXPOSE 8080

CMD [ "realize", "start", "--run" ]
