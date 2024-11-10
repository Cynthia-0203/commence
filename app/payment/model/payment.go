package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)


type PaymentLog struct{
	gorm.Model
	UserId uint32 `json:"user_id"`
	OrderId string `json:"order_id"`
	TrannsactionId string `json:"transaction_id"`
	Amount float32	`json:"amount"`
	PayAt time.Time `json:"pay_at"`
}

func (PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(db *gorm.DB, ctx context.Context,paymentLog *PaymentLog) error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(paymentLog).Error
}