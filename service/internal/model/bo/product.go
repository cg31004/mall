package bo

import "simon/mall/service/internal/constant"

type ProductCommon struct {
	Name      string
	Image     string
	Amount    int
	Inventory int
	Status    constant.ProductStatusEnum
}

type GetProductListCond struct {
	*PagerCond
	Name *string
}

type GetProductCond struct {
	Id string
}

type Product struct {
	Id       string
	Name     string
	Image    string
	Amount   int
	Quantity int
	Status   constant.ProductStatusEnum
}
