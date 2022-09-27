package dto

type MemberChart struct {
	Id        string `json:"id"`
	Name      string `json:"name"`      // product name
	Amount    int    `json:"amount"`    // 價格
	Quantity  int    `json:"quantity"`  // 數量
	Image     string `json:"image"`     // 圖片
	Inventory string `json:"inventory"` // 庫存
}

type MemberChartUpdateCond struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"` //數量
}

type MemberChartCreateCond struct {
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type MemberChartDelCond struct {
	Id string `json:"id"`
}
