package bo

import "simon/mall/service/internal/constant"

type Order struct {
	Id            string
	PaymentNumber string
	Amount        int
	Status        constant.TransactionStatusEnum
	Item          []*OrderItem
}

type OrderItem struct {
	Name     string
	Amount   int
	Quantity int
	Image    string
}

type GetTxnItemMapCond struct {
	MemberId string
	TxnId    []string
}

type DelTxnItemMapCond struct {
	MemberId string
}

type CreateOrderCond struct {
	PaymentNumber string
}
