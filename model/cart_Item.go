package model

import (
	"time"
)

type CartItem struct {
	CartItemID  int       `gorm:"column:cart_item_id;AUTO_INCREMENT;primary_key"`
	CartID      int       `gorm:"column:cart_id;NOT NULL"`
	ProductID   int       `gorm:"column:product_id;NOT NULL"`
	Quantity    int       `gorm:"column:quantity;NOT NULL"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	CartData    Cart      `gorm:"foreignKey:CartID;references:CartID"`
	ProductData Product   `gorm:"foreignKey:ProductID;references:ProductID"`
}

func (m *CartItem) TableName() string {
	return "cart_item"
}
