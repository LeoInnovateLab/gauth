package slack

import (
	"encoding/json"
	"errors"
	"github.com/LeoInnovateLab/gauth"
	"github.com/LeoInnovateLab/gauth/config"
	"github.com/LeoInnovateLab/gauth/utils"
	"github.com/go-resty/resty/v2"
)

type AuthRequest struct {
	*gauth.DefaultAuthRequest
}

func (a *AuthRequest) NewAuthRequest(c *config.AuthConfig) (gauth.AuthRequest, error) {
	authRequest := &AuthRequest{}
	authRequest.DefaultAuthRequest = gauth.NewDefaultAuthRequest(c, NewSlack(), authRequest)

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
	accessTokenUrl := a.DefaultAuthRequest.AccessTokenUrl(callback.Code)

	resp, err := resty.New().R().Post(accessTokenUrl)
	if err != nil {
		return gauth.AuthToken{}, gauth.ErrGetAccessToken
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(resp.Body(), &responseMap)
	if err != nil {
		return gauth.AuthToken{}, gauth.ErrGetAccessToken
	}

	if _, ok := responseMap["error"]; ok {
		return gauth.AuthToken{}, errors.New(responseMap["error"].(string))
	}

	return gauth.AuthToken{
		AccessToken: responseMap["access_token"].(string),
		Scope:       responseMap["scope"].(string),
		TokenType:   responseMap["token_type"].(string),
		Uid:         responseMap["authed_user"].(map[string]interface{})["id"].(string),
	}, nil
}

func (a *AuthRequest) GetUserInfo(token gauth.AuthToken) (gauth.AuthUser, error) {
	userinfoUrl := utils.UrlBuilderFromBaseUrl(a.Source.UserInfo()).
		QueryParam("access_token", token.AccessToken).
		QueryParam("user", token.Uid).
		Build(false)

	resp, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+token.AccessToken).
		Post(userinfoUrl)
	if err != nil {
		return gauth.AuthUser{}, gauth.ErrGetUserInfo
	}

	var resultMap map[string]interface{}
	json.Unmarshal(resp.Body(), &resultMap)
	rawUserInfo := string(resp.Body())

	if _, ok := resultMap["error"]; ok {
		return gauth.AuthUser{}, errors.New(resultMap["error"].(string))
	}

	userMap, ok := resultMap["user"].(map[string]interface{})
	if !ok {
		return gauth.AuthUser{}, gauth.ErrGetUserInfo
	}

	profileMap, ok := userMap["profile"].(map[string]interface{})
	if !ok {
		return gauth.AuthUser{}, gauth.ErrGetUserInfo
	}

	return gauth.AuthUser{
		RawUserInfo: rawUserInfo,
		UID:         userMap["id"].(string),
		Username:    userMap["name"].(string),
		Nickname:    userMap["real_name"].(string),
		Avatar:      profileMap["image_original"].(string),
		Email:       profileMap["email"].(string),
		Gender:      gauth.GenderUnknown,
		Token:       token,
		Source:      a.Source.GetSource(),
	}, nil
}
