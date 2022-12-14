package transaction

import (
	"context"

	"golang.org/x/xerrors"

	"simon/mall/service/internal/constant"
	"simon/mall/service/internal/model/bo"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/utils/timelogger"
)

//go:generate mockery --name ITxnItemCommon --structname MockTxnItemCommon --output mock_common_txn --outpkg mock_common_txn --filename mock_txn_item.go --with-expecter

type ITxnItemCommon interface {
	GetTxnItem(ctx context.Context, cond *bo.GetTxnItemMapCond) (map[string][]*bo.TxnItem, error)
	DeleteTxnItem(ctx context.Context, cond *bo.DelTxnItemMapCond)
}

func newTxnItemCommon(in digIn) ITxnItemCommon {
	return &txnItemCommon{in: in}
}

type txnItemCommon struct {
	in digIn
}

// todo txnItem 更新，需要刷新
func (c *txnItemCommon) GetTxnItem(ctx context.Context, cond *bo.GetTxnItemMapCond) (map[string][]*bo.TxnItem, error) {
	defer timelogger.LogTime(ctx)()

	if cacheProduct, ok := c.in.Cache.Get(constant.Cache_MemberTxnItem + cond.MemberId); ok {
		if temp, ok := cacheProduct.(map[string][]*bo.TxnItem); ok {
			return temp, nil
		}
	}

	db := c.in.DB.Session()
	// 不需要分頁查詢
	txnItem, err := c.in.TxnItemRepo.GetList(ctx, db, &po.GetTxnItemListCond{TxnId: cond.TxnId})
	if err != nil {
		return nil, xerrors.Errorf("productCommon.GetProduct -> ProductRepo.GetList: %w", err)
	}

	result := make(map[string][]*bo.TxnItem, len(txnItem))
	for _, val := range txnItem {
		tempItem := &bo.TxnItem{
			Name:     val.Name,
			Amount:   val.Amount,
			Quantity: val.Quantity,
			Image:    val.Image,
		}
		if _, ok := result[val.TransactionId]; !ok {
			result[val.TransactionId] = make([]*bo.TxnItem, 0)
		}
		result[val.TransactionId] = append(result[val.TransactionId], tempItem)
	}

	// 以memberId 當查詢條件
	c.in.Cache.Save(constant.Cache_MemberTxnItem+cond.MemberId, result)

	return result, nil
}

func (c *txnItemCommon) DeleteTxnItem(ctx context.Context, cond *bo.DelTxnItemMapCond) {
	c.in.Cache.Delete(constant.Cache_MemberTxnItem + cond.MemberId)
}
