package utils_test

import (
	"github.com/LeoInnovateLab/gauth/adapters/github"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDefaultScopes(t *testing.T) {

	githubScope := github.Scope{}
	scopes := githubScope.GetDefaultScopes()
	assert.Equal(t, 0, len(scopes))
}
