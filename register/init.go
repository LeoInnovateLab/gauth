package register

import (
	"github.com/LeoInnovateLab/gauth"
	"github.com/LeoInnovateLab/gauth/adapters/facebook"
	"github.com/LeoInnovateLab/gauth/adapters/github"
	"github.com/LeoInnovateLab/gauth/adapters/google"
	"github.com/LeoInnovateLab/gauth/adapters/linkedin"
	"github.com/LeoInnovateLab/gauth/adapters/slack"
)

func init() {
	gauth.Register("github", &github.AuthRequest{})
	gauth.Register("google", &google.AuthRequest{})
	gauth.Register("facebook", &facebook.AuthRequest{})
	gauth.Register("slack", &slack.AuthRequest{})
	gauth.Register("linkedin", &linkedin.AuthRequest{})
}
