package product

import (
	"context"

	"golang.org/x/xerrors"

	"simon/mall/service/internal/constant"
	"simon/mall/service/internal/model/bo"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/utils/timelogger"
)

//go:generate mockery --name IProductCommon --structname MockProductCommon --output mock_common_product --outpkg mock_common_product --filename mock_product.go --with-expecter

type IProductCommon interface {
	GetProduct(ctx context.Context) (map[string]*bo.ProductCommon, error)
	DeleteProductCache(ctx context.Context)
}

func newProductCommon(in digIn) IProductCommon {
	return &productCommon{in: in}
}

type productCommon struct {
	in digIn
}

// todo product 更新，需要刷新庫存
func (c *productCommon) GetProduct(ctx context.Context) (map[string]*bo.ProductCommon, error) {
	defer timelogger.LogTime(ctx)()

	if cacheProduct, ok := c.in.Cache.Get(constant.Cache_Product); ok {
		if temp, ok := cacheProduct.(map[string]*bo.ProductCommon); ok {
			return temp, nil
		}
	}

	db := c.in.DB.Session()
	// 不需要分頁查詢
	product, err := c.in.ProductRepo.GetList(ctx, db, &po.ProductSearch{}, nil)
	if err != nil {
		return nil, xerrors.Errorf("productCommon.GetProduct -> ProductRepo.GetList: %w", err)
	}

	result := make(map[string]*bo.ProductCommon, len(product))
	for _, val := range product {
		result[val.Id] = &bo.ProductCommon{
			Name:      val.Name,
			Image:     val.Image,
			Inventory: val.Inventory,
		}
		if val.Inventory < 1 {
			result[val.Id].Status = constant.ProductStatusEnum_Closed
		} else {
			result[val.Id].Status = val.Status
		}
	}

	c.in.Cache.Save(constant.Cache_Product, result)

	return result, nil
}

func (c *productCommon) DeleteProductCache(ctx context.Context) {
	c.in.Cache.Delete(constant.Cache_Product)
}
