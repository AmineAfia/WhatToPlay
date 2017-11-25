package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Run runs the router on the specified address
func Run(address string) {
	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/sanity", sanityCheck)
	v1.GET("/playlist", getPlaylist)
	v1.POST("/playlist", )

	router.Run(address)
}


func sanityCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getPlaylist(c *gin.Context) {
	c.Status(http.StatusOK)
}