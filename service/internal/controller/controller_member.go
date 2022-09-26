package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"mall/service/internal/controller/middleware"
	"mall/service/internal/errs"
	"mall/service/internal/model/bo"
	"mall/service/internal/model/dto"
	"mall/service/internal/utils/ctxs"
	"mall/service/internal/utils/timelogger"
)

type IMemberCtrl interface {
	// Session
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)

	//Chart
	GetChart(ctx *gin.Context)
	UpdateChart(ctx *gin.Context)
	CreateChart(ctx *gin.Context)
	DeleteChart(ctx *gin.Context)
}

func newMember(in digIn) IMemberCtrl {
	return &memberCtrl{
		in: in,
	}
}

type memberCtrl struct {
	in digIn
}

func (ctrl *memberCtrl) Login(ctx *gin.Context) {
	defer timelogger.LogTime(ctx)()

	req := &dto.MemberSessionCond{}
	if err := ctrl.in.Request.Bind(ctx, &req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.RequestParamParseFailed)
		return
	}

	cond := &bo.MemberSessionCond{}
	if err := copier.Copy(cond, req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	boResp, err := ctrl.in.MemberIn.Session.Login(ctx, cond)
	if err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, err)
		return
	}

	var dtoResp dto.MemberToken
	if err := copier.Copy(&dtoResp, boResp); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	ctrl.in.SetResponse.StandardResp(ctx, http.StatusOK, dtoResp)
}

func (ctrl *memberCtrl) Logout(ctx *gin.Context) {
	defer timelogger.LogTime(ctx)()

	session, ok := ctxs.GetSession(ctx)
	if !ok {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.MemberTokenError)
		return
	}

	cond := &bo.MemberToken{
		Token: session.Token,
	}

	err := ctrl.in.MemberIn.Session.Logout(ctx, cond)
	if err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, err)
		return
	}

	ctrl.in.SetResponse.StandardResp(ctx, http.StatusOK, middleware.Resp_Success)

}

func (ctrl *memberCtrl) GetChart(ctx *gin.Context) {
	defer timelogger.LogTime(ctx)()

	boResp, err := ctrl.in.MemberIn.Chart.GetMemberChart(ctx)
	if err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, err)
		return
	}

	var dtoResp dto.MemberChart
	if err := copier.Copy(&dtoResp, boResp); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	ctrl.in.SetResponse.StandardResp(ctx, http.StatusOK, dto.ListResp{List: dtoResp})
}

func (ctrl *memberCtrl) UpdateChart(ctx *gin.Context) {
	defer timelogger.LogTime(ctx)()

	req := &dto.MemberChartUpdateCond{}
	if err := ctrl.in.Request.Bind(ctx, &req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.RequestParamParseFailed)
		return
	}

	cond := &bo.MemberChartUpdateCond{}
	if err := copier.Copy(cond, req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	if err := ctrl.in.MemberIn.Chart.UpdateMemberChart(ctx, cond); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, err)
		return
	}

	ctrl.in.SetResponse.StandardResp(ctx, http.StatusOK, middleware.Resp_Success)
}

func (ctrl *memberCtrl) CreateChart(ctx *gin.Context) {
	defer timelogger.LogTime(ctx)()

	req := &dto.MemberChartCreateCond{}
	if err := ctrl.in.Request.Bind(ctx, &req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.RequestParamParseFailed)
		return
	}

	cond := &bo.MemberChartCreateCond{}
	if err := copier.Copy(cond, req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	if err := ctrl.in.MemberIn.Chart.CreateMemberChart(ctx, cond); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, err)
		return
	}

	ctrl.in.SetResponse.StandardResp(ctx, http.StatusOK, middleware.Resp_Success)
}

func (ctrl *memberCtrl) DeleteChart(ctx *gin.Context) {
	defer timelogger.LogTime(ctx)()

	req := &dto.MemberChartDelCond{}
	if err := ctrl.in.Request.Bind(ctx, &req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.RequestParamParseFailed)
		return
	}

	cond := &bo.MemberChartDelCond{}
	if err := copier.Copy(cond, req); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, errs.DataConvertError)
		return
	}

	if err := ctrl.in.MemberIn.Chart.DeleteMemberChart(ctx, cond); err != nil {
		ctrl.in.SetResponse.StandardResp(ctx, http.StatusBadRequest, err)
		return
	}

	ctrl.in.SetResponse.StandardResp(ctx, http.StatusOK, middleware.Resp_Success)
}
