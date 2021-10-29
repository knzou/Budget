package main

import (
	"context"
	"database/sql"
	"net"

	"github.com/knzou/BudgetCalculatorService/budget"
	proto "github.com/knzou/BudgetCalculatorService/proto"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
    host    = "localhost"
    port    = 5432
    user    = "kenzou"
    dbname  = "kenzou"
)
type server struct{
	db *sqlx.DB
}

func main() {

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "dbname=%s sslmode=disable",
    host, port, user, dbname)
    // open and connect at the same tim, panicing on error
    db := sqlx.MustConnect("postgres", psqlInfo)

	defer db.Close()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	// add our services into the grpc server with our created connection pool db
	proto.RegisterAddServiceServer(srv, &server{db: db})
	// reflection will set up serializing and deserializing
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) GetCategories(ctx context.Context, request *proto.Request) (*proto.GetCategoriesResponse, error) {
	var categories Category[]
	categories, err = GetCategories(s.db)
	if err != nil {
		panic(err)
	}
	return &proto.GetCategoriesResponse{categories: categories}, nil
}

func (s *server) GetTransactions(ctx context.Context, request *proto.Request) (*proto.GetTransactionsResponse, error) {
	return &proto.GetTransactionsResponse{transactions: []}, nil
}
