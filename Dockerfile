FROM golang:1.10-alpine

RUN apk -U add make git
RUN go get -u github.com/jaimelopez/chihuahua
RUN cp /go/bin/chihuahua /usr/local/bin/chihuahua

