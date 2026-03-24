package service

import (
	"go_demo/model"
	"go_demo/pkg/db"
)

func CreateUser(username, password, email string) error {
	user := model.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	return db.DB.Create(&user).Error

}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := db.DB.Where("email =?", email).First(&user).Error
	return &user, err
}

func GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	err := db.DB.First(&user, userID).Error
	return &user, err
}
