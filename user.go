package gauth

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
