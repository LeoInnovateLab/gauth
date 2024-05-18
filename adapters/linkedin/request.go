package linkedin

import (
	"encoding/json"
	"errors"
	"github.com/LeoInnovateLab/gauth"
	"github.com/LeoInnovateLab/gauth/config"
	"github.com/LeoInnovateLab/gauth/utils"
	"github.com/go-resty/resty/v2"
	"log"
)

type AuthRequest struct {
	*gauth.DefaultAuthRequest
}

func (a *AuthRequest) NewAuthRequest(c *config.AuthConfig) (gauth.AuthRequest, error) {
	authRequest := &AuthRequest{}
	authRequest.DefaultAuthRequest = gauth.NewDefaultAuthRequest(c, NewLinkedin(), authRequest)

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

	resp, err := resty.New().R().
		SetHeader("Host", "www.linkedin.com").
		Post(accessTokenUrl)
	if err != nil {
		return gauth.AuthToken{}, gauth.ErrGetAccessToken
	}

	var responseMap map[string]interface{}
	err = json.Unmarshal(resp.Body(), &responseMap)
	if err != nil {
		return gauth.AuthToken{}, gauth.ErrGetAccessToken
	}

	if _, ok := responseMap["error"]; ok {
		return gauth.AuthToken{}, errors.New(responseMap["error_description"].(string))
	}

	return gauth.AuthToken{
		AccessToken:  responseMap["access_token"].(string),
		ExpireIn:     int(responseMap["expires_in"].(float64)),
		RefreshToken: responseMap["refresh_token"].(string),
	}, nil
}

func (a *AuthRequest) GetUserInfo(token gauth.AuthToken) (gauth.AuthUser, error) {
	userinfoUrl := utils.UrlBuilderFromBaseUrl(a.Source.UserInfo()).
		QueryParam("projection", "(id,firstName,lastName,profilePicture(displayImage~:playableStreams))").
		Build(false)

	resp, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+token.AccessToken).
		SetHeader("Host", "api.linkedin.com").
		SetHeader("Connection", "Keep-Alive").
		Post(userinfoUrl)
	if err != nil {
		return gauth.AuthUser{}, gauth.ErrGetUserInfo
	}

	var resultMap map[string]interface{}
	json.Unmarshal(resp.Body(), &resultMap)
	rawUserInfo := string(resp.Body())

	if _, ok := resultMap["error"]; ok {
		return gauth.AuthUser{}, errors.New(resultMap["error_description"].(string))
	}

	userName := getUserName(resultMap)
	avatar := getAvatar(resultMap)
	email, err := getEmail(token.AccessToken)
	if err != nil {
		log.Printf("Error getting email: %v", err)
	}

	return gauth.AuthUser{
		RawUserInfo: rawUserInfo,
		UID:         resultMap["id"].(string),
		Username:    userName,
		Nickname:    userName,
		Avatar:      avatar,
		Email:       email,
		Gender:      gauth.GenderUnknown,
		Token:       token,
		Source:      a.Source.GetSource(),
	}, nil
}

func getEmail(accessToken string) (string, error) {
	resp, err := resty.New().R().
		SetHeader("Authorization", "Bearer "+accessToken).
		SetHeader("Host", "api.linkedin.com").
		SetHeader("Connection", "Keep-Alive").
		Get("https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))")

	if err != nil {
		return "", gauth.ErrGetUserInfo
	}

	var resultMap map[string]interface{}
	json.Unmarshal(resp.Body(), &resultMap)

	if _, ok := resultMap["error"]; ok {
		return "", errors.New(resultMap["error_description"].(string))
	}

	elements, ok := resultMap["elements"].([]interface{})
	if !ok || len(elements) == 0 {
		return "", gauth.ErrGetUserInfo
	}

	handleObj, ok := elements[0].(map[string]interface{})
	if !ok {
		return "", gauth.ErrGetUserInfo
	}

	emailAddress, ok := handleObj["emailAddress"].(string)
	if !ok {
		return "", gauth.ErrGetUserInfo
	}

	return emailAddress, nil
}

func getAvatar(resultMap map[string]interface{}) string {
	profilePictureObj, ok := resultMap["profilePicture"].(map[string]interface{})
	if !ok {
		return ""
	}

	displayImageObj, ok := profilePictureObj["displayImage~"].(map[string]interface{})
	if !ok {
		return ""
	}

	elements, ok := displayImageObj["elements"].([]interface{})
	if !ok || len(elements) == 0 {
		return ""
	}

	largestImageObj, ok := elements[len(elements)-1].(map[string]interface{})
	if !ok {
		return ""
	}

	identifiers, ok := largestImageObj["identifiers"].([]interface{})
	if !ok || len(identifiers) == 0 {
		return ""
	}

	identifierObj, ok := identifiers[0].(map[string]interface{})
	if !ok {
		return ""
	}

	identifier, ok := identifierObj["identifier"].(string)
	if !ok {
		return ""
	}

	return identifier
}

func getUserName(resultMap map[string]interface{}) string {
	firstName, ok := resultMap["localizedFirstName"].(string)
	if !ok {
		firstName = getUserNameByKey(resultMap, "firstName")
	}

	lastName, ok := resultMap["localizedLastName"].(string)
	if !ok {
		lastName = getUserNameByKey(resultMap, "lastName")
	}

	return firstName + " " + lastName
}

func getUserNameByKey(resultMap map[string]interface{}, nameKey string) string {
	nameObj, ok := resultMap[nameKey].(map[string]interface{})
	if !ok {
		return ""
	}

	localizedObj, ok := nameObj["localized"].(map[string]interface{})
	if !ok {
		return ""
	}

	preferredLocaleObj, ok := nameObj["preferredLocale"].(map[string]interface{})
	if !ok {
		return ""
	}

	language, ok := preferredLocaleObj["language"].(string)
	if !ok {
		return ""
	}

	country, ok := preferredLocaleObj["country"].(string)
	if !ok {
		return ""
	}

	name, ok := localizedObj[language+"_"+country].(string)
	if !ok {
		return ""
	}

	return name
}
