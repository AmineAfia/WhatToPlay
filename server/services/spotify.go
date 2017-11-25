package services

import (
	"net/http"

	"github.com/AmineAfia/WhatToPlay/server/config"
	"github.com/AmineAfia/WhatToPlay/server/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/zmb3/spotify"
)

const redirectURI = "http://localhost:8080/callback"

var (
	auth   = spotify.Authenticator{}
	setUp  = false
	state  = "abc123"
	client spotify.Client
)

func SetUpAuth() spotify.Authenticator {
	res := spotify.NewAuthenticator(config.Conf.SpotifyRedirectUrl, spotify.ScopeUserReadPrivate)
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

	_ = c.Query("code")  // code
	_ = c.Query("state") // state

	uid := uuid.NewV4().String()

	_ = models.Data.GetOrCreateRoom(uid)

	c.Redirect(http.StatusSeeOther, "http://localhost:8080/room/"+uid)
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
