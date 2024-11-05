package usecase

import (
	"errors"
	"github.com/dqtu39/GoWorkforce/internal/domain"
	"github.com/dqtu39/GoWorkforce/internal/repository"
	"github.com/dqtu39/GoWorkforce/pkg/auth"
)

type UserUseCase interface {
	Register(user *domain.User) error
	Login(username string, password string) (string, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo}
}

func (u *userUseCase) Register(user *domain.User) error {
	_, err := u.repo.FindById(user.Username)
	if err != nil {
		return err
	}
	return u.repo.Create(user)
}

func (u *userUseCase) Login(username string, password string) (string, error) {
	user, err := u.repo.FindById(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if !u.repo.ValidateUser(password, user.Password) {
		return "", errors.New("invalid credentials")
	}
	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
