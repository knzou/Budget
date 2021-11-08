package main

import (
	"context"
	_ "database/sql"
	"net"
	"fmt"
	"strconv"

	"github.com/knzou/Budget/db"
	proto "github.com/knzou/Budget/proto"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)
// Use flags to pass in db connection credentials by using -h flag in command line
const (
    host    = "localhost"
    port    = 5432
    user    = "test_user"
    dbname  = "kenzou"
)
type server struct{
	rdb *sqlx.DB
}

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "dbname=%s sslmode=disable",
    host, port, user, dbname)

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	// add our services into the grpc server with our db instance, which will close once app exit
	proto.RegisterAddServiceServer(srv, &server{rdb: sqlx.MustConnect("postgres", psqlInfo)})
	// reflection will set up serializing and deserializing
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) GetCategories(ctx context.Context, request *proto.Request) (*proto.GetCategoriesResponse, error) {
	categories, err := db.GetCategories(s.rdb)
	if err != nil {
		panic(err)
	}

	var cats []*proto.GetCategoriesResponse_Category
	for _, category := range categories {
		cats = append(cats, &proto.GetCategoriesResponse_Category{CatId: category.CatId , Name: category.Name, TypeId: category.TypeId})
	}
	return &proto.GetCategoriesResponse{Categories: cats}, nil
}

func (s *server) GetTransactions(ctx context.Context, request *proto.Request) (*proto.GetTransactionsResponse, error) {
	transactions, err := db.GetTransactions(s.rdb)
	if err != nil {
		panic(err)
	}
	var trans []*proto.GetTransactionsResponse_Transaction
	for _, transaction := range transactions {
		year, err := strconv.Atoi(transaction.TransDate[0:4])
		if err != nil {
			panic(err)
		}
		month, err := strconv.Atoi(transaction.TransDate[5:7])
		if err != nil {
			panic(err)
		}
		day, err := strconv.Atoi(transaction.TransDate[8:10])
		if err != nil {
			panic(err)
		}
		transformTransDate := &proto.Date{Year: int32(year), Month: int32(month), Day: int32(day)}
		trans = append(trans, &proto.GetTransactionsResponse_Transaction{TranId: transaction.TranId , CatId: transaction.CatId, TransDate: transformTransDate, Amount: transaction.Amount})
	}
	return &proto.GetTransactionsResponse{Transactions: trans}, nil
}

func (s *server) GetPeople(ctx context.Context, request *proto.GetPeopleRequest) (*proto.GetPeopleResponse, error) {
	people, err := db.GetPeople(s.rdb, request)
	if err != nil {
		panic(err)
	}
	var ppl []*proto.GetPeopleResponse_Person
	for _, person := range people {
		ppl = append(ppl, &proto.GetPeopleResponse_Person{Pid: person.Pid , Name: person.Name})
	}
	return &proto.GetPeopleResponse{People: ppl}, nil
}