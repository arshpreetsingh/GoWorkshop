For simple case
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/*.proto


To use as GO-plugins

protoc --go_out=plugins=grpc:. proto/users/service.proto
