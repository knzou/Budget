### Budget

# Settings
run bash ./scripts/setup

# Generate Protobuf 
protoc -I ~/Desktop/Budget/proto ~/Desktop/Budget/proto/service.proto --go_out=plugins=grpc:./.go/src

# Local 
go get -u github.com/knzou/Budget/proto

# Best practices
Explicit dependencies
No package level variables
No func init