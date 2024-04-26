package config

type AuthSource interface {
	Authorize() string
	AccessToken() string
	UserInfo() string
	GetSource() string
}
