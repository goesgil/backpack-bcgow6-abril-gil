package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/MockTrxs"
)

func ValidateQuerys(ctx *gin.Context, Key string, ValueToMatch string) bool {
	if ctx.Query(Key) == ValueToMatch {
		return true
	}
	return false
}

func HandlerHi(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hola Abril!"})
}

func GetAll(ctx *gin.Context) {
	var trxs mocktrxs.Transactions
	jsonTrxs, err := os.ReadFile("./goWeb/transactions.json")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error al leer el archivo"})
	} else {
		if err := json.Unmarshal(jsonTrxs, &trxs); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error al transpilar el archivo"})
		} else {
			var filterTrxs = []mocktrxs.Transaction{}
			for _, trx := range trxs.Transactions {
				if ValidateQuerys(ctx, "CodeTrxs", trx.CodeTrxs) {
					filterTrxs = append(filterTrxs, trx)
				}
			}
			ctx.JSON(http.StatusOK, filterTrxs)
		}
	}
}

func GetById(ctx *gin.Context) {
	var trxs mocktrxs.Transactions
	jsonTrxs, err := os.ReadFile("./goWeb/transactions.json")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error al leer el archivo"})
		return
	}
	if err := json.Unmarshal(jsonTrxs, &trxs); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error al transpilar el archivo"})
		return
	}
	indexById := 0
	for index, trx := range trxs.Transactions {
		if fmt.Sprint(trx.Id) == ctx.Param("id") {
			indexById = index
		}
	}
	if indexById == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No se encontró la transacción"})
		return
	}
	ctx.JSON(http.StatusOK, trxs.Transactions[indexById])

}

func main() {
	mocktrxs.Generate()

	router := gin.Default()

	router.GET("/v1/meeting", HandlerHi)
	groupV1 := router.Group("/v1")
	{
		groupV1.GET("/transactions", GetAll)
		groupV1.GET("/transactions/:id", GetById)
	}

	router.Run()
}
