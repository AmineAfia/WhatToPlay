package models

import (

)


// Room represents a room
type Room struct {
	Name	string
	Songs	[]Song
}

// Song represents the item in the list
type Song struct {
	Name	string
	ID		string
	Votes	int
}