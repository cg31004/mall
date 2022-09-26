package constant

type TransactionStatusEnum int

const (
	TransactionStatusEnum_Process TransactionStatusEnum = iota // 待處理
	TransactionStatusEnum_Success                              // 成功
	TransactionStatusEnum_Failed                               // 失敗
)

func (e TransactionStatusEnum) Dictionary() map[TransactionStatusEnum]string {
	return map[TransactionStatusEnum]string{
		TransactionStatusEnum_Process: "待處理",
		TransactionStatusEnum_Success: "成功",
		TransactionStatusEnum_Failed:  "失敗",
	}
}

func (e TransactionStatusEnum) String() string {
	return e.Dictionary()[e]
}
