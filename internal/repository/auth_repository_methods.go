package repository

import (
	"database/sql"
	"fmt"
	"forum/models"
	"log"
	"time"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) CreateUser(user models.User) error {
	query := "INSERT INTO user (email, username, password) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, user.Email, user.Username, user.Password)
	if err != nil {
		log.Printf("repo: create user %s", err)
		return fmt.Errorf("repository create user: %w", err)
	}
	return nil
}

func (r *AuthRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := "SELECT username, password FROM user where email=$1"
	row := r.db.QueryRow(query, email)
	if err := row.Scan(&user.Username, &user.Password); err != nil {
		return user, fmt.Errorf("repository get user by email: %w", err)
	}
	return user, nil
}

func (r *AuthRepo) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	query := "SELECT email, password FROM user where username=$1"
	row := r.db.QueryRow(query, username)
	if err := row.Scan(&user.Email, &user.Password); err != nil {
		return user, fmt.Errorf("repository get user by username: %w", err)
	}
	return user, nil
}

func (r *AuthRepo) GetUserByToken(token string) (models.User, error) {
	var user models.User
	query := "SELECT * from user where token=$1"
	row := r.db.QueryRow(query, token)
	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Token, &user.TokenDuration); err != nil {
		return user, fmt.Errorf("repository get user by token: %w", err)
	}
	return user, nil
}

func (r *AuthRepo) SaveToken(username, token string, duration time.Time) error {
	query := "UPDATE user set token=$1, token_duration=$2 where username=$3"
	_, err := r.db.Exec(query, token, duration, username)
	if err != nil {
		return fmt.Errorf("repository save token: %w", err)
	}
	return nil
}

func (r *AuthRepo) DeleteToken(token string) error {
	query := `UPDATE user set token=NULL, SET token_duration=NULL WHERE token=$1`
	_, err := r.db.Exec(query, token)
	if err != nil {
		return fmt.Errorf("repository delete token: %w", err)
	}
	return nil
}
