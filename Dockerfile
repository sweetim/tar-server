FROM golang:1.13

LABEL maintainer="hosweetim@gmail.com"

WORKDIR /go/src/github.com/sweetim/tar-server
COPY . .

RUN go get
RUN go install

ENTRYPOINT /go/bin/tar-server

EXPOSE 3000
