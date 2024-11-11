package model

import (
	"context"

	"gorm.io/gorm"
)

type Order struct{
	gorm.Model
	UserId uint32	`gorm:"type:int(11)"`
	UserCurrency string `gorm:"type:varchar(10)"`
	OrderId string `gorm:"uniqueIndex;type:varchar(100)"`
	Consignee Consignee `gorm:"embedded"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

type Consignee struct {
	Email string
	StreetAddress string
	City string
	State string
	ZipCode int32
	Country string
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context,db *gorm.DB,userId uint32) ([]*Order,error) {
	var orders []*Order
	err := db.WithContext(ctx).Where("user_id = ?",userId).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil,err
	}
	return orders,nil
}