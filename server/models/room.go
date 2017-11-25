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
	Title	string
	Artist	string
	Votes 	map[string]voteDirection	
}
type voteDirection int 
const (
	Up voteDirection = iota
	Down
)
