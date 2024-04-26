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
	err := godotenv.Load("../../.env.local")
	if err != nil {
		log.Println("Error loading .env file")
	}

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

func TestLogin(t *testing.T) {
	r := getMyGithub()
	callback := gauth.AuthCallback{
		Code:  "1eb6a1db53fd41e6640d",
		State: "12b87208-fbef-496b-b152-d328078dd746",
	}

	authToken, err := r.Login(callback)

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	log.Fatalf("Auth token: %v", authToken)
}

func TestRevoke(t *testing.T) {
	r := getMyGithub()
	token := gauth.AuthToken{}
	err := r.Revoke(token)
	assert.NotNil(t, err)
	assert.Equal(t, gauth.ErrNotImplemented, err)
}
