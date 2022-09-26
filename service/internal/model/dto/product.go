package dto

import "simon/mall/service/internal/constant"

type GetProductListCond struct {
	*PagerReq
	Name *string `json:"name"`
}

type ProductResp struct {
	Id       string                     `json:"id"`
	Name     string                     `json:"name"`     // 名稱
	Image    string                     `json:"image"`    // 圖片
	Amount   int                        `json:"amount"`   // 價格
	Quantity int                        `json:"quantity"` // 數量
	Status   constant.ProductStatusEnum `json:"status"`
}

type GetProductCond struct {
	Id string
}
