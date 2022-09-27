package product
// todo: 產品功能
//import (
//	"context"
//
//	"github.com/jinzhu/copier"
//	"golang.org/x/xerrors"
//
//	"simon/mall/service/internal/constant"
//	"simon/mall/service/internal/errs"
//	"simon/mall/service/internal/model/bo"
//	"simon/mall/service/internal/model/po"
//	"simon/mall/service/internal/utils/timelogger"
//)
//
//type IProductUseCase interface {
//	GetProductList(ctx context.Context, cond *bo.GetProductListCond) ([]*bo.Product, *bo.PagerResult, error)
//	GetProduct(ctx context.Context, cond *bo.GetProductCond) (*bo.Product, error)
//}
//
//func newProductUseCase(in digIn) IProductUseCase {
//	return &productUseCase{in: in}
//}
//
//type productUseCase struct {
//	in digIn
//}
//
//// todo product 更新，需要刷新
//func (uc *productUseCase) GetProduct(ctx context.Context, cond *bo.GetProductCond) (*bo.Product, error) {
//	defer timelogger.LogTime(ctx)()
//
//	if err := uc.validateGet(ctx, cond); err != nil {
//		return nil, xerrors.Errorf("memberChartUseCase.CreateMemberChart -> validateCreate : %w", err)
//	}
//
//	db := uc.in.DB.Session()
//	product, err := uc.in.ProductRepo.First(ctx, db, cond.Id)
//	if err != nil {
//		return nil, xerrors.Errorf("productUseCase.GetProduct -> ProductRepo.First: %w", err)
//	}
//
//	result := &bo.Product{
//		Id:       product.Id,
//		Name:     product.Name,
//		Image:    product.Image,
//		Amount:   product.Amount,
//		Quantity: product.Quantity,
//		Status:   product.Status,
//	}
//	// 無庫存更改狀態
//	if result.Quantity < 1 {
//		result.Status = constant.ProductStatusEnum_Closed
//	}
//
//	return result, nil
//}
//
//func (uc *productUseCase) validateGet(ctx context.Context, cond *bo.GetProductCond) error {
//	if cond.Id == "" {
//		return xerrors.Errorf("productUseCase.GetProduct -> validateGet.cond.Id :%v,  %w", cond.Id, errs.RequestParamInvalid)
//	}
//
//	return nil
//}
//
//// todo product 根據查詢+cache
//func (uc *productUseCase) GetProductList(ctx context.Context, cond *bo.GetProductListCond) ([]*bo.Product, *bo.PagerResult, error) {
//	defer timelogger.LogTime(ctx)()
//
//	poPager := &po.Pager{}
//	if err := copier.Copy(poPager, cond.PagerCond); err != nil {
//		return nil, nil, xerrors.Errorf("productUseCase.GetProductList -> poPager.copier.Copy: %w", errs.RequestParamParseFailed)
//	}
//
//	db := uc.in.DB.Session()
//	products, err := uc.in.ProductRepo.GetList(ctx, db, &po.ProductSearch{Name: cond.Name}, poPager)
//	if err != nil {
//		return nil, nil, xerrors.Errorf("productUseCase.GetProductList -> ProductRepo.GetList: %w", err)
//	}
//
//	pager, err := uc.in.ProductRepo.GetListPager(ctx, db, &po.ProductSearch{Name: cond.Name}, poPager)
//	if err != nil {
//		return nil, nil, xerrors.Errorf("productUseCase.GetProductList -> ProductRepo.GetListPager: %w", err)
//	}
//
//	boPager := &bo.PagerResult{}
//	if err := copier.Copy(&boPager, pager); err != nil {
//		return nil, nil, xerrors.Errorf("bankUseCase.GetBankList -> boPager.copier.Copy: %w", errs.RequestParamParseFailed)
//	}
//
//	result := make([]*bo.Product, 0, len(products))
//	for _, product := range products {
//		boP := &bo.Product{
//			Id:       product.Id,
//			Name:     product.Name,
//			Image:    product.Image,
//			Amount:   product.Amount,
//			Quantity: product.Quantity,
//			Status:   product.Status,
//		}
//		// 無庫存更改狀態
//		if product.Quantity < 1 {
//			boP.Status = constant.ProductStatusEnum_Closed
//		}
//
//		result = append(result, boP)
//	}
//
//	return result, boPager, nil
//}
