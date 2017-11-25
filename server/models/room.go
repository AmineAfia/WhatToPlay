package models

import (
	"log"
	"sort"

	"github.com/zmb3/spotify"
)

// Room represents a room
type Room struct {
	Name  string          `json:"name"`
	Songs map[string]Song `json:"songs"` //maps a songid to its class
	//Token  string          `json:"-"`
	Client        spotify.Client `json:"-"`
	UserID        string         `json:"hostid"`
	PlaylistID    spotify.ID     `json:"playlist"`
	RenewTreshold int            `json:"-"`
}

func (s Song) Upvote(user string) {
	s.Votes[user] = Up
}
func (s Song) Downvote(user string) {
	s.Votes[user] = Down
}

func (s Song) CalculateVotes() int {
	res := 0

	for _, v := range s.Votes {
		if v == Up {
			res++
		} else {
			res--
		}
	}
	s.VotesI = res
	return res
}

// Song represents the item in the list
type Song struct {
	ID     string                   `json:"id"`
	Title  string                   `json:"title"`
	Artist string                   `json:"artist"`
	Votes  map[string]voteDirection `json:"-"`
	VotesI int                      `json:"votes"`
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

			if _, ok := r.Songs[id]; !ok {
				r.Songs[id] = Song{ID: id, Title: title, Artist: artist, Index: i, Votes: make(map[string]voteDirection)}
			}
		}
	}
}

func (r *Room) CheckUpvotes() {

	//badSongs := make(map[string]Song)

	for id, s := range r.Songs {
		if s.CalculateVotes() < 0 {
			//badSongs[id] = s
			delete(r.Songs, id)
		}
	}
	/*
		for id, _ := range badSongs {
			delete(r.Songs, id)
		}*/

	if len(r.Songs) < r.RenewTreshold {
		r.AddMoreSongs()
	}
}

func rankByUpvotes(upvotes map[string]int) PairList {
	pl := make(PairList, len(upvotes))
	i := 0
	for k, v := range upvotes {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func (r *Room) AddMoreSongs() {
	if len(r.Songs) == 0 {
		return
	}

	bestSongs := make(map[string]int)

	for id, song := range r.Songs {
		bestSongs[id] = song.CalculateVotes()
	}

	pairs := rankByUpvotes(bestSongs)

	songIds := []spotify.ID{}

	for i, pair := range pairs {
		if i == 4 {
			break
		}
		songIds = append(songIds, spotify.ID(pair.Key))
		log.Println(songIds)
	}

	seed := spotify.Seeds{Tracks: songIds}
	trackAttributes := spotify.NewTrackAttributes()
	limit := 5
	options := spotify.Options{Limit: &limit}
	res, err := r.Client.GetRecommendations(seed, trackAttributes, &options)
	ids := []spotify.ID{}
	if err == nil {
		for i, song := range res.Tracks {
			sid := song.ID
			title := song.Name
			artist := song.Artists[0].Name
			id := string(sid)

			log.Println("nusong===   sid: ", sid, " title: ", title, " artist: ", artist, " i: ", i)

			r.Songs[id] = Song{ID: id, Title: title, Artist: artist, Index: i, Votes: make(map[string]voteDirection)}
			ids = append(ids, song.ID)

		}
	}

	r.Client.AddTracksToPlaylist(r.UserID, r.PlaylistID, ids...)
}

func (r Room) CreateSong(id string) Song {
	var song Song
	song = r.Songs[id]
	if song.ID == "" {
		r.Songs[id] = Song{ID: id, Title: id, Votes: make(map[string]voteDirection)}
	}
	return r.Songs[id]
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
