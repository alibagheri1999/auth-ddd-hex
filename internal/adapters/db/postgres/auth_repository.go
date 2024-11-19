package postgres

import (
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/ports/clients"
	"DDD-HEX/internal/ports/repository"
	"context"
)

// AuthRepository handles auth-related database operations.
type AuthRepository struct {
	DB clients.Database // Use the Database interface
}

var _ repository.AuthRepository = (*AuthRepository)(nil)

// Save inserts a new auth record into the database.
func (r *AuthRepository) Save(ctx context.Context, auth domain.AuthEntity) error {
	query := "INSERT INTO auth (user_id, access_token, refresh_token, expires) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(ctx, query, auth.UserID, auth.AccessToken, auth.RefreshToken, auth.Expires)
	return err
}

// FindByAccessToken retrieves an auth record by its access token.
func (r *AuthRepository) FindByAccessToken(ctx context.Context, token string) (domain.AuthEntity, error) {
	query := "SELECT user_id, access_token, refresh_token, expires FROM auth WHERE access_token = $1"
	row := r.DB.QueryRow(ctx, query, token)

	var auth domain.AuthEntity
	err := row.Scan(&auth.UserID, &auth.AccessToken, &auth.RefreshToken, &auth.Expires)
	return auth, err
}

// FindByRefreshToken retrieves a user and auth record by its refresh token.
func (r *AuthRepository) FindByRefreshToken(ctx context.Context, token string) (domain.UserAuthEntity, error) {
	query := `
		SELECT auth.user_id, auth.access_token, auth.refresh_token, auth.expires, 
		       users.email, users.rule, users.status 
		FROM auth 
		JOIN users ON users.id = auth.user_id 
		WHERE refresh_token = $1`
	row := r.DB.QueryRow(ctx, query, token)

	var auth domain.UserAuthEntity
	var rule string
	err := row.Scan(&auth.UserID, &auth.AccessToken, &auth.RefreshToken, &auth.Expires, &auth.Email, &rule, &auth.Status)
	auth.Rule.Scan(rule)
	return auth, err
}
