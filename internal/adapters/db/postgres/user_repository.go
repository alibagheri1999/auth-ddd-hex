package postgres

import (
	"DDD-HEX/internal/domain"
	"database/sql"
	"sync"
)

type UserRepository struct {
	DB *sql.DB
	mu sync.Mutex
}

func (r *UserRepository) Save(user domain.User) error {
	stmt, err := r.DB.Prepare("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password)
	return err
}

func (r *UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	stmt, err := r.DB.Prepare("SELECT id, name, email, password FROM users WHERE email=$1")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}

func (r *UserRepository) FindByID(id string) (domain.User, error) {
	var user domain.User
	stmt, err := r.DB.Prepare("SELECT id, name, email, password FROM users WHERE id=$1")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}
