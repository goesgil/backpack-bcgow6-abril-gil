package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_3/cmd/server/handler"
	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_3/internal/transactions"
	"github.com/goesgil/backpack-bcgow6-abril-gil/testing/clase_3/pkg/store"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	r := gin.Default()
	db := store.NewFileStore("./trxs.json")
	repo := transactions.NewRepository(db)
	serv := transactions.NewService(repo)
	p := handler.NewTrxs(serv)

	pr := r.Group("transactions")
	pr.GET("/", p.GetAll())

	if err := r.Run(); err != nil {
		panic(err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
