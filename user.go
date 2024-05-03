package gauth

import "strings"

type AuthUser struct {
	UID          string         `json:"uid"`
	Username     string         `json:"username"`
	Nickname     string         `json:"nickname"`
	Avatar       string         `json:"avatar"`
	Blog         string         `json:"blog"`
	Company      string         `json:"company"`
	Location     string         `json:"location"`
	Email        string         `json:"email"`
	Remark       string         `json:"remark"`
	Gender       AuthUserGender `json:"gender"`
	Source       string         `json:"source"`
	Token        AuthToken      `json:"token"`
	RawUserInfo  string         `json:"raw_user_info"`
	SnapshotUser bool           `json:"snapshot_user"`
}

type AuthUserGender int8

const (
	GenderMale    AuthUserGender = 1
	GenderFemale  AuthUserGender = 0
	GenderUnknown AuthUserGender = -1
)

func GetRealGender(originalGender string) AuthUserGender {
	if originalGender == "" || strings.EqualFold(originalGender, "UNKNOWN") {
		return GenderUnknown
	}
	males := map[string]bool{
		"m":    true,
		"男":    true,
		"1":    true,
		"male": true,
	}
	_, ok := males[strings.ToLower(originalGender)]
	if ok {
		return GenderMale
	}
	return GenderFemale
}
