package main

import (
	"fmt"
	"net/http"
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no tiene permisos para realizar la petición solicitada"})
	}
	fmt.Print("http")
	c.Next()
}

func handleFunc(c *gin.Context) {
	var req TransactionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req.Id = len(transactions) + 1
	transactions = append(transactions, req)
	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created"})
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.Use(ValidateSession)
	groupTrxs := v1.Group("/transactions")
	groupTrxs.POST("/", ValidateSession, handleFunc)
	groupTrxs.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, transactions)
	})

	router.Run()
}
