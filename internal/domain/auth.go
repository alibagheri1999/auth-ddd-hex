package domain

type AuthEntity struct {
	ID           int    `"db:"id""`
	UserID       int    `"db:"user_id""`
	AccessToken  string `"db:"access_token""`
	RefreshToken string `"db:"refresh_token""`
	Expires      int64  `"db:"expires""`
}

type UserAuthEntity struct {
	UserID       int      `"db:"user_id""`
	AccessToken  string   `"db:"access_token""`
	RefreshToken string   `"db:"refresh_token""`
	Expires      int64    `"db:"expires""`
	Email        string   `"db:"email""`
	Rule         UserRole `"db:"rule""`
	Status       string   `"db:"status""`
}
