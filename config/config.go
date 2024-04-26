package config

import "time"

type AuthConfig struct {
	ClientId       string        // App keys for respective platforms
	ClientSecret   string        // App secrets for respective platforms
	RedirectUri    string        // Callback URL after successful login
	Scopes         []string      // Custom scope content for authorization platforms
	StateCacheTime time.Duration // Cache time for state
}

type Option struct {
	apply func(*AuthConfig)
}

func NewAuthConfig(opts ...Option) *AuthConfig {
	conf := &AuthConfig{
		StateCacheTime: 3 * time.Minute,
	}
	for _, opt := range opts {
		opt.apply(conf)
	}
	return conf
}

func WithScopes(scopes []string) Option {
	return Option{func(c *AuthConfig) {
		c.Scopes = scopes
	}}
}

func WithClientId(clientId string) Option {
	return Option{func(c *AuthConfig) {
		c.ClientId = clientId
	}}
}

func WithClientSecret(clientSecret string) Option {
	return Option{func(c *AuthConfig) {
		c.ClientSecret = clientSecret
	}}
}

func WithRedirectUrl(redirectUrl string) Option {
	return Option{func(c *AuthConfig) {
		c.RedirectUri = redirectUrl
	}}
}
