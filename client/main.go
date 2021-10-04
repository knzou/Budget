package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/knzou/BudgetCalculatorService/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewAddServiceClient(conn)

	g := gin.Default()

        g.GET("/getCategories", func(ctx *gin.Context) {
                req := &proto.Request{}
                if response, err := client.GetCategories(ctx, req); err == nil {
                        ctx.JSON(http.StatusOK, gin.H{
                                "result": response,
                        })
                } else {
                        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                } 
        })
        g.GET("/getTransactions", func(ctx *gin.Context) {
                req := &proto.Request{}
                if response, err := client.GetTransactions(ctx, req); err == nil {
                        ctx.JSON(http.StatusOK, gin.H{
                                "result": response,
                        })
                } else {
                        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                } 
        })

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
