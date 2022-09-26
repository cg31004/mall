package repository

import (
	"context"

	"golang.org/x/xerrors"
	"gorm.io/gorm"

	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/utils/timelogger"
)

//go:generate mockery --name IMemberChartRepo --structname MockMemberChartRepo --output mock_repository --outpkg mock_repository --filename mock_member_chart.go --with-expecter

type IMemberChartRepo interface {
	First(ctx context.Context, db *gorm.DB, cond *po.MemberChartFirst) (*po.MemberChart, error)
	Insert(ctx context.Context, db *gorm.DB, chart *po.MemberChart) error
	GetList(ctx context.Context, db *gorm.DB, cond *po.MemberChartSearch) ([]*po.MemberChart, error)
	Delete(ctx context.Context, db *gorm.DB, cond *po.MemberChartDel) error
	Update(ctx context.Context, db *gorm.DB, cond *po.MemberChartUpdate) error
}

type memberChartRepo struct {
	in repositoryIn
}

func newMemberChartRepo(in repositoryIn) IMemberChartRepo {
	return &memberChartRepo{
		in: in,
	}
}

func (repo *memberChartRepo) Insert(ctx context.Context, db *gorm.DB, chart *po.MemberChart) error {
	defer timelogger.LogTime(ctx)()

	if err := db.Create(chart).Error; err != nil {
		return xerrors.Errorf("%w", errs.ConvertDB(err))
	}

	return nil
}

func (repo *memberChartRepo) GetList(ctx context.Context, db *gorm.DB, cond *po.MemberChartSearch) ([]*po.MemberChart, error) {
	defer timelogger.LogTime(ctx)()

	item := make([]*po.MemberChart, 0)
	if err := db.Model(po.MemberChart{MemberId: cond.MemberId}).Find(&item).Error; err != nil {
		return nil, xerrors.Errorf("%w", errs.ConvertDB(err))
	}

	return item, nil
}

func (repo *memberChartRepo) Delete(ctx context.Context, db *gorm.DB, cond *po.MemberChartDel) error {
	defer timelogger.LogTime(ctx)()

	chart := &po.MemberChart{}
	if err := db.Where("id = ?", cond.Id).Delete(chart).Error; err != nil {
		return xerrors.Errorf("%w", errs.ConvertDB(err))
	}

	return nil
}

func (repo *memberChartRepo) Update(ctx context.Context, db *gorm.DB, cond *po.MemberChartUpdate) error {
	defer timelogger.LogTime(ctx)()

	db = db.Where("`id` = ?", cond.Id)
	db = db.Where("`member_id` = ?", cond.MemberId)
	if err := db.Model(po.MemberChart{}).
		Updates(cond.GetUpdateChart()).Error; err != nil {
		return xerrors.Errorf("%w", errs.ConvertDB(err))
	}

	return nil
}

func (*memberChartRepo) First(ctx context.Context, db *gorm.DB, cond *po.MemberChartFirst) (*po.MemberChart, error) {
	defer timelogger.LogTime(ctx)()

	db = db.Where("`id` = ?", cond.ProductId)
	db = db.Where("`member_id` = ?", cond.MemberId)

	chart := &po.MemberChart{}
	if err := db.First(chart).Error; err != nil {
		return nil, xerrors.Errorf("%w", errs.ConvertDB(err))
	}

	return chart, nil

}
