package slack

type Slack struct{}

func NewGoogle() *Slack {
	return &Slack{}
}

func (s Slack) Authorize() string {
	return "https://slack.com/oauth/v2/authorize"
}

func (s Slack) AccessToken() string {
	return "https://slack.com/api/oauth.v2.access"
}

func (s Slack) UserInfo() string {
	return "https://slack.com/api/users.info"
}

func (s Slack) GetSource() string {
	return "slack"
}
