package main

import (
	"context"
	_ "database/sql"
	"net"
	"fmt"

	"github.com/knzou/BudgetCalculatorService/db"
	proto "github.com/knzou/BudgetCalculatorService/proto"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
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
    db1 := sqlx.MustConnect("postgres", psqlInfo)
	// defer db1.Close()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	// add our services into the grpc server with our created connection pool db
	proto.RegisterAddServiceServer(srv, &server{db: db1})
	// reflection will set up serializing and deserializing
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) GetCategories(ctx context.Context, request *proto.Request) (*proto.GetCategoriesResponse, error) {
	categories, err := db.GetCategories(s.db)
	if err != nil {
		panic(err)
	}
	// defer s.db.Close()// not needed, once app close, db connection close too

	var cats []*proto.GetCategoriesResponse_Category
	for _, category := range categories {
		cats = append(cats, &proto.GetCategoriesResponse_Category{CatId: category.CatId , Name: category.Name, TypeId: category.TypeId})
	}
	return &proto.GetCategoriesResponse{Categories: cats}, nil
}

func (s *server) GetTransactions(ctx context.Context, request *proto.Request) (*proto.GetTransactionsResponse, error) {
	return &proto.GetTransactionsResponse{}, nil
}
