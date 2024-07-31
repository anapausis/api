package usecase

import (
	"internal/domain"
)

type WorkRepository interface {
	Create(work *domain.Works) error
	GetByID(id int) (*domain.Works, error)
}

type WorkUsecase struct {
	WorkRepository WorkRepository
}

func (u *WorkUsecase) Creatework(work *domain.Works) error {
	return u.WorkRepository.Create(work)
}

func (u *WorkUsecase) GetworkByID(id int) (*domain.Works, error) {
	return u.WorkRepository.GetByID(id)
}
