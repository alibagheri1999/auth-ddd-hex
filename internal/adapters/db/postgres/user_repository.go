package postgres

import (
	"DDD-HEX/internal/domain"
	"DDD-HEX/internal/ports/clients"
	"DDD-HEX/internal/ports/repository"
	"context"
	"time"
)

// UserRepository handles user-related database operations.
type UserRepository struct {
	DB clients.Database
}

var _ repository.UserRepository = (*UserRepository)(nil)

// Save inserts a new user into the database.
func (r *UserRepository) Save(ctx context.Context, user domain.UserEntity) error {
	query := `
		INSERT INTO users 
		(name, email, password, two_fa_activated, phone_number, rule, status, date_created)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.DB.Exec(ctx, query, user.Name, user.Email, user.Password, false, user.PhoneNumber, domain.UserRole(1), "activated", time.Now())
	return err
}

// FindByEmail retrieves a user by email.
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.UserEntity, error) {
	query := `
		SELECT id, name, email, password, two_fa_activated, phone_number, rule, status 
		FROM users 
		WHERE email=$1`
	row := r.DB.QueryRow(ctx, query, email)

	var user domain.UserEntity
	var rule string
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.TowFAActivated, &user.PhoneNumber, &rule, &user.Status)
	if err != nil {
		return user, err
	}

	user.Rule.Scan(rule)
	return user, nil
}

// FindByID retrieves a user by ID.
func (r *UserRepository) FindByID(ctx context.Context, id string) (domain.UserEntity, error) {
	query := `
		SELECT id, name, email, password, two_fa_activated, phone_number, rule, status 
		FROM users 
		WHERE id=$1`
	row := r.DB.QueryRow(ctx, query, id)

	var user domain.UserEntity
	var rule string
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.TowFAActivated, &user.PhoneNumber, &rule, &user.Status)
	if err != nil {
		return user, err
	}

	user.Rule.Scan(rule)
	return user, nil
}
