package services

import (
	"net/http"

	"github.com/AmineAfia/WhatToPlay/server/config"
	"github.com/AmineAfia/WhatToPlay/server/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/zmb3/spotify"
)

const redirectURI = "callback"

var (
	auth   = spotify.Authenticator{}
	setUp  = false
	state  = "nil"
	client spotify.Client
)

func SetUpAuth() spotify.Authenticator {
	res := spotify.NewAuthenticator(config.Conf.BaseUrl+redirectURI,
		spotify.ScopeUserReadPrivate,
		spotify.ScopePlaylistReadPrivate,
		spotify.ScopePlaylistModifyPrivate,
		spotify.ScopePlaylistModifyPublic,
		spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserModifyPlaybackState)
	res.SetAuthInfo(config.Conf.SpotifyAppClientId, config.Conf.SpotifyAppClientSecret)
	setUp = true
	return res
}

func Auth(c *gin.Context) {
	if !setUp {
		auth = SetUpAuth()
	}
	url := auth.AuthURL(state)
	c.Redirect(http.StatusSeeOther, url)
}

func CallbHandler(c *gin.Context) {

	code := c.Query("code")   // code
	state := c.Query("state") // state

	if code == "" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	uid := uuid.NewV4().String()

	r := models.Data.GetOrCreateRoom(uid)
	//r.Token = code
	token, _ := auth.Token(state, c.Request)
	r.Client = auth.NewClient(token)

	/*
		user, err := client.GetUsersPublicProfile(spotify.ID(*userID))
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return
		}*/

	user, err := r.Client.CurrentUser()

	if err != nil {
		panic(err)
	} else {
		r.UserID = user.User.ID
	}

	r.FindOrCreatePlaylist()
	r.UpdatePlaylistSongs()

	c.Redirect(http.StatusSeeOther, config.Conf.BaseUrl+"qrs/"+uid+".png")
}
