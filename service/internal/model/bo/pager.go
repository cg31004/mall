package bo

type PagerCond struct {
	Index int // 頁碼
	Size  int // 比數
	Order string
}

type PagerResult struct {
	Index     int // 頁碼
	Size      int // 比數
	TotalPage int // 總頁數
	TotalRow  int // 總筆數
}
