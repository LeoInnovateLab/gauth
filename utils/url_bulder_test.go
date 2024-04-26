package utils

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"testing"
)

func TestUrlBuilder_Build(t *testing.T) {
	authorizeUrl := UrlBuilderFromBaseUrl("http://localhost").
		QueryParam("key1", "value1").
		QueryParam("key2", "value2").
		QueryParam("key3", "https://github.com").
		Build(false)

	u, err := url.Parse(authorizeUrl)
	assert.NoError(t, err)
	values := u.Query()
	assert.Equal(t, "value1", values.Get("key1"))
	assert.Equal(t, "value2", values.Get("key2"))
	assert.Equal(t, "https://github.com", values.Get("key3"))

	authorizeUrl = UrlBuilderFromBaseUrl("http://localhost").
		QueryParam("key1", "value1").
		QueryParam("key2", "value2").
		QueryParam("key3", "https://github.com").
		Build(true)

	assert.True(t, strings.Contains(authorizeUrl, "https%3A%2F%2Fgithub.com"))
}
