package repository

import (
	"context"

	"golang.org/x/xerrors"
	"gorm.io/gorm"

	"mall/service/internal/errs"
	"mall/service/internal/model/po"
	"mall/service/internal/utils/timelogger"
)

//go:generate mockery --name ITxnItemRepo --structname MockTxnItemRepo --output mock_repository --outpkg mock_repository --filename mock_transaction_item.go --with-expecter

type ITxnItemRepo interface {
	Insert(ctx context.Context, db *gorm.DB, txnItem []*po.TransactionItem) error
	GetListByTxnId(ctx context.Context, db *gorm.DB, txnId int64) ([]*po.TransactionItem, error)
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

func (repo *txnItemRepo) GetListByTxnId(ctx context.Context, db *gorm.DB, txnId int64) ([]*po.TransactionItem, error) {
	defer timelogger.LogTime(ctx)()

	item := make([]*po.TransactionItem, 0)
	if err := db.Model(po.TransactionItem{TransactionId: txnId}).Find(&item).Error; err != nil {
		return nil, xerrors.Errorf("%w", errs.ConvertDB(err))
	}

	return item, nil
}
