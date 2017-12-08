package router

import (
	"net/http"
	"time"

	"github.com/AmineAfia/WhatToPlay/server/models"
	"github.com/AmineAfia/WhatToPlay/server/services"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
)

// Run runs the router on the specified address
func Run(address string) {
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.StaticFS("front/", http.Dir("front/public_html"))

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/sanity", sanityCheck)
	v1.GET("/room", dumpDb)
	v1.GET("/room/:room", getRoom)
	v1.GET("/room/:room/update", updateRoom)

	v1.POST("/room/:room", createRoom) //use the spotify userid as the room name
	v1.POST("/room/:room/songs/:song", addSong)

	v1.POST("/room/:room/songs/:song/upvote", upvote)
	v1.POST("/room/:room/songs/:song/downvote", downvote)

	router.GET("/auth", services.Auth)
	router.GET("/callback", services.CallbHandler)

	router.StaticFS("/qrs", http.Dir("qrs"))

	router.Run(address)
}

func sanityCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func getRoom(c *gin.Context) {
	room := c.Param("room")
	r := models.Data.GetOrCreateRoom(room)
	r.UpdatePlaylistSongs()
	c.JSON(200, r)
}

func updateRoom(c *gin.Context) {
	room := c.Param("room")
	r := models.Data.GetOrCreateRoom(room)
	r.UpdatePlaylistSongs()
	r.CheckUpvotes()
	c.JSON(200, r)
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
