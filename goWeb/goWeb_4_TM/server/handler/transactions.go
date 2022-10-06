package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_4_TM/internal"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_4_TM/pkg/web"
)

func NewController(s internal.Service) Controller {
	return &controller{
		service: s,
	}
}

type Controller interface {
	Create(ctx *gin.Context)
	Put(ctx *gin.Context)
	Patch(ctx *gin.Context)
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
		ctx.JSON(http.StatusBadRequest, web.CustomResponse(http.StatusBadRequest, nil, "Error in body, the amount must be float and the other fields must be string. All fields are required"))
		return
	}
	trxs, err := c.service.Create(body.Amount, body.CodeTrxs, body.Coin, body.Transmitter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.CustomResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, web.CustomResponse(http.StatusCreated, trxs, ""))
	return
}

func (c *controller) Put(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.CustomResponse(http.StatusBadRequest, nil, "Error in params, the id must be integer"))
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
		ctx.JSON(http.StatusBadRequest, web.CustomResponse(http.StatusBadRequest, nil,
			"Error in body, the amount must be float and the other fields must be string. All fields are required"))
		return
	}
	err = c.service.Put(body.Amount, id, body.CodeTrxs, body.Coin, body.Transmitter)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.CustomResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
	return
}

func (c *controller) Patch(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.CustomResponse(http.StatusBadRequest, nil, "Error in params, the id must be integer"))
		return
	}
	body := struct {
		Amount   float64 `json:"amount" binding:"required"`
		CodeTrxs string  `json:"code_trxs" binding:"required"`
	}{}
	err = ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.CustomResponse(http.StatusBadRequest, nil,
			"Error in body, the amount must be float and the other fields must be string. All fields are required"))
		return
	}
	err = c.service.Patch(body.Amount, id, body.CodeTrxs)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.CustomResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
	return
}

func (c *controller) GetAll(ctx *gin.Context) {
	transactions := c.service.GetAll()
	ctx.JSON(http.StatusOK, web.CustomResponse(http.StatusOK, gin.H{
		"data":  transactions,
		"count": len(transactions),
	}, ""))
	return
}

func (c *controller) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Params.ByName("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.CustomResponse(http.StatusBadRequest, nil, "Error in params, the id must be integer"))
		return
	}
	err = c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.CustomResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
	return
}
