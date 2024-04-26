package gauth

type AuthCallback struct {
	Code              string `json:"code" schema:"code"`
	AuthCode          string `json:"auth_code" schema:"auth_code"`
	State             string `json:"state" schema:"state"`
	AuthorizationCode string `json:"authorization_code" schema:"authorization_code"`
	OauthToken        string `json:"oauth_token" schema:"oauth_token"`
	OauthVerifier     string `json:"oauth_verifier" schema:"oauth_verifier"`
}
