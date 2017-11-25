package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Run(address string) {
	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/sanity")

	router.Run(address)
}


func sanityCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}