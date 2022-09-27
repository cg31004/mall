package bo

import "simon/mall/service/internal/constant"

type MemberChart struct {
	Id        string
	Name      string
	Amount    int
	Quantity  int
	Image     string
	Inventory int
	Status    constant.ProductStatusEnum
}

type MemberChartUpdateCond struct {
	Id       string
	Quantity int
}

type MemberChartCreateCond struct {
	ProductId string
	Quantity  int
}

type MemberChartDelCond struct {
	Id string
}
