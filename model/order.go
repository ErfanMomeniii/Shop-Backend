package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderSituation int

const (
	NotStart = iota
	InProgress
	Delivered
)

type Order struct {
	gorm.Model
	Products   []int
	DeliveryID int
	Delivery   Delivery `gorm:"foreignKey:DeliveryID"`
	CustomerID int
	Customer   Customer `gorm:"foreignKey:CustomerID"`
	Situation  OrderSituation
}

type OrderService interface{
	Validate() error
	AddOrderToDB(db *gorm.DB) error
	DeleteOrderFromDB(db *gorm.DB) error
	UpdateOrderInDB(db *gorm.DB) error
}

func (order *Order)Validate()error{
	return validation.ValidateStruct(&order,
		validation.Field(&order.DeliveryID, validation.Required),
		validation.Field(&order.CustomerID, validation.Required),
	)
}


func (order *Order) AddOrderToDB(db *gorm.DB) error {
	err := order.Validate()
	if err != nil {
		return err
	}

	result := db.Create(&order)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (order *Order) DeleteOrderFromDB(db *gorm.DB) error {
	err := DeleteOrderById(db, int(order.ID))

	if err != nil {
		log.Error(err)
	}

	return err
}

func DeleteOrderById(db *gorm.DB, id int) error {
	result := db.Where("id = ?", id).Delete(&Order{})

	if result.Error != nil {
		return result.Error
	} else if db.RowsAffected < 1 {
		log.Error("row with id=%d cannot be deleted because it doesn't exist", id)
	}

	return nil
}

func (oldOrder *Order) UpdateOrderInDB(updatedOrder *User, db *gorm.DB) error {
	err := updatedOrder.Validate()

	if err != nil {
		return err
	}

	result := db.Model(&oldOrder).Updates(&updatedOrder)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
