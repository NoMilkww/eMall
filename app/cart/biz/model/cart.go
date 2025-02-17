package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductId uint32 `gorm:"type:int(11);not null;"`
	Quantity  uint32 `gorm:"type:int(11);not null;"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, item *Cart) error {
	var row Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row.ID > 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserId: item.UserId, ProductId: item.ProductId}).
			UpdateColumn("quantity", gorm.Expr("quantity + ?", item.Quantity)).
			Error
	}
	return db.WithContext(ctx).Create(item).Error
}

func GetCart(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	var items []*Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: userId}).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	return db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserId: userId}).Delete(&Cart{}).Error
}
