package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	transactions []TransactionsRequest
)

type TransactionsRequest struct {
	Id          int     `json:"id"`
	CodeTrxs    string  `json:"codeTrxs" binding:"required"`
	Coin        string  `json:"coin" binding:"required"`
	Amount      float64 `json:"amount" binding:"required,min=0.00000000"`
	Transmitter string  `json:"transmitter" binding:"required"`
}

func ValidateSession(c *gin.Context) {
	if token := c.GetHeader("Authorization"); token == "" || strings.Split(token, " ")[1] != "123456" {
		c.AbortWithStatusJSON(401, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
	}
	c.Next()
}

func handleFunc(c *gin.Context) {
	var req TransactionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	req.Id = len(transactions) + 1
	transactions = append(transactions, req)
	c.JSON(201, gin.H{"message": "Transaction created"})
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.Use(ValidateSession)
	groupTrxs := v1.Group("/transactions")
	groupTrxs.POST("/", ValidateSession, handleFunc)
	groupTrxs.GET("/", func(c *gin.Context) {
		c.JSON(200, transactions)
	})

	router.Run()
}
