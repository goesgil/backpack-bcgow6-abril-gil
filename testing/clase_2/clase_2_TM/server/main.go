package main

import (
	"fmt"
	//"os"

	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_2/clase_2_TM/internal"
	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_2/clase_2_TM/pkg/db"
	//"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_2/clase_2_TM/docs"
	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_2/clase_2_TM/server/handler"
	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_2/clase_2_TM/server/middleware"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @contact.name Goesgil
// @contact.url http://www.example.com
// @host localhost:8080
// @BasePath /api/v1

func main() {
	_ = godotenv.Load()
	err := db.GenerateMock()
	if err != nil {
		fmt.Println(err.Error())
	}
	db := db.NewDB()
	repo := internal.NewRepository(db)
	service := internal.NewService(repo)
	controller := handler.NewController(service)
	
	router := gin.Default()
	//docs.SwaggerInfo.Host = os.Getenv("HOST")
	
	groupV1 := router.Group("/api/v1")
	groupV1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	groupTrxs := groupV1.Group("/transactions", middleware.ValidateSession)
	{
		groupTrxs.POST("/", controller.Create)
		groupTrxs.PATCH("/:id", controller.Patch)
		groupTrxs.PUT("/:id", controller.Put)
		groupTrxs.GET("/", controller.GetAll)
		groupTrxs.DELETE("/:id", controller.Delete)
	}
	
	router.Run("localhost:8081")
}
