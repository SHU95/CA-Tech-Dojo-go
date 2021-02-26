package usecase

import (
	"github.com/SHU95/CA-Tech-Dojo-go/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByID(id)
	return
}

func (interactor *UserInteractor) Add(createUser domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Store(createUser)
	return
}

func (interactor *UserInteractor) Update(updateUser domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Update(updateUser)
	return
}
