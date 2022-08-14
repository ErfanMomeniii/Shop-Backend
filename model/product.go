package model

import "time"

type Product struct {
	Id         int
	Name       string
	Price      string
	Company_id int
	Created_At time.Time
}
