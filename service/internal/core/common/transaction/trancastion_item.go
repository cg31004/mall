package transaction

import (
	"context"

	"golang.org/x/xerrors"

	"simon/mall/service/internal/constant"
	"simon/mall/service/internal/model/bo"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/utils/timelogger"
)

type ITxnItemCommon interface {
	GetTxnItem(ctx context.Context, cond *bo.GetTxnItemMapCond) (map[string][]*bo.OrderItem, error)
}

func newTxnItemCommon(in digIn) ITxnItemCommon {
	return &txnItemCommon{in: in}
}

type txnItemCommon struct {
	in digIn
}

// todo txnItem 更新，需要刷新
func (c *txnItemCommon) GetTxnItem(ctx context.Context, cond *bo.GetTxnItemMapCond) (map[string][]*bo.OrderItem, error) {
	defer timelogger.LogTime(ctx)()

	if cacheProduct, ok := c.in.Cache.Get(constant.CacheMemberTxnItem + cond.MemberId); ok {
		if temp, ok := cacheProduct.(map[string][]*bo.OrderItem); ok {
			return temp, nil
		}
	}

	db := c.in.DB.Session()
	// 不需要分頁查詢
	txnItem, err := c.in.TxnItemRepo.GetList(ctx, db, &po.GetTxnItemListCond{TxnId: cond.TxnId})
	if err != nil {
		return nil, xerrors.Errorf("productCommon.GetProduct -> ProductRepo.GetList: %w", err)
	}

	result := make(map[string][]*bo.OrderItem, len(txnItem))
	for _, val := range txnItem {
		tempItem := &bo.OrderItem{
			Name:     val.Name,
			Amount:   val.Amount,
			Quantity: val.Quantity,
			Image:    val.Image,
		}
		if _, ok := result[val.TransactionId]; !ok {
			result[val.TransactionId] = make([]*bo.OrderItem, 0)
		}
		result[val.TransactionId] = append(result[val.TransactionId], tempItem)
	}

	// 以memberId 當查詢條件
	c.in.Cache.Save(constant.CacheMemberTxnItem+cond.MemberId, result)

	return result, nil
}
