package postgres

import (
	"DDD-HEX/internal/domain"
	"database/sql"
	"sync"
	"time"
)

type UserRepository struct {
	DB *sql.DB
	mu sync.Mutex
}

func (r *UserRepository) Save(user domain.UserEntity) error {
	stmt, err := r.DB.Prepare("INSERT INTO users (name, email, password, two_fa_activated ,phone_number ,rule ,status ,date_created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email, user.Password, false, user.PhoneNumber, domain.UserRole(1), "activated", time.Now())
	return err
}

func (r *UserRepository) FindByEmail(email string) (domain.UserEntity, error) {
	var user domain.UserEntity
	var rule string
	stmt, err := r.DB.Prepare("SELECT id, name, email, password, two_fa_activated ,phone_number ,rule ,status FROM users WHERE email=$1")
	if err != nil {
		return user, err
	}

	defer stmt.Close()
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.TowFAActivated, &user.PhoneNumber, &rule, &user.Status)
	user.Rule.Scan(rule)
	return user, err
}

func (r *UserRepository) FindByID(id string) (domain.UserEntity, error) {
	var user domain.UserEntity
	var rule string
	stmt, err := r.DB.Prepare("SELECT id, name, email, password, two_fa_activated ,phone_number ,rule ,status FROM users WHERE id=$1")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.TowFAActivated, &user.PhoneNumber, &rule, &user.Status)
	user.Rule.Scan(rule)
	return user, err
}
