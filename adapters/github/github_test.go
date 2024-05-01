package github_test

import (
	"cmp"
	"github.com/LeoInnovateLab/gauth"
	"github.com/LeoInnovateLab/gauth/utils"
	"github.com/joho/godotenv"
	"log"
	"net/url"
	"os"
	"testing"

	_ "github.com/LeoInnovateLab/gauth/register"
	"github.com/stretchr/testify/assert"
)

var (
	gitlabClientID, gitlabSecret string
)

func init() {
	godotenv.Load("../../.env.local")

	gitlabClientID = cmp.Or(os.Getenv("GITHUB_CLIENT_ID"), "client_id")
	gitlabSecret = cmp.Or(os.Getenv("GITHUB_SECRET"), "secret")

}

func TestGithubAuthRequest_Authorize(t *testing.T) {
	r := getMyGithub()

	state := utils.CreateState()
	authorizeUrl, err := r.Authorize(state)

	assert.Nil(t, err)

	u, authErr := url.Parse(authorizeUrl)
	assert.NoError(t, authErr)
	assert.Equal(t, "github.com", u.Host)
	values := u.Query()
	assert.Equal(t, gitlabClientID, values.Get("client_id"))
	assert.Equal(t, state, values.Get("state"))
	assert.Equal(t, "http://localhost:8080/auth/github/callback", values.Get("redirect_uri"))
	assert.Equal(t, "", values.Get("scope"))
}

func getMyGithub() gauth.AuthRequest {
	r, err := gauth.New().
		Source("github").
		ClientId(gitlabClientID).
		ClientSecret(gitlabSecret).
		RedirectUrl("http://localhost:8080/auth/github/callback").
		Build()
	if err != nil {
		log.Fatalf("Failed to build request:%v", err)
	}
	return r
}

func TestMyGithub(t *testing.T) {
	r := getMyGithub()

	state := utils.CreateState()
	authorizeUrl, err := r.Authorize(state)
	assert.Nil(t, err)
	assert.NotNil(t, authorizeUrl)
}

func TestRevoke(t *testing.T) {
	r := getMyGithub()
	token := gauth.AuthToken{}
	err := r.Revoke(token)
	assert.NotNil(t, err)
	assert.Equal(t, gauth.ErrNotImplemented, err)
}
