package postgres

import (
	"DDD-HEX/internal/domain"
	"database/sql"
	"sync"
)

type AuthRepository struct {
	DB *sql.DB
	mu sync.Mutex
}

func (r *AuthRepository) Save(auth domain.AuthEntity) error {
	query := "INSERT INTO auth (user_id, access_token, refresh_token, expires) VALUES ($1, $2, $3, $4)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(auth.UserID, auth.AccessToken, auth.RefreshToken, auth.Expires)
	return err
}

func (r *AuthRepository) FindByAccessToken(token string) (domain.AuthEntity, error) {
	var auth domain.AuthEntity
	query := "SELECT user_id, access_token, refresh_token, expires FROM auth WHERE access_token = $1"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return auth, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(token).Scan(&auth.UserID, &auth.AccessToken, &auth.RefreshToken, &auth.Expires)
	return auth, err
}

func (r *AuthRepository) FindByRefreshToken(token string) (domain.UserAuthEntity, error) {
	var auth domain.UserAuthEntity
	var rule string
	query := "SELECT auth.user_id, auth.access_token, auth.refresh_token, auth.expires, users.email, users.rule, users.status FROM auth JOIN users ON users.id = auth.user_id WHERE refresh_token = $1"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return auth, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(token).Scan(&auth.UserID, &auth.AccessToken, &auth.RefreshToken, &auth.Expires, &auth.Email, &rule, &auth.Status)
	auth.Rule.Scan(rule)
	return auth, err
}
