package constant

type ProductStatusEnum int

const (
	ProductStatusEnum_Closed ProductStatusEnum = iota
	ProductStatusEnum_Open
)

func (e ProductStatusEnum) Dictionary() map[ProductStatusEnum]string {
	return map[ProductStatusEnum]string{
		ProductStatusEnum_Closed: "無庫存",
		ProductStatusEnum_Open:   "尚有庫存",
	}
}

func (e ProductStatusEnum) String() string {
	return e.Dictionary()[e]
}
