package linkedin

type Linkedin struct{}

func NewLinkedin() *Linkedin {
	return &Linkedin{}
}

func (s Linkedin) Authorize() string {
	return "https://www.linkedin.com/oauth/v2/authorization"
}

func (s Linkedin) AccessToken() string {
	return "https://www.linkedin.com/oauth/v2/accessToken"
}

func (s Linkedin) UserInfo() string {
	return "https://api.linkedin.com/v2/me"
}

func (s Linkedin) GetSource() string {
	return "linkedin"
}
