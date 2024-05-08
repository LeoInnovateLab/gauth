package slack_test

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
	clientID, clientSecret string
)

func init() {
	godotenv.Load("../../.env.local")

	clientID = cmp.Or(os.Getenv("SLACK_CLIENT_ID"), "client_id")
	clientSecret = cmp.Or(os.Getenv("SLACK_SECRET"), "secret")
}

func getMySlack() gauth.AuthRequest {
	r, err := gauth.New().
		Source("slack").
		ClientId(clientID).
		ClientSecret(clientSecret).
		RedirectUrl("http://localhost:8080/auth/slack/callback").
		Build()
	if err != nil {
		log.Fatalf("Failed to build request:%v", err)
	}
	return r
}

func TestSlackAuthRequest_Authorize(t *testing.T) {
	r := getMySlack()

	state := utils.CreateState()
	authorizeUrl, err := r.Authorize(state)

	assert.Nil(t, err)

	u, authErr := url.Parse(authorizeUrl)
	assert.NoError(t, authErr)
	assert.Equal(t, "slack.com", u.Host)
	values := u.Query()
	assert.Equal(t, clientID, values.Get("client_id"))
	assert.Equal(t, state, values.Get("state"))
	assert.Equal(t, "http://localhost:8080/auth/slack/callback", values.Get("redirect_uri"))
	assert.Equal(t, "users.profile:read users:read users:read.email", values.Get("scope"))
}
