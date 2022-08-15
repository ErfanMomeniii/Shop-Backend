package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Company struct {
	ID      uint
	Name    string
	Address string
	Tel     string
}

type CompanyService interface {
	Validate() error
	AddCompanyToDB(db *gorm.DB) error
	DeleteCompanyFromDB(db *gorm.DB) error
	UpdateCompanyInDB(db *gorm.DB) error
}

func (company *Company) Validate() error {
	return validation.ValidateStruct(&company,
		validation.Field(&company.Name, validation.Required, validation.Length(5, 20)),
	)
}

func (company *Company) AddCompanyToDB(db *gorm.DB) error {
	err := company.Validate()
	if err != nil {
		return err
	}

	result := db.Create(&company)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (company *Company) DeleteCompanyFromDB(db *gorm.DB) error {
	err := DeleteCompanyById(db, int(company.ID))

	if err != nil {
		log.Error(err)
	}

	return err
}

func DeleteCompanyById(db *gorm.DB, id int) error {
	result := db.Where("id = ?", id).Delete(&Company{})

	if result.Error != nil {
		return result.Error
	} else if db.RowsAffected < 1 {
		log.Error("row with id=%d cannot be deleted because it doesn't exist", id)
	}

	return nil
}

func (oldCompany *Company) UpdateCompanyInDB(UpdatedCompany *Company, db *gorm.DB) error {
	err := UpdatedCompany.Validate()

	if err != nil {
		return err
	}

	result := db.Model(&oldCompany).Updates(&UpdatedCompany)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
