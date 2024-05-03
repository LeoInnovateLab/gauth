package facebook

type Facebook struct {
}

func NewFacebook() *Facebook {
	return &Facebook{}
}

func (f Facebook) Authorize() string {
	return "https://www.facebook.com/v19.0/dialog/oauth"
}

func (f Facebook) AccessToken() string {
	return "https://graph.facebook.com/v19.0/oauth/access_token"
}

func (f Facebook) UserInfo() string {
	return "https://graph.facebook.com/v19.0/me"
}

func (f Facebook) GetSource() string {
	return "facebook"
}
