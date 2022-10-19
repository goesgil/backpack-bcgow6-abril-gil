package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_2_TT/internal"
)

func NewController(s internal.Service) Controller {
	return &controller{
		service: s,
	}
}

type Controller interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	service internal.Service
}

func (c *controller) Create(ctx *gin.Context) {
	body := struct {
		Amount      float64 `json:"amount" binding:"required"`
		CodeTrxs    string  `json:"code_trxs" binding:"required"`
		Coin        string  `json:"coin" binding:"required"`
		Transmitter string  `json:"transmitter" binding:"required"`
	}{}
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in body, the amount must be float and the other fields must be string. All fields are required",
		})
		return
	}
	trxs, err := c.service.Create(body.Amount, body.CodeTrxs, body.Coin, body.Transmitter)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": trxs,
	})
	return
}

func (c *controller) Update(ctx *gin.Context) {
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
	err = c.service.Update(body.Amount, id, body.CodeTrxs, body.Coin, body.Transmitter)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
	return
}

func (c *controller) GetAll(ctx *gin.Context) {
	transactions := c.service.GetAll()
	ctx.JSON(http.StatusOK, gin.H{
		"data":  transactions,
		"count": len(transactions),
	})
	return
}

func (c *controller) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in params, the id must be number integer",
		})
		return
	}
	err = c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
	return
}
