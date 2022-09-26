package dto

import "simon/mall/service/internal/constant"

type OrderResp struct {
	Id            string                         `json:"id"`
	PaymentNumber string                         `json:"payment_number"` // 卡號
	Amount        int                            `json:"amount"`         // 價格
	Status        constant.TransactionStatusEnum `json:"status"`
	Item          []*OrderItem                   `json:"item"`
}

type OrderItem struct {
	Name     string `json:"name"`     // 名字
	Amount   int    `json:"amount"`   // 價格
	Quantity int    `json:"quantity"` //數量
	Image    string `json:"image"`    // 圖片
}

type CreateOrderCond struct {
	PaymentNumber string `json:"payment_number"` //卡號
}
