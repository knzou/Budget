package main

import (
    	"context"
    	"net"

    	"github.com/knzou/BudgetCalculatorService/proto"
    	"google.golang.org/grpc"
        "google.golang.org/grpc/reflection"
)
type server struct{}

func main() {
    listener, err := net.Listen("tcp", ":4040")
    if err != nil {
        panic(err)
    }
    srv := grpc.NewServer()
    // add our services into the grpc server
    proto.RegisterAddServiceServer(srv, &server{})
    // reflection will set up serializing and deserializing
    reflection.Register(srv)

    if e := srv.Serve(listener); e != nil {
        panic(err)
    }
}

func (s *server) GetCategories(ctx context.Context, request *proto.Request) (*proto.GetCategoriesResponse, error) {
	return &proto.GetCategoriesResponse{CatId: "id", Name: "name", TypeId: "typeId"}, nil
}

func (s *server) GetTransactions(ctx context.Context, request *proto.Request) (*proto.GetTransactionsResponse, error) {
	return &proto.GetTransactionsResponse{TranId: "id", CatId: "id", Amount: "amount"}, nil
}