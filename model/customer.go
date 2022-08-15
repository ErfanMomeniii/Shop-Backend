package model

import "gorm.io/gorm"

type Customer struct {
	ID     uint
	UserID int
	User   User `gorm:"foreignKey:UserID"`
}

type CustomerService interface {
	Validate() error
	AddCustomerToDB(db *gorm.DB) error
	DeleteCustomerFromDB(db *gorm.DB) error
	UpdateCustomerInDB(db *gorm.DB) error
}
