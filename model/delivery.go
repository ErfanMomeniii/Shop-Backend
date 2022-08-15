package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Delivery struct {
	ID           uint
	UserID       int
	User         User `gorm:"foreignKey:UserID"`
	Score        float64
	DeliveryRate int
}

type DeliveryServie interface {
	Validate() error
	AddDeliveryToDB(db *gorm.DB) error
	DeleteDeliveryFromDB(db *gorm.DB) error
	UpdateCompanyInDB(updatedDelivery *Delivery, db *gorm.DB) error
	ChangeDeliveryRate(newDeliveryRate int, db *gorm.DB) error
	ChangeScore(newScore float64, db *gorm.DB) error
}

func (delivery *Delivery) Validate() error {
	return validation.ValidateStruct(&delivery,
		validation.Field(&delivery.UserID, validation.Required),
		validation.Field(&delivery.Score, is.Float),
		validation.Field(&delivery.DeliveryRate, is.Int),
	)
}

func (delivery *Delivery) AddDeliveryToDB(db *gorm.DB) error {
	err := delivery.Validate()
	if err != nil {
		return err
	}

	result := db.Create(&delivery)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (delivery *Delivery) DeleteDeliveryFromDB(db *gorm.DB) error {
	err := DeleteDeliveryById(db, delivery.ID)

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

func (oldDelivery *Delivery) UpdateCompanyInDB(updatedDelivery *Delivery, db *gorm.DB) error {
	err := updatedDelivery.Validate()

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
