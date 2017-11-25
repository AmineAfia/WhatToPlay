package models

import (

)


// Room represents a room
type Room struct {
	Name	string
	Songs	map[string]Song //maps a songid to its class
}

func (r Room) Upvote(song string, user string) {
	r.Songs[song].Votes[user] = Up
}
func (r Room) Downvote(song string, user string) {
	r.Songs[song].Votes[user] =  Down
}

// Song represents the item in the list
type Song struct {
	ID		string
	Title	string
	Artist	string
	Votes 	map[string]voteDirection	
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