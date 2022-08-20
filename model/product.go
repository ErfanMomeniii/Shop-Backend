package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID         uint
	Name       string
	Price      string
	CompanyID  int
	Company    Company `gorm:"foreignKey:CompanyID"`
	Created_At time.Time
}

type ProductService interface {
	Validate() error
	AddProductToDB(db *gorm.DB) error
	UpdateProductInDB(db *gorm.DB) error
}

func (product *Product) Validate() error {
	return validation.ValidateStruct(&product,
		validation.Field(&product.Name, validation.Required),
		validation.Field(&product.Price, validation.Required),
	)
}

func (product *Product) AddProductToDB(db *gorm.DB)error{
	err := product.Validate()
	if err != nil {
		return err
	}

	result := db.Create(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (oldProduct *Product) UpdateProductInDB(updatedProduct *Product, db *gorm.DB) error {
	err := updatedProduct.Validate()

	if err != nil {
		return err
	}

	result := db.Model(&oldProduct).Updates(&updatedProduct)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
