package models

import "github.com/AmineAfia/WhatToPlay/server/qrcode"

type DB struct {
<<<<<<< HEAD
	Rooms	map[string]Room
}

func (db DB) GetRoom(id string) Room {
=======
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

>>>>>>> 7384e8623721dacb08700b04d8a1de204ca6f546
	return db.Rooms[id]
}

var Data DB

func InitDB() {
	Data = DB{Rooms: make(map[string]*Room)}
}
