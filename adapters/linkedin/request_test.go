package linkedin_test

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

	clientID = cmp.Or(os.Getenv("LINKIN_CLIENT_ID"), "client_id")
	clientSecret = cmp.Or(os.Getenv("LINKIN_SECRET"), "secret")
}

func getMyLinkedin() gauth.AuthRequest {
	r, err := gauth.New().
		Source("linkedin").
		ClientId(clientID).
		ClientSecret(clientSecret).
		RedirectUrl("http://localhost:8080/auth/linkedin/callback").
		Build()
	if err != nil {
		log.Fatalf("Failed to build request:%v", err)
	}
	return r
}

func TestLinkedinAuthRequest_Authorize(t *testing.T) {
	r := getMyLinkedin()

	state := utils.CreateState()
	authorizeUrl, err := r.Authorize(state)

	assert.Nil(t, err)

	u, authErr := url.Parse(authorizeUrl)
	assert.NoError(t, authErr)
	assert.Equal(t, "www.linkedin.com", u.Host)
	values := u.Query()
	assert.Equal(t, clientID, values.Get("client_id"))
	assert.Equal(t, state, values.Get("state"))
	assert.Equal(t, "http://localhost:8080/auth/linkedin/callback", values.Get("redirect_uri"))
	assert.Equal(t, "r_liteprofile r_emailaddress w_member_social", values.Get("scope"))
}
