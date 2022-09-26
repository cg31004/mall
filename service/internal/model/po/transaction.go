package po

import (
	"time"

	"simon/mall/service/internal/constant"
)

type Transaction struct {
	Id            string                         `gorm:"column:id;primary_key"` // id
	MemberId      string                         `gorm:"column:member_id"`      // 會員id
	PaymentNumber string                         `gorm:"column:payment_number"` // 支付號碼
	Amount        int                            `gorm:"column:amount"`         // 金額
	Status        constant.TransactionStatusEnum `gorm:"column:status"`         // 訂單狀態 0： 待處理, 1： 成功, 2： 失敗
	CreatedAt     time.Time                      `gorm:"<-:create;column:created_at"`
	UpdatedAt     time.Time                      `gorm:"<-:create;column:updated_at"`
}

func (Transaction) TableName() string {
	return "transaction"
}

type TransactionSearch struct {
	MemberId *string
}

//
type TransactionItem struct {
	Id            int                            `gorm:"column:id"`             // id
	TransactionId string                         `gorm:"column:transaction_id"` // transaction id
	Name          string                         `gorm:"column:name"`           // 商品名稱
	Amount        int                            `gorm:"column:"amount"`        //商品價格
	Quantity      int                            `gorm:"column:"quantity"`      // 庫存
	Image         string                         `gorm:"column:"image"`         // 圖片
}

func (TransactionItem) TableName() string {
	return "transaction_item"
}

type GetTxnItemListCond struct {
	TxnId []string
}
