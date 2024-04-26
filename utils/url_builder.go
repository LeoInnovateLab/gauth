package utils

import (
	"fmt"
	"net/url"
	"strings"
)

type UrlBuilder struct {
	baseUrl string
	params  map[string]string
}

func UrlBuilderFromBaseUrl(baseUrl string) *UrlBuilder {
	return &UrlBuilder{
		baseUrl: baseUrl,
		params:  make(map[string]string),
	}
}

func (b *UrlBuilder) QueryParam(key string, value string) *UrlBuilder {
	b.params[key] = value
	return b
}

func (b *UrlBuilder) Build(encode bool) string {
	var paramString string
	if encode {
		params := url.Values{}
		for k, v := range b.params {
			params.Add(k, v)
		}
		paramString = params.Encode()
	} else {
		for k, v := range b.params {
			paramString += fmt.Sprintf("%s=%s&", k, v)
		}
		paramString = strings.TrimSuffix(paramString, "&") // remove trailing '&'
	}

	// Check if baseUrl already contains a '?'
	if strings.Contains(b.baseUrl, "?") {
		// If it does, use '&' to append new parameters
		return fmt.Sprintf("%s&%s", b.baseUrl, paramString)
	} else {
		// If not, use '?' to append new parameters
		return fmt.Sprintf("%s?%s", b.baseUrl, paramString)
	}
}
