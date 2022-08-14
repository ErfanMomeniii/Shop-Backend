package model

type Order_situation int

const (
	NotStart = iota
	InProgress
	Delivered
)

type Order struct {
	Id          int
	Products    []int //products_id
	Delivery_id int
	Customer_id int
	Situation   Order_situation
}
