package google

import (
	"encoding/json"
	"errors"
	"fmt"
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
	authRequest.DefaultAuthRequest = gauth.NewDefaultAuthRequest(c, NewGoogle(), authRequest)

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

	var responseMap map[string]interface{}
	err = json.Unmarshal([]byte(response), &responseMap)
	if err != nil {
		return gauth.AuthToken{}, gauth.ErrGetAccessToken
	}

	if _, ok := responseMap["error"]; ok {
		return gauth.AuthToken{},
			errors.New(fmt.Sprintf("%s:%s",
				responseMap["error"].(string),
				responseMap["error_description"].(string)))
	}
	return gauth.AuthToken{
		AccessToken: responseMap["access_token"].(string),
		ExpireIn:    int(responseMap["expires_in"].(float64)),
		Scope:       responseMap["scope"].(string),
		TokenType:   responseMap["token_type"].(string),
		IdToken:     responseMap["id_token"].(string),
	}, nil
}

func (a *AuthRequest) GetUserInfo(token gauth.AuthToken) (gauth.AuthUser, error) {
	userinfoUrl := utils.UrlBuilderFromBaseUrl(a.Source.UserInfo()).
		QueryParam("access_token", token.AccessToken).
		Build(false)

	resp, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+token.AccessToken).
		Post(userinfoUrl)
	if err != nil {
		return gauth.AuthUser{}, gauth.ErrGetUserInfo
	}

	var userMap map[string]interface{}
	json.Unmarshal(resp.Body(), &userMap)

	rawUserInfo := string(resp.Body())

	if _, ok := userMap["error"]; ok {
		return gauth.AuthUser{},
			errors.New(fmt.Sprintf("%s:%s",
				userMap["error"].(string),
				userMap["error_description"].(string)))
	}

	return gauth.AuthUser{
		RawUserInfo: rawUserInfo,
		UID:         userMap["sub"].(string),
		Email:       userMap["email"].(string),
		Username:    userMap["email"].(string),
		Avatar:      userMap["picture"].(string),
		Location:    userMap["locale"].(string),
		Nickname:    userMap["name"].(string),
		Token:       token,
		Gender:      gauth.GenderUnknown,
		Source:      a.Source.GetSource(),
	}, nil
}
