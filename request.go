package gauth

type AuthRequest interface {
	Authorize(state string) (string, error)
	Login(callback AuthCallback) (AuthResponse, error)
	GetAccessToken(callback AuthCallback) (AuthToken, error)
	GetUserInfo(token AuthToken) (AuthUser, error)
	Revoke(token AuthToken) error
}
