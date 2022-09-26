package repository

import (
	"context"

	"golang.org/x/xerrors"
	"gorm.io/gorm"

	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/utils/timelogger"
)

//go:generate mockery --name ITxnItemRepo --structname MockTxnItemRepo --output mock_repository --outpkg mock_repository --filename mock_transaction_item.go --with-expecter

type ITxnItemRepo interface {
	Insert(ctx context.Context, db *gorm.DB, txnItem []*po.TransactionItem) error
	GetList(ctx context.Context, db *gorm.DB, cond *po.GetTxnItemListCond) ([]*po.TransactionItem, error)
}

type txnItemRepo struct {
	in repositoryIn
}

func newTxnItemRepo(in repositoryIn) ITxnItemRepo {
	return &txnItemRepo{
		in: in,
	}
}

func (repo *txnItemRepo) Insert(ctx context.Context, db *gorm.DB, txnItem []*po.TransactionItem) error {
	defer timelogger.LogTime(ctx)()
	if err := db.Create(txnItem).Error; err != nil {
		return xerrors.Errorf("%w", errs.ConvertDB(err))
	}

	return nil
}

func (repo *txnItemRepo) GetList(ctx context.Context, db *gorm.DB, cond *po.GetTxnItemListCond) ([]*po.TransactionItem, error) {
	defer timelogger.LogTime(ctx)()

	item := make([]*po.TransactionItem, 0)
	if err := db.Model(&item).Scopes(repo.txnItemListCond(cond)).Find(&item).Error; err != nil {
		return nil, xerrors.Errorf("%w", errs.ConvertDB(err))
	}

	return item, nil
}

func (repo *txnItemRepo) txnItemListCond(cond *po.GetTxnItemListCond) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Model(&po.TransactionItem{})

		if len(cond.TxnId) != 0 {
			db = db.Where("`transaction_id` IN ?", cond.TxnId)
		}

		return db
	}
}
