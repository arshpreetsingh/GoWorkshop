FROM golang:1.12.0-alpine3.9

ENV SRV_NAME bounce

# add your package path here
ENV PKG_PATH src/relay/reports/bounce
# mention your "proto" library location here, better to inside your package-path
ADD proto $PKG_PATH/proto
# This your service name aligning with Package path
ADD $SRV_NAME $PKG_PATH/$SRV_NAME
WORKDIR $PKG_PATH/$SRV_NAME
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN go get google.golang.org/grpc github.com/golang/protobuf/proto github.com/sirupsen/logrus
RUN go get -v
RUN go install
EXPOSE 8080
