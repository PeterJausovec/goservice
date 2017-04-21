FROM golang:1.7
# ENV GOPATH /go
COPY . /go/src/app
WORKDIR /go/src/app
RUN go-wrapper download

