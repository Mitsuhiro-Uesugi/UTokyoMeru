package db

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique;not null"`
	Avatar   string `gorm:"type:text"`
	Id       string `gorm:"unique;not null"`
}

type UserGet struct {
	UserName string
	Avatar   string
	Id       string
}

type UserLogin struct {
	Id       string `json:"Id" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type UserCRUD struct{}

func (crud UserCRUD) CreateByObject(u *User) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	if u == nil {
		return errors.New("User not exists")
	}
	result := db.Create(u)
	if result.Error != nil {
		return result.Error
	}

	return result.Error
}

func (crud UserCRUD) GetUserByName(name string) ([]User, error) {
	db, err := GetDatabaseInstance()
	if err != nil {
		return nil, err
	}

	var Users []User
	result := db.Where("name LIKE ?", "%"+name+"%").Find(&Users)
	return Users, result.Error
}

func (crud UserCRUD) UpdateByObject(u *User) error {
	db, err := GetDatabaseInstance()
	if err != nil {
		return err
	}
	result := db.Save(&u)
	return result.Error
}
