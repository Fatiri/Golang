package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
	//gorm unique_index itu
	Email string `gorm:"unique_index"json:"email"`
}

//Untuk mengkustomisasi nama table
//Default name functionya harus TableName
func (User) TableName() string {
	return "tb_user"
}
