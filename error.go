package gauth

import (
	"errors"
)

var (
	ErrAuthFailed           = errors.New("auth failed")
	ErrGetAccessToken       = errors.New("get access token error")
	ErrGetUserInfo          = errors.New("get user info error")
	ErrIllegalStatus        = errors.New("illegal status error")
	ErrNotImplemented       = errors.New("not implemented")
	ErrClientIdNotFound     = errors.New("client id is required")
	ErrClientSecretNotFound = errors.New("client secret is required")
	ErrSourceNotFound       = errors.New("auth source is required")
)
