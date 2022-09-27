package order

import (
	"context"
	"sync"

	"golang.org/x/xerrors"
	"gorm.io/gorm"

	"simon/mall/service/internal/constant"
	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/model/bo"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/utils/ctxs"
	"simon/mall/service/internal/utils/timelogger"
)

type IOrderUseCase interface {
	CreateOrder(ctx context.Context, cond *bo.CreateOrderCond) error
	GetOrderList(ctx context.Context) ([]*bo.Order, error)
}

func newOrderUseCase(in digIn) IOrderUseCase {
	return &orderUseCase{in: in}
}

type orderUseCase struct {
	in digIn

	orderLock sync.Mutex
}

// todo: order by channel
// todo: mutex to redis
func (uc *orderUseCase) CreateOrder(ctx context.Context, cond *bo.CreateOrderCond) error {
	defer timelogger.LogTime(ctx)()

	memberInfo, ok := ctxs.GetSession(ctx)
	if !ok {
		return errs.MemberTokenError
	}

	uc.orderLock.Lock()
	defer uc.orderLock.Unlock()
	// refresh product quantity
	defer uc.in.ProductCommon.DeleteProductCache(ctx)
	defer uc.in.TxnItemCommon.DeleteTxnItem(ctx, &bo.DelTxnItemMapCond{MemberId: memberInfo.Id})

	db := uc.in.DB.Session()
	// 取得個人購物車
	chart, err := uc.in.MemberChartRepo.GetList(ctx, db, &po.MemberChartSearch{MemberId: memberInfo.Id})
	if err != nil {
		return xerrors.Errorf("orderUseCase.CreateOrder -> MemberChartRepo.GetList: %w ", err)
	}
	// 各產品價格
	products, err := uc.in.ProductCommon.GetProduct(ctx)
	if err != nil {
		return xerrors.Errorf("orderUseCase.CreateOrder -> ProductCommon.GetProduct: %w ", err)
	}

	txn := func(tx *gorm.DB) error {
		transactionId := uc.in.Uuid.GetUUID()
		// amount 訂單總額
		var amount int

		// txnItem
		poItem := make([]*po.TransactionItem, len(chart))
		for _, c := range chart {
			// 商品已下架或沒有庫存
			if val, ok := products[c.ProductId]; !ok || val.Inventory < c.Quantity || val.Status != constant.ProductStatusEnum_Open {
				return xerrors.Errorf("orderUseCase.CreateOrder -> ProductCommon.GetProduct: %w ", errs.OrderProductNoMatch)
			}

			item := &po.TransactionItem{
				TransactionId: transactionId,
				Name:          products[c.ProductId].Name,
				Amount:        products[c.ProductId].Amount,
				Quantity:      c.Quantity,
				Image:         products[c.ProductId].Image,
			}

			poItem = append(poItem, item)
			amount += products[c.ProductId].Amount * c.Quantity
		}

		if err := uc.in.TxnItemRepo.Insert(ctx, db, poItem); err != nil {
			return xerrors.Errorf("orderUseCase.CreateOrder -> TxnItemRepo.Insert: %w ", err)
		}

		transaction := &po.Transaction{
			Id:            transactionId,
			MemberId:      memberInfo.Id,
			PaymentNumber: cond.PaymentNumber,
			Status:        constant.TransactionStatusEnum_Success,
			Amount:        amount,
		}

		if err := uc.in.TxnRepo.Insert(ctx, db, transaction); err != nil {
			return xerrors.Errorf("orderUseCase.CreateOrder -> TxnRepo.Insert: %w ", err)
		}

		return nil
	}
	if err := db.Transaction(txn); err != nil {
		return xerrors.Errorf("sb *syncBankUseCase -> db.Transaction: %w", err)
	}

	return nil
}

func (uc *orderUseCase) GetOrderList(ctx context.Context) ([]*bo.Order, error) {
	defer timelogger.LogTime(ctx)()

	memberInfo, ok := ctxs.GetSession(ctx)
	if !ok {
		return nil, errs.MemberTokenError
	}

	db := uc.in.DB.Session()

	txn, err := uc.in.TxnRepo.GetList(ctx, db, &po.TransactionSearch{MemberId: &memberInfo.Id})
	if err != nil {
		return nil, xerrors.Errorf("orderUseCase.GetOrderList -> TxnRepo.GetList : %w", err)
	}

	// 取得各訂單資料
	var txnItemCond []string
	for _, t := range txn {
		txnItemCond = append(txnItemCond, t.Id)
	}
	txnItem, err := uc.in.TxnItemCommon.GetTxnItem(ctx, &bo.GetTxnItemMapCond{
		MemberId: memberInfo.Id,
		TxnId:    txnItemCond,
	})
	if err != nil {
		return nil, xerrors.Errorf("orderUseCase.GetOrderList -> TxnItemCommon.GetTxnItem : %w", err)
	}

	// make order data
	result := make([]*bo.Order, 0, len(txn))
	for _, t := range txn {
		result = append(result, &bo.Order{
			Id:            t.Id,
			PaymentNumber: t.PaymentNumber,
			Amount:        t.Amount,
			Status:        t.Status,
			Item:          txnItem[t.Id],
		})
	}

	return result, nil
}
