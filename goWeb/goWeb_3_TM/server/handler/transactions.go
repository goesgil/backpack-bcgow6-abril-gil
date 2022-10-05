package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_3_TM/internal"
)

func NewController(s internal.Service) Controller {
	return &controller{
		service: s,
	}
}

type Controller interface {
	Put(ctx *gin.Context)
	Patch(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	service internal.Service
}

func (c *controller) Put(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in params, the id must be integer",
		})
		return
	}
	body := struct {
		Amount      float64 `json:"amount" binding:"required"`
		CodeTrxs    string  `json:"code_trxs" binding:"required"`
		Coin        string  `json:"coin" binding:"required"`
		Transmitter string  `json:"transmitter" binding:"required"`
	}{}
	err = ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in body, the amount must be float and the other fields must be string. All fields are required",
		})
		return
	}
	err = c.service.Put(body.Amount, id, body.CodeTrxs, body.Coin, body.Transmitter)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Error",
		})
		return
	}
	ctx.AbortWithStatus(http.StatusNoContent)
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Transaction updated",
	})
	return
}

func (c *controller) Patch(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in params, the id must be integer",
		})
		return
	}
	body := struct {
		Amount   float64 `json:"amount" binding:"required"`
		CodeTrxs string  `json:"code_trxs" binding:"required"`
	}{}
	err = ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in body, the amount must be float and the code_trxs must be string. All fields are required",
		})
		return
	}
	err = c.service.Patch(body.Amount, id, body.CodeTrxs)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Transaction does not exist",
		})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Transaction updated",
	})
	return
}

func (c *controller) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in params, the id must be integer",
		})
		return
	}
	err = c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Error",
		})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Transaction deleted",
	})
	return
}
