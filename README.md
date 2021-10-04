### BudgetCalculatorService

# Generate Protobuf 
protoc -I ./ ./proto/service.proto --go_out=plugins=grpc:./proto

# Local 
go get -u github.com/knzou/BudgetCalculatorService/proto