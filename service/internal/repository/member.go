package repository

import (
	"context"

	"golang.org/x/xerrors"
	"gorm.io/gorm"

	"mall/service/internal/errs"
	"mall/service/internal/model/po"
	"mall/service/internal/utils/timelogger"
)

//go:generate mockery --name IMemberRepo --structname MockMemberRepo --output mock_repository --outpkg mock_repository --filename mock_member.go --with-expecter

type IMemberRepo interface {
	First(ctx context.Context, db *gorm.DB, id string) (*po.Member, error)
	FirstByAccount(ctx context.Context, db *gorm.DB, account string) (*po.Member, error)
}

type memberRepo struct {
	in repositoryIn
}

func newMemberRepo(in repositoryIn) IMemberRepo {
	return &memberRepo{
		in: in,
	}
}

func (repo *memberRepo) First(ctx context.Context, db *gorm.DB, id string) (*po.Member, error) {
	defer timelogger.LogTime(ctx)()

	db = db.Where("`id` = ?", id)

	member := &po.Member{}
	if err := db.First(member).Error; err != nil {
		return nil, xerrors.Errorf("memberRepo.First: %w", errs.ConvertDB(err))
	}

	return member, nil
}
func (repo *memberRepo) FirstByAccount(ctx context.Context, db *gorm.DB, account string) (*po.Member, error) {
	defer timelogger.LogTime(ctx)()

	db = db.Where("`member`.`account` = ?", account)

	member := &po.Member{}
	if err := db.First(member).Error; err != nil {
		return nil, xerrors.Errorf("memberRepo.FirstByAccount: %w", errs.ConvertDB(err))
	}

	return member, nil
}
