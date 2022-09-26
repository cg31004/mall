package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/model/dto"
	"simon/mall/service/internal/utils/timelogger"
)

type IOrderCtrl interface {
	Order(ctx *gin.Context)
	GetOrderList(ctx *gin.Context)
}

func newOrder(in digIn) IOrderCtrl {
	return &orderCtrl{
		in: in,
	}
}

type orderCtrl struct {
	in digIn
}

func (ctrl *orderCtrl) Order(ctx *gin.Context) {

}

func (ctrl *orderCtrl) GetOrderList(ctx *gin.Context) {
	defer timelogger.LogTime(ctx)()

	boResp, err := ctrl.in.OrderIn.Order.GetOrderList(ctx)
	if err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, err)
		return
	}

	dtoResp := make([]*dto.OrderResp, len(boResp))
	if err := copier.Copy(&dtoResp, boResp); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	resp := &dto.ListResp{
		List: dtoResp,
	}

	ctrl.in.SetResponse.StandardResp(ctx, http.StatusOK, resp)
}
