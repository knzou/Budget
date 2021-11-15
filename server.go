package main

import (
	"context"
	"log"
	_ "database/sql"
	"net"
	"fmt"
	"time"
	"strconv"
	"sync"

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
	pool := sqlx.MustConnect("postgres", psqlInfo)
	pool.SetConnMaxIdleTime(10 * time.Second)
	// add our services into the grpc server with our db instance, which will close once app exit
	proto.RegisterAddServiceServer(srv, &server{rdb: pool})
	// reflection will set up serializing and deserializing
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) GetCategories(ctx context.Context, request *proto.Request) (*proto.GetCategoriesResponse, error) {
	stats := s.rdb.Stats()
	log.Printf("Pool Status \n Open Connections: %d \n InUse: %d \n Idle: %d", stats.OpenConnections, stats.InUse, stats.Idle)
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

func (s *server) GetTotalTransactionAmount(ctx context.Context, request *proto.GetTotalTransactionAmountRequest) (*proto.GetTotalTransactionAmountResponse, error) {
	start := time.Now()
	transactions, err := db.GetTransactions(s.rdb)
	if err != nil {
		panic(err)
	}
	var totalAmount int64
	if request.GetIsParallel() {
		processTransactionInParallel(start, transactions)
	} else {
		for _, transaction := range transactions {
			totalAmount = totalAmount + transaction.Amount
			waitTwentyMilliseconds()
		}
	}
	return &proto.GetTotalTransactionAmountResponse{TotalAmount: totalAmount, TotalTime: int64(time.Now().Sub(start)) / int64(time.Millisecond)}, nil
}

func processTransactionInParallel(startTime time.Time ,transactions []db.Transaction) int64 {
	var results = make(chan int64)
	var totalAmount int64
	// Setup wait group to process all transactions
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(transactions))

	for _, transaction := range transactions {

		go func(transaction db.Transaction) {
			results <- transaction.Amount
			waitTwentyMilliseconds()
			defer waitGroup.Done()
		}(transaction)
		
	}
	// monitor when all work is done
	go func() {
		waitGroup.Wait()
		close(results)
		log.Printf("\n total amount: %d \n total time used in ms: %d", totalAmount, int64(time.Now().Sub(startTime)) / int64(time.Millisecond))
	}()

	for amount := range results {
		totalAmount = totalAmount + amount
	}
	
	select {
	case <-time.After(time.Duration(1) * time.Second):
		return totalAmount
	}
}

func waitTwentyMilliseconds() {
	time.Sleep(20 * time.Millisecond)
}