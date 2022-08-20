package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Firstname string
	Lastname  string
	Age       int
	Address   string
	Tel       string
}

type UserService interface {
	Validate() error
	AddUserToDB(db *gorm.DB) error
	DeleteUserFromDB(db *gorm.DB) error
	UpdateUserInDB(db *gorm.DB) error
}

func (user *User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Firstname, validation.Required),
		validation.Field(&user.Lastname, validation.Required),
		validation.Field(&user.Age, is.Int),
	)
}

func (user *User) AddUserToDB(db *gorm.DB) error {
	err := user.Validate()
	if err != nil {
		return err
	}

	result := db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (user *User) DeleteUserFromDB(db *gorm.DB) error {
	err := DeleteUserById(db, int(user.ID))

	if err != nil {
		log.Error(err)
	}

	return err
}

func DeleteUserById(db *gorm.DB, id int) error {
	result := db.Where("id = ?", id).Delete(&User{})

	if result.Error != nil {
		return result.Error
	} else if db.RowsAffected < 1 {
		log.Error("row with id=%d cannot be deleted because it doesn't exist", id)
	}

	return nil
}

func (oldUser *User) UpdateUserInDB(updatedUser *User, db *gorm.DB) error {
	err := updatedUser.Validate()

	if err != nil {
		return err
	}

	result := db.Model(&oldUser).Updates(&updatedUser)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
