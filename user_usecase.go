package usecase

import (
	"internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int) (*domain.User, error)
}

type UserUsecase struct {
	UserRepository UserRepository
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.UserRepository.Create(user)
}

func (u *UserUsecase) GetUserByID(id int) (*domain.User, error) {
	return u.UserRepository.GetByID(id)
}
