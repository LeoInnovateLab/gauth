package utils

import (
	"github.com/google/uuid"
	"strings"
)

func CreateState() string {
	return uuid.New().String()
}

func ParseStringToMap(str string) map[string]string {
	m := make(map[string]string)
	if str == "" {
		return m
	}
	for _, kv := range strings.Split(str, "&") {
		kvPair := strings.Split(kv, "=")
		m[kvPair[0]] = kvPair[1]
	}
	return m
}
