package user

import (
	"users/model"
)

type UserUseCase interface {
	Register(data *model.User) (*model.User, error)
	GetById(id int) (*model.User, error)
	GetAllDataUser()(*[]model.User, error)
	DeleteUserById(id int) error
	UpdateUserById(id int, user *model.User)(*model.User, error)
}
