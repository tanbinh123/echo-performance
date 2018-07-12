FROM golang:1.9.4

ENV GOPATH $GOPATH:/go/src
ENV DB_SOURCE="root:root@docker.for.mac.localhost:3306/performance_test?charset=utf8&parseTime=True"

RUN go get github.com/labstack/echo && \
    go get github.com/jinzhu/gorm && \
    go get github.com/jinzhu/gorm/dialects/mysql

RUN mkdir /go/tomo0111/echo-performance

CMD go run main.go