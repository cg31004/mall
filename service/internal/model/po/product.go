package po

import (
	"time"

	"simon/mall/service/internal/constant"
)

type Product struct {
	Id        string                     `gorm:"column:id;primary_key"` // 商品 id
	Name      string                     `gorm:"column:name"`           // 商品名稱
	Amount    int                        `gorm:"column:amount"`         //商品價格
	Inventory int                        `gorm:"column:inventory"`      // 庫存
	Image     string                     `gorm:"column:image"`          // 圖片
	Status    constant.ProductStatusEnum `gorm:"column:status"`         // 商品目前狀態

	CreatedAt *time.Time `gorm:"<-:create;column:created_at"`
	UpdatedAt time.Time  `gorm:"<-:create;column:updated_at"`
}

func (Product) TableName() string {
	return "product"
}

type ProductSearch struct {
	Status *constant.ProductStatusEnum
}
