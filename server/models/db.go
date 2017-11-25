package models

type DB struct {
	Rooms	map[string]*Room
}

func (db DB) GetOrCreateRoom(id string) *Room {
	var room = db.Rooms[id]
	if room != nil {
		return room
	}
	db.Rooms[id] = &Room{Name: id, Songs: make(map[string]Song)}
	return db.Rooms[id]
}

var Data DB 

func InitDB()  {
	Data = DB{Rooms: make(map[string]*Room)}
}