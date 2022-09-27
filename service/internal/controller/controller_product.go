package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/model/bo"
	"simon/mall/service/internal/model/dto"
	"simon/mall/service/internal/utils/timelogger"
)

// todo: 產品控制
type IProductCtrl interface {
	Get(ctx *gin.Context)
	GetList(ctx *gin.Context)
}

func newProduct(in digIn) IProductCtrl {
	return &productCtrl{
		in: in,
	}
}

type productCtrl struct {
	in digIn
}

// todo 查詢條件
func (ctrl *productCtrl) GetList(ctx *gin.Context) {
	defer timelogger.LogTime(ctx)()

	req := &dto.GetProductListCond{}
	if err := ctrl.in.Request.Bind(ctx, &req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.RequestParamParseFailed)
		return
	}

	cond := &bo.GetProductListCond{}
	if err := copier.Copy(cond, req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	boResp, pager, err := ctrl.in.ProductIn.Product.GetProductList(ctx, cond)
	if err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, err)
		return
	}

	dtoResp := make([]*dto.ProductResp, len(boResp))
	if err := copier.Copy(&dtoResp, boResp); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	dtoPagerResp := &dto.PagerResp{}
	if err := copier.Copy(dtoPagerResp, pager); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	resp := &dto.ListResp{
		List:  dtoResp,
		Pager: dtoPagerResp,
	}

	ctrl.in.SetResponse.StandardResp(ctx, http.StatusOK, resp)
}

func (ctrl *productCtrl) Get(ctx *gin.Context) {
	defer timelogger.LogTime(ctx)()

	req := &dto.GetProductCond{}
	if err := ctrl.in.Request.Bind(ctx, &req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.RequestParamParseFailed)
		return
	}

	cond := &bo.GetProductCond{}
	if err := copier.Copy(cond, req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	boResp, err := ctrl.in.ProductIn.Product.GetProduct(ctx, cond)
	if err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, err)
		return
	}

	var dtoResp dto.ProductResp
	if err := copier.Copy(&dtoResp, boResp); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	ctrl.in.SetResponse.StandardResp(ctx, http.StatusOK, dtoResp)
}
