package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_2/clase_2_TM/internal"
	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_2/clase_2_TM/pkg/web"
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

// Transaction godoc
// @Summary Transactions
// @Schemes
// @Description Create a new transaction
// @Tags trxs
// @Security token header string "Bearer <token>"
// @Accept json
// @Produce json
// @Success 201 {object} web.response "Created"
// @Failure 401 {} web.response "Unauthorized"
// @Failure 400 {} web.response "Bad Request"
// @Failure 500 {} web.response "Internal Server Error"
// @Router /transactions [post]
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

// Transaction godoc
// @Summary Transactions
// @Schemes
// @Description Update a transaction
// @Tags trxs
// @Pamameter id path int true "Transaction ID"
// @Security token header string "Bearer <token>"
// @Accept json
// @Success 204
// @Failure 401 {object} web.response "Unauthorized"
// @Failure 400 {object} web.response "Bad Request"
// @Failure 500 {object} web.response "Internal Server Error"
// @Router /transactions/{id} [put]
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

// Transaction godoc
// @Summary Transactions
// @Schemes
// @Description Update codeTrxs and amount of a transaction
// @Tags trxs
// @Pamameter id path int true "Transaction ID"
// @Security token header string "Bearer <token>"
// @Accept json
// @Success 204
// @Failure 401 {object} web.response "Unauthorized"
// @Failure 400 {object} web.response "Bad Request"
// @Failure 500 {object} web.response "Internal Server Error"
// @Router /transactions/{id} [patch]
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

// Transaction godoc
// @Summary Transactions
// @Schemes
// @Description Get all transactions
// @Tags trxs
// @Pamameter id path int true "Transaction ID"
// @Security token header string "Bearer <token>"
// @Accept json
// @Produce json
// @Success 200 {object} web.PaginationResponse "Transactions"
// @Router /transactions [get]

func (c *controller) GetAll(ctx *gin.Context) {
	transactions := c.service.GetAll()
	var pagination *web.Pagination
	pagination.Data = transactions
	pagination.Count = len(transactions)
	ctx.JSON(http.StatusOK, web.PaginationResponse(http.StatusOK,
		pagination, ""))
	return
}

// Transaction godoc
// @Summary Transactions
// @Schemes
// @Description Update codeTrxs and amount of a transaction
// @Tags trxs
// @Pamameter id path int true "Transaction ID"
// @Security token header string "Bearer <token>"
// @Accept json
// @Success 204
// @Failure 401 {object} web.response "Unauthorized"
// @Failure 400 {object} web.response "Bad Request"
// @Failure 500 {object} web.response "Internal Server Error"
// @Router /transactions/{id} [delete]
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
