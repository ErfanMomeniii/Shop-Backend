package model

import "time"

type Product struct {
	ID         uint
	Name       string
	Price      string
	CompanyID  int
	Company    Company `gorm:"foreignKey:CompanyID"`
	Created_At time.Time
}
