package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	proto "github.com/knzou/Budget/proto"
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
        g.GET("/getPeople", func(ctx *gin.Context) {
                var req proto.GetPeopleRequest
                ctx.BindJSON(&req)
                if response, err := client.GetPeople(ctx, &req); err == nil {
                        ctx.JSON(http.StatusOK, gin.H{
                                "result": response,
                                "req": req,
                        })
                } else {
                        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                } 
        })
        g.GET("/getTotalTransactionAmount", func(ctx *gin.Context) {
                var req proto.GetTotalTransactionAmountRequest
                ctx.BindJSON(&req)
                if response, err := client.GetTotalTransactionAmount(ctx, &req); err == nil {
                        ctx.JSON(http.StatusOK, gin.H{
                                "result": response,
                                "req": req,
                        })
                } else {
                        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                } 
        })

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
