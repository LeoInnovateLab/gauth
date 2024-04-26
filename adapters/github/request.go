package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/LeoInnovateLab/gauth"
	"github.com/LeoInnovateLab/gauth/config"
	"github.com/LeoInnovateLab/gauth/utils"
	"github.com/go-resty/resty/v2"
	"net/url"
)

type AuthRequest struct {
	*gauth.DefaultAuthRequest
}

func (a *AuthRequest) NewAuthRequest(c *config.AuthConfig) (gauth.AuthRequest, error) {
	authRequest := &AuthRequest{}
	authRequest.DefaultAuthRequest = gauth.NewDefaultAuthRequest(c, NewGithub(), authRequest)

	// Ensure that the AuthRequest instance has implemented all methods of the AuthRequest interface
	var _ gauth.AuthRequest = authRequest

	return authRequest, nil
}

func (a *AuthRequest) Authorize(state string) (string, error) {
	url, err := a.DefaultAuthRequest.Authorize(state)
	if err != nil {
		return "", err
	}
	scope := Scope{}
	return utils.UrlBuilderFromBaseUrl(url).
		QueryParam("scope", a.GetScopes(" ", true, scope.GetDefaultScopes())).
		Build(false), nil
}

func (a *AuthRequest) GetAccessToken(callback gauth.AuthCallback) (gauth.AuthToken, error) {
	response, err := a.DefaultAuthRequest.DoPostAuthorizationCode(callback.Code)
	if err != nil {
		return gauth.AuthToken{}, gauth.ErrGetAccessToken
	}

	resMap := utils.ParseStringToMap(response)
	if _, ok := resMap["error"]; ok {
		return toAuthError(resMap["error_description"])
	}

	return gauth.AuthToken{
		AccessToken: resMap["access_token"],
		Scope:       resMap["scope"],
		TokenType:   resMap["token_type"],
	}, nil
}

func toAuthError(urlErr string) (gauth.AuthToken, error) {
	errDesc, err := url.QueryUnescape(urlErr)
	if err != nil {
		return gauth.AuthToken{}, gauth.ErrGetAccessToken
	}
	return gauth.AuthToken{}, errors.New(errDesc)
}

func (a *AuthRequest) GetUserInfo(token gauth.AuthToken) (gauth.AuthUser, error) {
	resp, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+token.AccessToken).
		Get(a.Source.UserInfo())
	if err != nil {
		return gauth.AuthUser{}, gauth.ErrGetUserInfo
	}

	var userMap map[string]interface{}
	json.Unmarshal(resp.Body(), &userMap)

	if _, ok := userMap["error"]; ok {
		return gauth.AuthUser{}, errors.New(userMap["error_description"].(string))
	}

	rawUserInfo := string(resp.Body())

	user := gauth.AuthUser{
		RawUserInfo: rawUserInfo,
		Source:      a.Source.GetSource(),
		Token:       token,
	}

	if val, ok := userMap["id"]; ok && val != nil {
		user.UID = fmt.Sprintf("%d", int(val.(float64)))
	}

	if val, ok := userMap["login"]; ok && val != nil {
		user.Username = val.(string)
	}

	if val, ok := userMap["avatar_url"]; ok && val != nil {
		user.Avatar = val.(string)
	}

	if val, ok := userMap["blog"]; ok && val != nil {
		user.Blog = val.(string)
	}

	if val, ok := userMap["name"]; ok && val != nil {
		user.Nickname = val.(string)
	}

	if val, ok := userMap["location"]; ok && val != nil {
		user.Location = val.(string)
	}

	if val, ok := userMap["email"]; ok && val != nil {
		user.Email = val.(string)
	}

	if val, ok := userMap["bio"]; ok && val != nil {
		user.Remark = val.(string)
	}

	if val, ok := userMap["company"]; ok && val != nil {
		user.Company = val.(string)
	}

	user.Gender = gauth.GenderUnknown

	return user, nil
}
