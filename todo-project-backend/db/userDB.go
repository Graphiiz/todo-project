package db

import (
	"gorm.io/gorm"
)

type UserDB struct {
	gorm.Model
	// Code  string
	Name  string
	Email string
	Todos []byte
}

func ReadEmail(email string) (*UserDB, error) {
	db := DB()

	var userDB UserDB
	if err := db.First(&userDB, "Email = ?", email).Error; err != nil {
		return nil, err
	}

	return &userDB, nil

}

func Create(user UserDB) error {
	db := DB()
	userData := UserDB{
		// Code:  user.Code,
		Name:  user.Name,
		Email: user.Email,
		Todos: user.Todos,
	}

	if err := db.Create(&userData).Error; err != nil {
		return err
	}

	return nil
}

func Save(user UserDB) error {
	db := DB()
	var userDB UserDB
	if err := db.First(&userDB, "Email = ?", user.Email).Error; err != nil {
		return err
	}

	userDB.Todos = user.Todos

	if err := db.Save(&userDB).Error; err != nil {
		return err
	}

	return nil
}
