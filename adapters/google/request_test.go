package google_test

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
	googleClientID, googleClientSecret string
)

func init() {
	godotenv.Load("../../.env.local")

	googleClientID = cmp.Or(os.Getenv("GOOGLE_CLIENT_ID"), "client_id")
	googleClientSecret = cmp.Or(os.Getenv("GOOGLE_SECRET"), "secret")

}

func getMyGoogle() gauth.AuthRequest {
	r, err := gauth.New().
		Source("google").
		ClientId(googleClientID).
		ClientSecret(googleClientSecret).
		RedirectUrl("http://localhost:8080/auth/google/callback").
		Build()
	if err != nil {
		log.Fatalf("Failed to build request:%v", err)
	}
	return r
}

func TestGoogleAuthRequest_Authorize(t *testing.T) {
	r := getMyGoogle()

	state := utils.CreateState()
	authorizeUrl, err := r.Authorize(state)

	assert.Nil(t, err)

	u, authErr := url.Parse(authorizeUrl)
	assert.NoError(t, authErr)
	assert.Equal(t, "accounts.google.com", u.Host)
	values := u.Query()
	assert.Equal(t, googleClientID, values.Get("client_id"))
	assert.Equal(t, state, values.Get("state"))
	assert.Equal(t, "http://localhost:8080/auth/google/callback", values.Get("redirect_uri"))
	assert.Equal(t, "openid email profile", values.Get("scope"))
}
