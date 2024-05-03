package facebook_test

import (
	"cmp"
	"github.com/LeoInnovateLab/gauth"
	_ "github.com/LeoInnovateLab/gauth/register"
	"github.com/LeoInnovateLab/gauth/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/url"
	"os"
	"testing"
)

var (
	clientID, secret string
)

func init() {
	godotenv.Load("../../.env.local")

	clientID = cmp.Or(os.Getenv("FACEBOOK_APP_ID"), "client_id")
	secret = cmp.Or(os.Getenv("FACEBOOK_APP_SECRET"), "secret")
}

func getMyFacebook() gauth.AuthRequest {
	r, err := gauth.New().
		Source("facebook").
		ClientId(clientID).
		ClientSecret(secret).
		RedirectUrl("http://localhost:8080/auth/facebook/callback").
		Build()
	if err != nil {
		log.Fatalf("Failed to build request:%v", err)
	}
	return r
}

func TestFacebookAuthRequest_Authorize(t *testing.T) {
	r := getMyFacebook()

	state := utils.CreateState()
	authorizeUrl, err := r.Authorize(state)

	assert.Nil(t, err)

	u, authErr := url.Parse(authorizeUrl)
	assert.NoError(t, authErr)
	assert.Equal(t, "www.facebook.com", u.Host)
	values := u.Query()
	assert.Equal(t, clientID, values.Get("client_id"))
	assert.Equal(t, state, values.Get("state"))
	assert.Equal(t, "http://localhost:8080/auth/facebook/callback", values.Get("redirect_uri"))
	assert.Equal(t, "public_profile", values.Get("scope"))
}
