package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Incoming struct{}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.POST("/v1/transfers", func(c *gin.Context) {
		var req Incoming
		err := c.ShouldBind(&req)
		c.String(http.StatusOK, fmt.Sprintf("%+v %s", req, err))
	})
	r.GET("/v1/transfers/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, id)
	})

	_ = r.Run(":9090")
}
