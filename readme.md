protocolbuffer compiler:
https://protobuf.dev/downloads/
necessary plugins:
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
create the pb folder in your project
add this to account.proto - option go_package = "./pb";
finally run this command - protoc --go_out=./pb --go-grpc_out=./pb account.proto
