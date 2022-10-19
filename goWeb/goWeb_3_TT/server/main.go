package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_3_TT/internal"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_3_TT/server/handler"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_3_TT/pkg/db"
)

func main() {
	db := db.NewDB()
	repo := internal.NewRepository(db)
	service := internal.NewService(repo)
	controller := handler.NewController(service)

	router := gin.Default()

	groupV1 := router.Group("/v1")
	groupTrxs := groupV1.Group("/transactions")
	{
		groupTrxs.POST("/", controller.Create)
		groupTrxs.PATCH("/:id", controller.Patch)
		groupTrxs.PUT("/:id", controller.Put)
		groupTrxs.GET("/", controller.GetAll)
		groupTrxs.DELETE("/:id", controller.Delete)
	}

	router.Run()
}
