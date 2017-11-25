package models

import (

)


// Room represents a room
type Room struct {
	Name	string				`json:"name`
	Songs	map[string]Song 	`json:"songs"`	//maps a songid to its class 
	Token	string 				`json:"-"`
}

func (s Song) Upvote(user string) {
	s.Votes[user] = Up
}
func (s Song) Downvote(user string) {
	s.Votes[user] =  Down
}

// Song represents the item in the list
type Song struct {
	ID		string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Votes 	map[string]voteDirection `json:"votes"`	

}
type voteDirection int 
const (
	Up voteDirection = iota
	Down
)

func (r Room) CreateSong(id string) Song {
	var song Song 
	song = r.Songs[id]
	if song.ID == "" {
		r.Songs[id] = Song{ID: id,	Title: id, Votes: make(map[string]voteDirection)}
	}
	return r.Songs[id]
}
