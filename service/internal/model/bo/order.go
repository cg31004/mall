package bo

import "simon/mall/service/internal/constant"

type Txn struct {
	Id            string
	PaymentNumber string
	Amount        int
	Status        constant.TransactionStatusEnum
	Item          []*TxnItem
}

type TxnItem struct {
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
