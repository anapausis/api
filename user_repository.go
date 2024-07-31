package database

import (
	"database/sql"
	"internal/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(user *domain.User) error {
	_, err := r.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	return err
}

func (r *UserRepository) GetByID(id int) (*domain.User, error) {
	row := r.DB.QueryRow("SELECT id, name, email, password FROM users WHERE id = ?", id)
	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
