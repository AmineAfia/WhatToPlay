package models

type DB struct {
	Rooms	map[string]*Room
}

func (db DB) GetRoom(id string) *Room {
	return db.Rooms[id]
}

var Data DB 

func InitDB()  {
	Data = DB{}
}