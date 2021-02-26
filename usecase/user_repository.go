package usecase

import "github.com/SHU95/CA-Tech-Dojo-go/domain"

type UserRepository interface {
	FindByID(id int) (domain.User, error)
	Store(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
}
