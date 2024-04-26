package register

import (
	"github.com/LeoInnovateLab/gauth"
	"github.com/LeoInnovateLab/gauth/adapters/github"
	"github.com/LeoInnovateLab/gauth/adapters/google"
)

func init() {
	gauth.Register("github", &github.AuthRequest{})
	gauth.Register("google", &google.AuthRequest{})
}
