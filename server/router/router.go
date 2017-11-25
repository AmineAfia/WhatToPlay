package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/AmineAfia/WhatToPlay/server/models"
	
)

// Run runs the router on the specified address
func Run(address string) {
	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/sanity", sanityCheck)
	v1.GET("/room/:room", getRoom)
	v1.POST("/playlist", )

	router.Run(address)
}


func sanityCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getRoom(c *gin.Context) {
	room := c.Param("room")
	c.JSON(200, models.Data.GetRoom(room))
}