package models

const StatusNew = 1
const StatusCompleted = 2
const StatusCanceled = 3

type Delivery struct {
	ID        int `gorm:"index;type:int" json:"id"`
	CourierID int `gorm:"index;type:int" json:"courierId"`
	OrederID  int `gorm:"index;type:int" json:"orderId"`
	Created   int `gorm:"index;type:bigint" json:"created"`
	Completed int `gorm:"index;type:bigint" json:"completed"`
	Status    int `gorm:"index;type:int" json:"status"`
}

func (s Delivery) TableName() string { return "delivery_orders" }
