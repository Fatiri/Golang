package repo

import (
	"fmt"
	"users/model"
	"users/user"

	"github.com/jinzhu/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func (u UserRepoImpl) UpdateUserById(id int, data *model.User) (*model.User, error) {
	user := model.User{}
	err := u.db.Model(&user).Where("id = ?", id).Update(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[UserRepoImpl.Update] Got Error when update by id %v\n", err)
	}
	return data, nil
}

func (u UserRepoImpl) DeleteUserById(id int) error {
	data := model.User{}
	err := u.db.Where("id = ?", id).Delete(&data).Error
	if err != nil {
		return fmt.Errorf("[UserRepoImpl] got error when delete user by id %v\n", err)
	}
	return nil
}

func (u UserRepoImpl) GetAllDataUser() (*[]model.User, error) {
	var dataUser []model.User
	//Inget harus di kasih pointer
	err := u.db.Find(&dataUser).Error
	if err != nil {
		return nil, fmt.Errorf("[UserRepoImpl.getAllDataUser] got error when get all data user %v\n", err)
	}
	return &dataUser, nil
}

func (u UserRepoImpl) GetById(id int) (*model.User, error) {
	data := model.User{}
	//Menerima error saat melakukan query .Error
	err := u.db.First(&data, id).Error
	if err != nil {
		return nil, fmt.Errorf("[UserRepoImpl.GetById] Get error when get user by id %w\n", err)
	}
	return &data, nil
}

func (u UserRepoImpl) Register(data *model.User) (*model.User, error) {
	err := u.db.Save(&data).Error
	if err != nil {
		return nil, fmt.Errorf("Got Error on Implement Register %v\n", err)
	}
	return data, nil
}

func CreateUserRepoImpl(db *gorm.DB) user.UserRepo {
	return &UserRepoImpl{db}
}
