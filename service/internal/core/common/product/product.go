package product

import (
	"context"

	"golang.org/x/xerrors"

	"mall/service/internal/constant"
	"mall/service/internal/model/bo"
	"mall/service/internal/model/po"
	"mall/service/internal/utils/timelogger"
)

type IProductCommon interface {
	GetProduct(ctx context.Context) (map[string]*bo.ProductCommon, error)
}

func newProductCommon(in digIn) IProductCommon {
	return &productCommon{in: in}
}

type productCommon struct {
	in digIn
}

// todo product 更新，需要刷新
func (c *productCommon) GetProduct(ctx context.Context) (map[string]*bo.ProductCommon, error) {
	defer timelogger.LogTime(ctx)()

	if cacheProduct, ok := c.in.Cache.Get(constant.CacheProduct); ok {
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
			Name:  val.Name,
			Image: val.Image,
		}
	}

	c.in.Cache.Save(constant.CacheProduct, result)

	return result, nil
}
