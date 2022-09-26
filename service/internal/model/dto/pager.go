package dto

type PagerReq struct {
	Index int    `form:"pi"` // 頁碼
	Size  int    `form:"ps"` // 比數
	Order string `form:"po"`
}

type PagerResp struct {
	Index     int `json:"pi"`         // 頁碼
	Size      int `json:"ps"`         // 比數
	TotalPage int `json:"total_page"` // 總頁數
	TotalRow  int `json:"total_row"`  // 總筆數
}

type ListResp struct {
	List  interface{} `json:"list"`
	Pager *PagerResp  `json:"pager,omitempty"`
}
