package po

const (
	pagingDefaultIndex = 1
	pagingDefaultSize  = 10
)

type Pager struct {
	Index int    // 頁碼
	Size  int    // 筆數
	Order string // 排序
}

func (p *Pager) GetIndex() int {
	if p.Index < 1 {
		return pagingDefaultIndex
	}

	return p.Index
}

func (p *Pager) GetSize() int {
	if p.Size < 1 {
		return pagingDefaultSize
	}

	return p.Size
}

func (p *Pager) GetOffset() int {
	return p.GetSize() * (p.GetIndex() - 1)
}

func (p *Pager) GetPaging() (int, int) {
	return p.GetIndex(), p.GetSize()
}

func (p *Pager) PagingPtr() *Pager {
	return p
}

func NewPagerResult(paging *Pager, count int64) *PagingResult {
	totalPage := int(count) / paging.GetSize()
	if int(count)%paging.GetSize() > 0 {
		totalPage++
	}

	return &PagingResult{
		Index:     paging.GetIndex(),
		Size:      paging.GetSize(),
		TotalPage: totalPage,
		TotalRow:  int(count),
	}
}

type PagingResult struct {
	Index     int // 頁碼
	Size      int // 比數
	TotalPage int // 總頁數
	TotalRow  int // 總筆數
}
