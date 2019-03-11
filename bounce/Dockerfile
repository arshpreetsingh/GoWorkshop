FROM golang:1.9-alpine

ENV SRV_NAME bounce

ENV PKG_PATH /go/src/github.com/wizelineacademy/GoWorkshop

ADD proto $PKG_PATH/proto
ADD $SRV_NAME $PKG_PATH/$SRV_NAME
WORKDIR $PKG_PATH/$SRV_NAME
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN go get google.golang.org/grpc github.com/golang/protobuf/proto github.com/sirupsen/logrus
RUN go get -v
RUN go install
EXPOSE 8080
