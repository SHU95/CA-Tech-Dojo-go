package database

import (
	"github.com/SHU95/CA-Tech-Dojo-go/domain"
)

type UserRepository struct {
	SqlHandler
}

func (userRepo *UserRepository) FindByID(id int) (user domain.User, err error) {
	err = userRepo.Find(&user, id).Error
	if err != nil {
		return
	}
	return
}

//mysqlに作ったユーザーを格納
func (userRepo *UserRepository) Store(storeUser domain.User) (user domain.User, err error) {
	err = userRepo.Create(&storeUser).Error
	if err != nil {
		return
	}
	user = storeUser
	return
}

func (userRepo *UserRepository) Update(updateUser domain.User) (user domain.User, err error) {
	err = userRepo.Save(&updateUser).Error
	if err != nil {
		return
	}
	user = updateUser
	return
}
