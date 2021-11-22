package db

import (
	"res-task/config"
	"res-task/model"
	
)


func GetUsers() (*[]model.User, error) {

	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return &[]model.User{}, err
	}
	return &users, nil

}

func GetbyIDUser(id int) (*model.User, error) {
	var user model.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return &model.User{}, err
	}

	return &user, nil
}

func StoreUser(user model.User) (*model.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return &model.User{}, err
	}

	return &user, nil
}

func UpdateUser(id int, user model.User) (*model.User, error) {

	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return &model.User{}, err
	}

	return &user, nil
}

func DeleteUser(id int) (*model.User, error) {
	var user model.User

	if err := config.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return &model.User{}, err
	}

	return &user, nil
}