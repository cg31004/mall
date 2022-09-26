package repository

import (
	"context"

	"golang.org/x/xerrors"
	"gorm.io/gorm"

	"mall/service/internal/errs"
	"mall/service/internal/model/po"
	"mall/service/internal/utils/timelogger"
)

//go:generate mockery --name IPaymentRepo --structname MockPaymentRepo --output mock_repository --outpkg mock_repository --filename mock_payment.go --with-expecter

type IPaymentRepo interface {
	First(ctx context.Context, db *gorm.DB, memberId string) (*po.Payment, error)
}

type paymentRepo struct {
	in repositoryIn
}

func newPaymentRepo(in repositoryIn) IPaymentRepo {
	return &paymentRepo{
		in: in,
	}
}

func (repo *paymentRepo) First(ctx context.Context, db *gorm.DB, memberId string) (*po.Payment, error) {
	defer timelogger.LogTime(ctx)()

	db = db.Where("`member_id` = ?", memberId)

	payment := &po.Payment{}
	if err := db.First(payment).Error; err != nil {
		return nil, xerrors.Errorf("%w", errs.ConvertDB(err))
	}

	return payment, nil
}
