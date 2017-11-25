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

	c.Redirect(http.StatusSeeOther, config.Conf.BaseUrl+"qrs/"+uid+".png")
}

// the user will eventually be redirected back to your redirect URL
// typically you'll have a handler set up like the following:
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// use the same state string here that you used to generate the URL
	token, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusNotFound)
		return
	}
	// create a client using the specified token
	client = auth.NewClient(token)

	// the client can now be used to make authenticated requests
}
