package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goesgil/backpack-bcgow6-abril-gil/goWeb/goWeb_4_TT/pkg/web"
)

func ValidateSession(c *gin.Context) {
	fmt.Println(c.GetHeader("Authorization"))
	fmt.Println("c.GetHeader")
	fmt.Println(c.Request.Header.Get("Authorization"))
	fmt.Println("c.Request.Header.Get")
	if token := c.GetHeader("Authorization"); token == "" || strings.Split(token, " ")[1] != "123456" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, web.CustomResponse(http.StatusUnauthorized, nil, "no tiene permisos para realizar la petición solicitada"))
		return
	}
	c.Next()
}
