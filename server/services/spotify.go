package services
/*

import ( 
	"net/http"
	"github.com/zmb3/spotify"
)

func RunCallbackServer(address string) {
	auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate)
	
	
	auth.SetAuthInfo(clientID, secretKey)
	
	// get the user to this URL - how you do that is up to you
	// you should specify a unique state string to identify the session
	url := auth.AuthURL(state)
	
	http.ServeMux  

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
      client := auth.NewClient(token)

      // the client can now be used to make authenticated requests
}
*/