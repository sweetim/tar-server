FROM golang:1.9

LABEL maintainer="hosweetim@gmail.com"

COPY . /go/src/github.com/sweetim/tar-server
WORKDIR /go/src/github.com/sweetim/tar-server

RUN go get
RUN go install

ENTRYPOINT /go/bin/tar-server
EXPOSE 3000
