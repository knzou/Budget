### BudgetCalculatorService

# Settings
run bash ./scripts/setup

# Generate Protobuf 
protoc ./proto/service.proto -I. --go_out=plugins=grpc:./.go/src

# Local 
go get -u github.com/knzou/BudgetCalculatorService/proto