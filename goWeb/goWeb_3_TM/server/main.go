package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_3_TM/internal"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_3_TM/server/handler"
)

func main() {
	// Generate mock data
	internal.GenerateMock()

	// Started task
	repo := internal.NewRepository()
	service := internal.NewService(repo)
	controller := handler.NewController(service)

	router := gin.Default()

	groupV1 := router.Group("/v1")
	groupTrxs := groupV1.Group("/transactions")
	{
		groupTrxs.PATCH("/:id", controller.Patch)
		groupTrxs.PUT("/:id", controller.Put)
		groupTrxs.DELETE("/:id", controller.Delete)
	}
	router.Run()
}
