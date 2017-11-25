package router

import (
	"net/http"

	"github.com/AmineAfia/WhatToPlay/server/models"
	"github.com/AmineAfia/WhatToPlay/server/qrcode"
	"github.com/gin-gonic/gin"
)

// Run runs the router on the specified address
func Run(address string) {
	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/sanity", sanityCheck)
	v1.GET("/room/:room", getRoom)
	v1.POST("/room/:room", addSong)

	router.StaticFS("/qrs", http.Dir("qrs"))

	qrcode.CreateQr("http://localhost/", "2345")

	router.Run(address)
}

func sanityCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getRoom(c *gin.Context) {
	room := c.Param("room")
	c.JSON(200, models.Data.GetRoom(room))
}

func addSong(c *gin.Context) {
	room := c.Param("room")
	r := models.Data.GetRoom(room)
	r.Songs[c.Param("song")] = models.Song{}
	c.JSON(200, r)
}
