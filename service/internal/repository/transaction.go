package repository

import (
	"context"

	"golang.org/x/xerrors"
	"gorm.io/gorm"

	"mall/service/internal/errs"
	"mall/service/internal/model/po"
	"mall/service/internal/utils/converter"
	"mall/service/internal/utils/timelogger"
)

//go:generate mockery --name ITxnRepo --structname MockTxnRepo --output mock_repository --outpkg mock_repository --filename mock_transaction.go --with-expecter

type ITxnRepo interface {
	Insert(ctx context.Context, db *gorm.DB, txnItem *po.Transaction) error
	GetList(ctx context.Context, db *gorm.DB, cond *po.TransactionSearch) ([]*po.Transaction, error)
}

type txnRepo struct {
	in repositoryIn
}

func newTxnRepo(in repositoryIn) ITxnRepo {
	return &txnRepo{
		in: in,
	}
}

func (repo *txnRepo) Insert(ctx context.Context, db *gorm.DB, txn *po.Transaction) error {
	defer timelogger.LogTime(ctx)()
	if err := db.Create(txn).Error; err != nil {
		return xerrors.Errorf("%w", errs.ConvertDB(err))
	}

	return nil
}

func (repo *txnRepo) GetList(ctx context.Context, db *gorm.DB, cond *po.TransactionSearch) ([]*po.Transaction, error) {
	defer timelogger.LogTime(ctx)()

	item := make([]*po.Transaction, 0)
	if err := db.Scopes(repo.listCond(cond)).Find(&item).Error; err != nil {
		return nil, xerrors.Errorf("%w", errs.ConvertDB(err))
	}

	return item, nil
}

func (repo *txnRepo) listCond(cond *po.TransactionSearch) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(converter.ConvertInt64ToStr(cond.MemberId)) > 0 {
			db = db.Where("`member_id` = ?", cond.MemberId)
		}

		return db
	}
}
