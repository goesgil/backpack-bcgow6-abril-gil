package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	transactions []TransactionsRequest
)

type TransactionsRequest struct {
	Id          int     `json:"id"`
	CodeTrxs    string  `json:"codeTrxs" binding:"required" nullable:"false"`
	Coin        string  `json:"coin" binding:"required" nullable:"false"`
	Amount      float64 `json:"amount" binding:"required" nullable:"false"`
	Transmitter string  `json:"transmitter" binding:"required" nullable:"false"`
}

func ValidateSession(c *gin.Context) {
	if token := c.GetHeader("Authorization"); token == "" || strings.Split(token, " ")[1] != "123456" {
		c.JSON(401, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
		c.Abort()
	}
	c.Next()
}

func handleFunc(c *gin.Context) {
	var req TransactionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.Id = len(transactions) + 1
	transactions = append(transactions, req)
	c.JSON(201, gin.H{"message": "Transaction created"})
}

func main() {
	router := gin.Default()

	router.POST("/transactions", ValidateSession, handleFunc)
	router.Run()
}
