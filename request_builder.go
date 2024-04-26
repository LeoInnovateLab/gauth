package gauth

import (
	"errors"
	"github.com/LeoInnovateLab/gauth/config"
)

type AuthRequestFactory interface {
	NewAuthRequest(*config.AuthConfig) (AuthRequest, error)
}

var factories = make(map[string]AuthRequestFactory)

func Register(source string, factory AuthRequestFactory) {
	factories[source] = factory
}

type AuthRequestBuilder struct {
	source       string
	clientId     string
	clientSecret string
	redirectUrl  string
}

func New() *AuthRequestBuilder {
	return &AuthRequestBuilder{}
}

func (b *AuthRequestBuilder) Source(source string) *AuthRequestBuilder {
	b.source = source
	return b
}

func (b *AuthRequestBuilder) ClientId(clientId string) *AuthRequestBuilder {
	b.clientId = clientId
	return b
}

func (b *AuthRequestBuilder) ClientSecret(clientSecret string) *AuthRequestBuilder {
	b.clientSecret = clientSecret
	return b
}

func (b *AuthRequestBuilder) RedirectUrl(redirectUrl string) *AuthRequestBuilder {
	b.redirectUrl = redirectUrl
	return b
}

func (b *AuthRequestBuilder) Build() (AuthRequest, error) {
	if b.clientId == "" {
		return nil, ErrClientIdNotFound
	}
	if b.source == "" {
		return nil, ErrSourceNotFound
	}
	if b.clientSecret == "" {
		return nil, ErrClientSecretNotFound
	}

	factory, ok := factories[b.source]
	if !ok {
		return nil, errors.New("source not supported yet")
	}
	return factory.NewAuthRequest(config.NewAuthConfig(
		config.WithClientId(b.clientId),
		config.WithClientSecret(b.clientSecret),
		config.WithRedirectUrl(b.redirectUrl),
	))
}
