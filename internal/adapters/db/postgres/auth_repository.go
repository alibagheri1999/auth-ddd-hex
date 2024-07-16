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

func (r *AuthRepository) Save(auth domain.Auth) error {
	query := "INSERT INTO auth (user_id, access_token, refresh_token, expires) VALUES ($1, $2, $3, $4)"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(auth.UserID, auth.AccessToken, auth.RefreshToken, auth.Expires)
	return err
}

func (r *AuthRepository) FindByAccessToken(token string) (domain.Auth, error) {
	var auth domain.Auth
	query := "SELECT user_id, access_token, refresh_token, expires FROM auth WHERE access_token = $1"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return auth, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(token).Scan(&auth.UserID, &auth.AccessToken, &auth.RefreshToken, &auth.Expires)
	return auth, err
}

func (r *AuthRepository) FindByRefreshToken(token string) (domain.Auth, error) {
	var auth domain.Auth
	query := "SELECT user_id, access_token, refresh_token, expires FROM auth WHERE refresh_token = $1"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return auth, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(token).Scan(&auth.UserID, &auth.AccessToken, &auth.RefreshToken, &auth.Expires)
	return auth, err
}
