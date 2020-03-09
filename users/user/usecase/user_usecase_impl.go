package usecase

import (
	"users/model"
	"users/user"
)

type UserUserCaseImpl struct {
	userRepo user.UserRepo
}

func (u UserUserCaseImpl) UpdateUserById(id int, data *model.User) (*model.User, error) {
	return u.userRepo.UpdateUserById(id, data)
}

func (u UserUserCaseImpl) DeleteUserById(id int) error {
	return u.userRepo.DeleteUserById(id)
}

func (u UserUserCaseImpl) GetAllDataUser() (*[]model.User, error) {
	return u.userRepo.GetAllDataUser()
}

func (u UserUserCaseImpl) GetById(id int) (*model.User, error) {
	return u.userRepo.GetById(id)
}

func (u UserUserCaseImpl) Register(data *model.User) (*model.User, error) {
	return u.userRepo.Register(data)
}

func CreateUserUseCaseImpl(userRepo user.UserRepo) user.UserRepo {
	return &UserUserCaseImpl{userRepo}
}
