package model

type OrderSituation int

const (
	NotStart = iota
	InProgress
	Delivered
)

type Order struct {
	ID         uint
	Products   []int
	DeliveryID int
	Delivery   Delivery `gorm:"foreignKey:DeliveryID"`
	CustomerID int
	Customer   Customer `gorm:"foreignKey:CustomerID"`
	Situation  OrderSituation
}
