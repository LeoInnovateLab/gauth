package google

type Google struct{}

func NewGoogle() *Google {
	return &Google{}
}

func (g Google) Authorize() string {
	return "https://accounts.google.com/o/oauth2/v2/auth"
}

func (g Google) AccessToken() string {
	return "https://www.googleapis.com/oauth2/v4/token"
}

func (g Google) UserInfo() string {
	return "https://www.googleapis.com/oauth2/v3/userinfo"
}

func (g Google) GetSource() string {
	return "google"
}
