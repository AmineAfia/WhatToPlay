package router

import (
	"net/http"

	"github.com/AmineAfia/WhatToPlay/server/models"
	"github.com/gin-gonic/gin"
)

// Run runs the router on the specified address
func Run(address string) {
	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/sanity", sanityCheck)
	v1.GET("/room", dumpDb)
	v1.GET("/room/:room", getRoom)
<<<<<<< HEAD
	v1.POST("/playlist", )
=======
	v1.POST("/room/:room", createRoom) //use the spotify userid as the room name
	v1.POST("/room/:room/songs/:song", addSong)

	v1.POST("/room/:room/songs/:song/upvote", upvote)
	v1.POST("/room/:room/songs/:song/downvote", downvote)

	router.StaticFS("/qrs", http.Dir("qrs"))
>>>>>>> 7384e8623721dacb08700b04d8a1de204ca6f546

	router.Run(address)
}

func sanityCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getRoom(c *gin.Context) {
	room := c.Param("room")
<<<<<<< HEAD
	c.JSON(200, models.Data.GetRoom(room))
}
=======
	c.JSON(200, models.Data.GetOrCreateRoom(room))
}

/*
func addSong(c *gin.Context) {
	room := c.Param("room")
	r := models.Data.GetOrCreateRoom(room)
	r.Songs[c.Param("song")] = models.Song{}
	c.JSON(200, r)
} */
func createRoom(c *gin.Context) {
	room := c.Param("room")
	r := models.Data.GetOrCreateRoom(room)
	c.JSON(200, r)
}
func dumpDb(c *gin.Context) {
	c.JSON(200, models.Data)
}
func addSong(c *gin.Context) {
	room := c.Param("room")
	r := models.Data.GetOrCreateRoom(room)
	r.CreateSong(c.Param("song"))
	c.JSON(200, r)
}

func upvote(c *gin.Context) {
	room := c.Param("room")
	r := models.Data.GetOrCreateRoom(room)
	s := r.CreateSong(c.Param("song"))
	s.Upvote(c.DefaultQuery("user", "Guest"))
	c.JSON(200, r)
}
func downvote(c *gin.Context) {
	room := c.Param("room")
	r := models.Data.GetOrCreateRoom(room)
	s := r.CreateSong(c.Param("song"))
	s.Downvote(c.DefaultQuery("user", "Guest"))
	c.JSON(200, r)
}

>>>>>>> 7384e8623721dacb08700b04d8a1de204ca6f546
