package gauth

import (
	"github.com/LeoInnovateLab/gauth/config"
	"github.com/LeoInnovateLab/gauth/utils"
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"log"
	"net/url"
	"strings"
)

type DefaultAuthRequest struct {
	AuthRequest
	config *config.AuthConfig
	Source config.AuthSource
	cache  *cache.Cache
}

func NewDefaultAuthRequest(config *config.AuthConfig, source config.AuthSource,
	authRequest AuthRequest) *DefaultAuthRequest {
	return &DefaultAuthRequest{
		config:      config,
		Source:      source,
		AuthRequest: authRequest,
		cache:       cache.New(config.StateCacheTime, config.StateCacheTime),
	}
}

func (r *DefaultAuthRequest) Authorize(state string) (string, error) {
	url := utils.UrlBuilderFromBaseUrl(r.Source.Authorize()).
		QueryParam("response_type", "code").
		QueryParam("client_id", r.config.ClientId).
		QueryParam("redirect_uri", r.config.RedirectUri).
		QueryParam("state", r.saveCache(state)).
		Build(false)
	return url, nil
}

func (r *DefaultAuthRequest) Login(callback AuthCallback) (AuthResponse, error) {
	authToken, err := r.GetAccessToken(callback)
	if err != nil {
		return AuthResponse{}, err
	}

	err = r.checkState(callback.State)
	if err != nil {
		return AuthResponse{}, err
	}

	userInfo, authError := r.GetUserInfo(authToken)
	if authError != nil {
		return AuthResponse{}, authError
	}
	return AuthResponse{
		Data: userInfo,
	}, nil
}

func (r *DefaultAuthRequest) checkState(state string) error {
	_, found := r.cache.Get(state)
	if !found {
		return ErrIllegalStatus
	}
	return nil
}

func (r *DefaultAuthRequest) GetScopes(separator string, encode bool, defaultScopes []string) string {
	scopes := r.config.Scopes
	if scopes == nil || len(scopes) == 0 {
		if defaultScopes == nil || len(defaultScopes) == 0 {
			return ""
		}
		scopes = defaultScopes
	}
	if separator == "" {
		// default separator is space
		separator = " "
	}
	scopeStr := strings.Join(scopes, separator)
	if encode {
		return url.QueryEscape(scopeStr)
	}
	return scopeStr
}

func (r *DefaultAuthRequest) DoPostAuthorizationCode(code string) (string, error) {
	accessTokenUrl := r.AccessTokenUrl(code)

	resp, err := resty.New().R().Post(accessTokenUrl)
	if err != nil {
		log.Printf("An error occurred while getting access token: %v", err)
		return "", err
	}
	body := resp.Body()

	return string(body), nil
}

func (r *DefaultAuthRequest) AccessTokenUrl(code string) string {
	return utils.UrlBuilderFromBaseUrl(r.Source.AccessToken()).
		QueryParam("code", code).
		QueryParam("client_id", r.config.ClientId).
		QueryParam("client_secret", r.config.ClientSecret).
		QueryParam("grant_type", "authorization_code").
		QueryParam("redirect_uri", r.config.RedirectUri).
		Build(false)
}

func (r *DefaultAuthRequest) saveCache(state string) string {
	r.cache.Set(state, state, cache.DefaultExpiration)
	return state
}

func (r *DefaultAuthRequest) Revoke(token AuthToken) error {
	return ErrNotImplemented
}
