package bo

type MemberChart struct {
	Id        string
	Name      string
	Amount    int
	Quantity  int
	Image     string
	Inventory int
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
