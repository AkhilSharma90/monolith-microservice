FROM golang:1.13

RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/akhilsharma90/monolith-microservice
COPY . .

RUN dep ensure
RUN go get github.com/cespare/reflex
