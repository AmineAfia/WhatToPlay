package models

import (
	"github.com/AmineAfia/WhatToPlay/server/config"
	"github.com/AmineAfia/WhatToPlay/server/qrcode"
)

type DB struct {
	Rooms map[string]*Room
}

func (db DB) GetOrCreateRoom(id string) *Room {
	var room = db.Rooms[id]
	if room != nil {
		return room
	}
	db.Rooms[id] = &Room{Name: id, Songs: make(map[string]Song)}

	qrcode.CreateQr(config.Conf.BaseUrl, id)

	return db.Rooms[id]
}

var Data DB

func InitDB() {
	Data = DB{Rooms: make(map[string]*Room)}
}
