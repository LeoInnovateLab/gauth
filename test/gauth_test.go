package main

import (
	"github.com/LeoInnovateLab/gauth"
	"github.com/LeoInnovateLab/gauth/adapters/github"
	_ "github.com/LeoInnovateLab/gauth/register"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestBuilderRequest(t *testing.T) {
	authRequest, err := gauth.New().
		Source("github").
		ClientId("gitlabClientID").
		ClientSecret("gitlabSecret").
		RedirectUrl("http://localhost:8080/auth/github/callback").
		Build()
	if err != nil {
		log.Fatalf("Failed to build request:%v", err)
	}

	assert.NotNil(t, authRequest)
	assert.IsType(t, &github.AuthRequest{}, authRequest)
}
