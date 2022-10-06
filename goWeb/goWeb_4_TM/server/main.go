package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_4_TM/internal"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_4_TM/pkg/db"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_4_TM/server/handler"
)

func main() {
	err := db.GenerateMock()
	if err != nil {
		fmt.Println(err.Error())
	}
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
