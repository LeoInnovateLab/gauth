package gauth

type AuthToken struct {
	AccessToken            string `json:"access_token"`
	ExpireIn               int    `json:"expire_in"`
	RefreshToken           string `json:"refresh_token"`
	RefreshTokenExpireIn   int    `json:"refresh_token_expire_in"`
	Uid                    string `json:"uid"`
	OpenId                 string `json:"open_id"`
	AccessCode             string `json:"access_code"`
	UnionId                string `json:"union_id"`
	Scope                  string `json:"scope"`                    // Google
	TokenType              string `json:"token_type"`               // Google
	IdToken                string `json:"id_token"`                 // Google
	MacAlgorithm           string `json:"mac_algorithm"`            // Xiao Mi
	MacKey                 string `json:"mac_key"`                  // Xiao Mi
	Code                   string `json:"code"`                     // Work WeChat
	SnapshotUser           bool   `json:"snapshot_user"`            // WeChat Official Account - Available for web authorization login
	OauthToken             string `json:"oauth_token"`              // Twitter
	OauthTokenSecret       string `json:"oauth_token_secret"`       // Twitter
	UserId                 string `json:"user_id"`                  // Twitter
	ScreenName             string `json:"screen_name"`              // Twitter
	OauthCallbackConfirmed *bool  `json:"oauth_callback_confirmed"` // Twitter
}
