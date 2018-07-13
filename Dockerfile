FROM golang:1.9.4

ENV GOPATH $GOPATH:/go/src
ENV DB_SOURCE="root:root@tcp(docker.for.mac.localhost:3306)/performance_test?charset=utf8&parseTime=True"

RUN go get github.com/tomo0111/echo-performance
RUN cd /go/src/github.com/tomo0111/echo-performance

CMD dep ensure
CMD go run main.go