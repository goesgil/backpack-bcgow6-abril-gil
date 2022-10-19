package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_2_TT/internal"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_2_TT/server/handler"
)

func main() {
	repo := internal.NewRepository()
	service := internal.NewService(repo)
	controller := handler.NewController(service)

	router := gin.Default()

	groupV1 := router.Group("/v1")
	groupTrxs := groupV1.Group("/transactions")
	{
		groupTrxs.POST("/", controller.Create)
		groupTrxs.PATCH("/:id", controller.Update)
		groupTrxs.GET("/", controller.GetAll)
		groupTrxs.DELETE("/:id", controller.Delete)
	}

	router.Run()
}
