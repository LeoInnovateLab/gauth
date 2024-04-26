package github

type Github struct{}

func NewGithub() *Github {
	return &Github{}
}

func (g Github) Authorize() string {
	return "https://github.com/login/oauth/authorize"
}

func (g Github) AccessToken() string {
	return "https://github.com/login/oauth/access_token"
}

func (g Github) UserInfo() string {
	return "https://api.github.com/user"
}

func (g Github) GetSource() string {
	return "github"
}
