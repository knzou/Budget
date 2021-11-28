### Budget

A service to keep track of people, catagory and transactions data

### Endpoints

Get Requests:
`/getCategories`
`/getTransactions`

`/getPeople` - Fuzzy name search
```bash
curl --location --request GET '104.162.136.185/getPeople' --header 'Content-Type: application/json' --data-raw '{
    "name": "ken"
}'
```
`/getTotalTransactionAmount` - To turn on concurrency set `isParallel` to true
```bash
curl --location --request GET '104.162.136.185/getTotalTransactionAmount' --header 'Content-Type: application/json' --data-raw '{
    "isParallel": false
}'
```

# Quick setup
run bash ./scripts/setup

# Generate Protobuf 

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/knzou/Budget/proto
protoc -I ~/Desktop/Budget/proto ~/Desktop/Budget/proto/service.proto --go_out=plugins=grpc:./.go/src
protoc -I ~/Desktop/Budget/proto ~/Desktop/Budget/proto/date.proto --go_out=plugins=grpc:./.go/src
```

# Best practices
Explicit dependencies
No package level variables
No func init