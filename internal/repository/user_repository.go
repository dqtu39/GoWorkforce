package repository

import (
	"database/sql"
	"github.com/dqtu39/GoWorkforce/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindById(username string) (*domain.User, error)
	ValidateUser(password string, hash string) bool
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	query := "INSERT INTO user (username, password) VALUES (?, ?)"
	_, err = r.db.Exec(query, user.Username, user.Password)
	return err
}

func (r *userRepository) FindById(username string) (*domain.User, error) {
	user := &domain.User{}
	query := "SELECT * FROM user WHERE username = ?"
	row := r.db.QueryRow(query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) ValidateUser(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
