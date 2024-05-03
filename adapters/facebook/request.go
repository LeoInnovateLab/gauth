package facebook

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
	authRequest.DefaultAuthRequest = gauth.NewDefaultAuthRequest(c, NewFacebook(), authRequest)

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

	if errorMap, ok := responseMap["error"].(map[string]interface{}); ok {
		if message, ok := errorMap["message"].(string); ok {
			return gauth.AuthToken{}, errors.New(fmt.Sprintf("%s", message))
		}
	}
	return gauth.AuthToken{
		AccessToken: responseMap["access_token"].(string),
		ExpireIn:    int(responseMap["expires_in"].(float64)),
		TokenType:   responseMap["token_type"].(string),
	}, nil
}

func (a *AuthRequest) GetUserInfo(token gauth.AuthToken) (gauth.AuthUser, error) {
	userinfoUrl := utils.UrlBuilderFromBaseUrl(a.Source.UserInfo()).
		QueryParam("access_token", token.AccessToken).
		Build(false)

	resp, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+token.AccessToken).
		Get(userinfoUrl)
	if err != nil {
		return gauth.AuthUser{}, gauth.ErrGetUserInfo
	}

	var userMap map[string]interface{}
	json.Unmarshal(resp.Body(), &userMap)

	rawUserInfo := string(resp.Body())

	if errorMap, ok := userMap["error"].(map[string]interface{}); ok {
		if message, ok := errorMap["message"].(string); ok {
			return gauth.AuthUser{}, errors.New(fmt.Sprintf("%s", message))
		}
	}

	user := gauth.AuthUser{
		RawUserInfo: rawUserInfo,
		Source:      a.Source.GetSource(),
		Token:       token,
	}

	user.UID = userMap["id"].(string)
	user.Nickname = userMap["name"].(string)
	user.Username = userMap["name"].(string)

	avatar := ""
	if picture, ok := userMap["picture"].(map[string]interface{}); ok {
		if data, ok := picture["data"].(map[string]interface{}); ok {
			avatar = data["url"].(string)
		}
	}

	user.Avatar = avatar

	gender := gauth.GenderUnknown
	if g, ok := userMap["gender"]; ok {
		gender = gauth.GetRealGender(g.(string))
	}
	user.Gender = gender

	if email, ok := userMap["email"].(string); ok {
		user.Email = email
	}

	if location, ok := userMap["location"].(string); ok {
		user.Location = location
	}

	if link, ok := userMap["link"].(string); ok {
		user.Blog = link
	}

	return user, nil
}
