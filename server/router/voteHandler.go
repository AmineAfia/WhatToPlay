package router 

import (
	"github.com/AmineAfia/WhatToPlay/server/models"
)

func Upvote(c *gin.Context) {
	room :=	GetRoom("RoomId")
	room.Upvote("songid", "userid")
} 

func Downvote(c *gin.Context) {
	room :=	GetRoom("RoomId")
	room.Downvote("songid", "userid")
}

/* 
AddSong() 

Upvote() 
POST /api/v1/
Header: Bearer: 	uuid //user token
ID		string //song id
RoomId 	string 


Downvote() 


*/

func GetRoom(id string) models.Room {
	return models.Room{}
}