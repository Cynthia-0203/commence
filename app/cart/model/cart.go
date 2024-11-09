package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct{
	gorm.Model
	UserID uint32 `gorm:"not null;type:int(11);index:idx_user_id"`
	ProductID uint32	`gorm:"not null;type:int(11)"`
	Qty uint32	`gorm:"not null;type:int(11)"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context,db *gorm.DB,item *Cart)error{
	var row Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: item.UserID,ProductID: item.ProductID}).First(&row).Error
	if err != nil && !errors.Is(err,gorm.ErrRecordNotFound){
		return err
	}
	if row.ID > 0{
		return db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: item.UserID,ProductID: item.ProductID}).UpdateColumn("qty",gorm.Expr("qty + ?",item.Qty)).Error
	}
	return db.WithContext(ctx).Create(item).Error
}

func EmptyCart(ctx context.Context,db *gorm.DB,userID uint32)error{
	if userID == 0{
		return errors.New("user id is required")
	}
	return db.WithContext(ctx).Delete(&Cart{},"user_id = ?",userID).Error
}

func GetCartByUserId(ctx context.Context,db *gorm.DB,userID uint32) ([]*Cart,error){
	var rows []*Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where("user_id = ?",userID).Find(&rows).Error
	if err != nil {
		return nil, err
	}

	return rows,nil
}