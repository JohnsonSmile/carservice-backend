package database

import (
	"carservice/model"
)

func CreateUser(phone string, password string) (*model.User, error) {

	user := &model.User{
		Phone:    phone,
		Username: phone,
		Password: password,
		Avatar:   "",
		Bio:      "",
	}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByPhone(phone string) (*model.User, error) {
	user := &model.User{}
	if err := db.Where("phone = ?", phone).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserById(id uint) (*model.User, error) {
	user := &model.User{}
	if err := db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
