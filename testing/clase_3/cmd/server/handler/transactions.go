package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_3/internal/transactions"
	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_3/pkg/web"
)

type Trxs struct {
	service transactions.Service
}

type request struct {
	CodeTrxs    string  `json:"codeTrxs"`
	Coin        string  `json:"coin"`
	Amount      float64 `json:"amount"`
	Transmitter string  `json:"transmitter"`
}

func NewTrxs(s transactions.Service) *Trxs {
	return &Trxs{service: s}
}

func (t *Trxs) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, ""))
			return
		}

		trxs, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(404, nil, ""))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(200, trxs, ""))
	}
}

func (t *Trxs) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, ""))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}

		trx, err := t.service.Update(int(id), req.Amount, req.CodeTrxs, req.Coin, req.Transmitter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}

		ctx.JSON(http.StatusOK, trx)
	}
}

func (t *Trxs) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(401, nil, "token inválido")) //401
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error())) //400
			return
		}

		err = t.service.Delete(int(id))
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(404, nil, err.Error())) //404
			return
		}

		c.JSON(http.StatusNoContent, web.NewResponse(204, nil, ""))
	}
}
