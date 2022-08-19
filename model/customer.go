package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Customer struct {
	ID     uint
	UserID int
	User   User    `gorm:"foreignKey:UserID"`
	Orders []Order `gorm:"foreignKey:CustomerID;references:ID"`
}

type CustomerService interface {
	Validate() error
	AddCustomerToDB(db *gorm.DB) error
	AddOrder(order Order, db *gorm.DB) error
	DeleteCustomerFromDB(db *gorm.DB) error
	UpdateCustomerInDB(db *gorm.DB) error
}

func (customer *Customer) Validate() error {
	return validation.ValidateStruct(&customer,
		validation.Field(&customer.UserID, validation.Required),
	)
}

func (customer *Customer) AddCustomerToDB(db *gorm.DB) error {
	err := customer.Validate()
	if err != nil {
		return err
	}

	result := db.Create(&customer)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (customer *Customer) AddOrder(order Order, db *gorm.DB) error {
	customer.Orders = append(customer.Orders, order)
	result := db.Save(&customer)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (customer *Customer) DeleteCustomerFromDB(db *gorm.DB) error {
	err := DeleteDeliveryById(db, int(customer.ID))

	if err != nil {
		log.Error(err)
	}

	return err
}

func DeleteCustomerById(db *gorm.DB, id int) error {
	result := db.Where("id = ?", id).Delete(&Customer{})

	if result.Error != nil {
		return result.Error
	} else if db.RowsAffected < 1 {
		log.Error("row with id=%d cannot be deleted because it doesn't exist", id)
	}

	return nil
}

func (oldCustomer *Customer) UpdateCustomerInDB(updatedCustomer *Customer, db *gorm.DB) error {
	err := updatedCustomer.Validate()

	if err != nil {
		return err
	}

	result := db.Model(&oldCustomer).Updates(&updatedCustomer)

	if result.Error != nil {
		return result.Error
	}

	return nil
}