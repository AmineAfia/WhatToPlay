package models

import "github.com/AmineAfia/WhatToPlay/server/qrcode"

type DB struct {
	Rooms map[string]*Room
}

const hostname = "http://localhost:8080/"

func (db DB) GetOrCreateRoom(id string) *Room {
	var room = db.Rooms[id]
	if room != nil {
		return room
	}
	db.Rooms[id] = &Room{Name: id, Songs: make(map[string]Song)}

	qrcode.CreateQr(hostname, id)

	return db.Rooms[id]
}

var Data DB

func InitDB() {
	Data = DB{Rooms: make(map[string]*Room)}
}
