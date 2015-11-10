# Pull golang image, with tag 1.5
FROM golang:1.5
# Set workdir
WORKDIR /go/src/github.com/castillobg/example/server

ADD . /go

RUN go install github.com/castillobg/example/server

ENTRYPOINT ["server"]

EXPOSE 8080
