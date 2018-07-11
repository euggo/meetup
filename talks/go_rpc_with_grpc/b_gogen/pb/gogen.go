package pb

//go:generate protoc -I../proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. ../proto/grpcbasic0.proto
//go:generate protoc -I../proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ../proto/grpcbasic0.proto
//go:generate protoc -I../proto -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. ../proto/grpcbasic0.proto
