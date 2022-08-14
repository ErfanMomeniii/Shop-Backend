package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Delivery struct {
	Id           int
	User_id      int
	Score        float64
	DeliveryRate int
}

type DeliveryServie interface {
	validateDelivery() error
	CreateDelivery(db *gorm.DB) error
	DeleteDelivery(db *gorm.DB) error
	UpdateDelivery(updatedDelivery *Delivery, db *gorm.DB) error
	ChangeDeliveryRate(newDeliveryRate int, db *gorm.DB) error
	ChangeScore(newScore float64, db *gorm.DB) error
}

func (delivery *Delivery) validateDelivery() error {
	return validation.ValidateStruct(&delivery,
		validation.Field(&delivery.User_id, validation.Required),
		validation.Field(&delivery.Score, is.Float),
		validation.Field(&delivery.DeliveryRate, is.Int),
	)
}

func (delivery *Delivery) CreateDelivery(db *gorm.DB) error {
	err := delivery.validateDelivery()
	if err != nil {
		return err
	}

	result := db.Create(&delivery)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (delivery *Delivery) DeleteDelivery(db *gorm.DB) error {
	err := DeleteDeliveryById(db, delivery.Id)

	if err != nil {
		log.Error(err)
	}

	return err
}

func DeleteDeliveryById(db *gorm.DB, id int) error {
	result := db.Where("id = ?", id).Delete(&Delivery{})

	if result.Error != nil {
		return result.Error
	} else if db.RowsAffected < 1 {
		log.Error("row with id=%d cannot be deleted because it doesn't exist", id)
	}

	return nil
}

func (oldDelivery *Delivery) UpdateDelivery(updatedDelivery *Delivery, db *gorm.DB) error {
	err := updatedDelivery.validateDelivery()

	if err != nil {
		return err
	}

	result := db.Model(&oldDelivery).Updates(&updatedDelivery)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (delivery *Delivery) ChangeDeliveryRate(newDeliveryRate int, db *gorm.DB) error {
	delivery.DeliveryRate = newDeliveryRate
	result := db.Save(&delivery)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (delivery *Delivery) ChangeScore(newScore float64, db *gorm.DB) error {
	delivery.Score = newScore
	result := db.Save(&delivery)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
