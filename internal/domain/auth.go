package domain

type Auth struct {
	UserID       string `"json:"user_id""`
	AccessToken  string `"json:"access_token""`
	RefreshToken string `"json:"refresh_token""`
	Expires      int64  `"json:"expires""`
}
