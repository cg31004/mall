package repository

import (
	"context"

	"golang.org/x/xerrors"
	"gorm.io/gorm"

	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/utils/timelogger"
)

//go:generate mockery --name IProductRepo --structname MockPaymentRepo --output mock_repository --outpkg mock_repository --filename mock_product.go --with-expecter

type IProductRepo interface {
	First(ctx context.Context, db *gorm.DB, id string) (*po.Product, error)
	GetList(ctx context.Context, db *gorm.DB, cond *po.ProductSearch, pager *po.Pager) ([]*po.Product, error)
	// todo: 產品功能:頁面
	//GetListPager(ctx context.Context, db *gorm.DB, cond *po.ProductSearch, pager *po.Pager) (*po.PagingResult, error)
}

type productRepo struct {
	in repositoryIn
}

func newProductRepo(in repositoryIn) IProductRepo {
	return &productRepo{
		in: in,
	}
}

func (repo *productRepo) First(ctx context.Context, db *gorm.DB, id string) (*po.Product, error) {
	defer timelogger.LogTime(ctx)()

	db = db.Where("`id` = ?", id)

	product := &po.Product{}
	if err := db.First(product).Error; err != nil {
		return nil, xerrors.Errorf("productRepo.First: %w", errs.ConvertDB(err))
	}

	return product, nil
}

func (repo *productRepo) GetList(ctx context.Context, db *gorm.DB, cond *po.ProductSearch, pager *po.Pager) ([]*po.Product, error) {
	defer timelogger.LogTime(ctx)()

	products := make([]*po.Product, 0)

	if err := db.Model(&products).Scopes(repo.productListCond(cond, pager)).Find(&products).Error; err != nil {
		return nil, xerrors.Errorf("productRepo.GetList: %w", errs.ConvertDB(err))
	}

	return products, nil
}

// todo: 產品功能:頁面
//func (repo *productRepo) GetListPager(ctx context.Context, db *gorm.DB, cond *po.ProductSearch, pager *po.Pager) (*po.PagingResult, error) {
//	defer timelogger.LogTime(ctx)()
//
//	var count int64
//	if err := db.
//		Scopes(repo.productListCond(cond, nil)).
//		Count(&count).
//		Error; err != nil {
//		return nil, xerrors.Errorf("%w", errs.ConvertDB(err))
//	}
//
//	return po.NewPagerResult(pager, count), nil
//}

func (repo *productRepo) productListCond(cond *po.ProductSearch, pager *po.Pager) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Model(&po.Product{})

		if cond != nil {
			if cond.Status != nil {
				db = db.Where("`status` = ?", cond.Status)
			}
		}

		if pager != nil {
			db.Scopes(parsePaging(pager))
		}

		return db
	}
}
