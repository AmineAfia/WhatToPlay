package models

import (
	"log"

	"github.com/zmb3/spotify"
)

// Room represents a room
type Room struct {
	Name  string          `json:"name"`
	Songs map[string]Song `json:"songs"` //maps a songid to its class
	//Token  string          `json:"-"`
	Client     spotify.Client `json:"-"`
	UserID     string         `json:"hostid"`
	PlaylistID spotify.ID     `json:"playlist"`
}

func (s Song) Upvote(user string) {
	s.Votes[user] = Up
}
func (s Song) Downvote(user string) {
	s.Votes[user] = Down
}

// Song represents the item in the list
type Song struct {
	ID     string                   `json:"id"`
	Title  string                   `json:"title"`
	Artist string                   `json:"artist"`
	Votes  map[string]voteDirection `json:"votes"`
	Index  int                      `json:"index"`
}
type voteDirection int

const (
	Up voteDirection = iota
	Down
)

func (r *Room) FindOrCreatePlaylist() {
	//TODO
	list, err := r.Client.GetPlaylistsForUser(r.UserID)

	found := spotify.ID("")
	if err == nil {
		for _, plist := range list.Playlists {
			log.Println("playlistfound: ", plist.Name)

			if plist.Name == "WhatToPlay" {
				found = plist.ID
			}
		}
	}
	if found.String() == "" {
		plist, _ := r.Client.CreatePlaylistForUser(r.UserID, "WhatToPlay", true)
		found = plist.ID
	}

	log.Println("playlistid: ", found)

	r.PlaylistID = found
}

func (r *Room) UpdatePlaylistSongs() {
	list, err := r.Client.GetPlaylistTracks(r.UserID, r.PlaylistID)

	if err == nil && list != nil {
		for i, song := range list.Tracks {
			sid := song.Track.ID
			title := song.Track.Name
			artist := song.Track.Artists[0].Name
			id := string(sid)

			log.Println("song===   sid: ", sid, " title: ", title, " artist: ", artist, " i: ", i)

			r.Songs[id] = Song{ID: id, Title: title, Artist: artist, Index: i, Votes: make(map[string]voteDirection)}
		}
	}
}

func (r Room) CreateSong(id string) Song {
	var song Song
	song = r.Songs[id]
	if song.ID == "" {
		r.Songs[id] = Song{ID: id, Title: id, Votes: make(map[string]voteDirection)}
	}
	return r.Songs[id]
}
