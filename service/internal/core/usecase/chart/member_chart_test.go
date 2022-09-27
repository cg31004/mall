package chart

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"simon/mall/service/internal/constant"
	"simon/mall/service/internal/core/common/product/mock_common_product"
	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/model/bo"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/repository/mock_repository"
	"simon/mall/service/internal/thirdparty/mysqlcli"
	"simon/mall/service/internal/utils/ctxs"
	"simon/mall/service/internal/utils/uuid/mock_uuid"
)

type chartSuit struct {
	suite.Suite

	ctx context.Context
	*memberChartUseCase
}

func TestChart(t *testing.T) {
	suite.Run(t, &chartSuit{})
}

// 測試初始化
func (s *chartSuit) SetupSuite() {
	gin.SetMode(gin.TestMode)
	s.memberChartUseCase = &memberChartUseCase{}
}

const (
	memberDefaultId      = "111"
	memberDefaultAccount = "simon"
	memberDefaultName    = "simon"
	memberDefaultToken   = "222"
	chartUuid            = "1"
)

func (s *chartSuit) SetupTest() {
	s.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
	ctxs.SetSession(s.ctx.(*gin.Context), &bo.MemberSession{Id: memberDefaultId, Account: memberDefaultAccount, Name: memberDefaultName, Token: memberDefaultToken})
	s.in.DB = mysqlcli.NewMockClient()
	s.in.ProductCommon = mock_common_product.NewMockProductCommon(s.T())
	s.in.MemberChartRepo = mock_repository.NewMockMemberChartRepo(s.T())
	s.in.Uuid = mock_uuid.NewMockUuid(s.T())
}

func (s *chartSuit) Test_MemberChart_Delete() {
	var cond *bo.MemberChartDelCond
	var err error

	//
	s.SetupTest()
	s.T().Log("delete session fail")
	s.ctx = context.Background()
	cond = &bo.MemberChartDelCond{}
	err = s.memberChartUseCase.DeleteMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.MemberTokenError, err)

	//
	s.SetupTest()
	s.T().Log("delete fail return error")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().Delete(s.ctx, mock.Anything, mock.Anything).Return(errs.CommonUnknownError)
	cond = &bo.MemberChartDelCond{}
	err = s.memberChartUseCase.DeleteMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("delete ok return nil")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().Delete(s.ctx, mock.Anything, mock.Anything).Return(nil)
	cond = &bo.MemberChartDelCond{}
	err = s.memberChartUseCase.DeleteMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(nil, err)

	//
	s.SetupTest()
	s.T().Log("delete value test")
	cond = &bo.MemberChartDelCond{Id: "111"}
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().Delete(s.ctx, mock.Anything, &po.MemberChartDel{Id: "111", MemberId: memberDefaultId}).Return(nil)
	_ = s.memberChartUseCase.DeleteMemberChart(s.ctx, cond)
}

func (s *chartSuit) Test_MemberChart_Update() {
	var cond *bo.MemberChartUpdateCond
	var err error
	memberChartId := "123"

	//
	s.SetupTest()
	s.T().Log("update session fail")
	s.ctx = context.Background()
	cond = &bo.MemberChartUpdateCond{}
	err = s.memberChartUseCase.UpdateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.MemberTokenError, err)

	//
	s.SetupTest()
	s.T().Log("update validate fail: no id")
	cond = &bo.MemberChartUpdateCond{}
	err = s.memberChartUseCase.UpdateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.RequestParamInvalid, err)

	//
	s.SetupTest()
	s.T().Log("update validate fail: no quantity")
	cond = &bo.MemberChartUpdateCond{Id: memberChartId, Quantity: -1}
	err = s.memberChartUseCase.UpdateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.RequestParamInvalid, err)

	//
	s.SetupTest()
	s.T().Log("update validate fail: quantity = 0")
	cond = &bo.MemberChartUpdateCond{Id: memberChartId, Quantity: 0}
	err = s.memberChartUseCase.UpdateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.RequestParamInvalid, err)

	//
	s.SetupTest()
	s.T().Log("update fail return error")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().Update(s.ctx, mock.Anything, mock.Anything).Return(errs.CommonUnknownError)
	cond = &bo.MemberChartUpdateCond{Id: memberChartId, Quantity: 1}
	err = s.memberChartUseCase.UpdateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("update ok return nil")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().Update(s.ctx, mock.Anything, mock.Anything).Return(nil)
	cond = &bo.MemberChartUpdateCond{Id: memberChartId, Quantity: 1}
	err = s.memberChartUseCase.UpdateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(nil, err)

	//
	s.SetupTest()
	s.T().Log("delete value test")
	cond = &bo.MemberChartUpdateCond{Id: memberChartId, Quantity: 1}
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().Update(s.ctx, mock.Anything, &po.MemberChartUpdate{Id: memberChartId, MemberId: memberDefaultId, Quantity: 1}).Return(nil)
	_ = s.memberChartUseCase.UpdateMemberChart(s.ctx, cond)
}

func (s *chartSuit) Test_MemberChart_Create() {
	var cond *bo.MemberChartCreateCond
	var err error
	productId := "123"

	//
	s.SetupTest()
	s.T().Log("create session fail")
	s.ctx = context.Background()
	cond = &bo.MemberChartCreateCond{}
	err = s.memberChartUseCase.CreateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.MemberTokenError, err)

	//
	s.SetupTest()
	s.T().Log("create validate fail: no quantity")
	cond = &bo.MemberChartCreateCond{ProductId: productId, Quantity: -1}
	err = s.memberChartUseCase.CreateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.RequestParamInvalid, err)

	//
	s.SetupTest()
	s.T().Log("create validate fail: quantity = 0")
	cond = &bo.MemberChartCreateCond{ProductId: productId, Quantity: 0}
	err = s.memberChartUseCase.CreateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.RequestParamInvalid, err)

	//
	s.SetupTest()
	s.T().Log("create: first fail return error, check first parameter")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().First(s.ctx, mock.Anything, &po.MemberChartFirst{
		MemberId:  memberDefaultId,
		ProductId: productId,
	}).Return(nil, errs.CommonUnknownError)
	cond = &bo.MemberChartCreateCond{ProductId: productId, Quantity: 1}
	err = s.memberChartUseCase.CreateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("create: first fail than return no rows -> insert")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().First(s.ctx, mock.Anything, mock.Anything).Return(nil, errs.ErrDB.NoRow)
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().Insert(s.ctx, mock.Anything, mock.Anything).Return(errs.CommonUnknownError)
	s.in.Uuid.(*mock_uuid.MockUuid).EXPECT().GetUUID().Return(chartUuid)
	cond = &bo.MemberChartCreateCond{ProductId: productId, Quantity: 1}
	err = s.memberChartUseCase.CreateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("create: first fail than return no rows -> insert fail , check insert parameter")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().First(s.ctx, mock.Anything, mock.Anything).Return(nil, errs.ErrDB.NoRow)
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().Insert(s.ctx, mock.Anything, &po.MemberChart{
		Id:        chartUuid,
		MemberId:  memberDefaultId,
		ProductId: productId,
		Quantity:  1,
	}).Return(errs.CommonUnknownError)
	s.in.Uuid.(*mock_uuid.MockUuid).EXPECT().GetUUID().Return(chartUuid)
	cond = &bo.MemberChartCreateCond{ProductId: productId, Quantity: 1}
	err = s.memberChartUseCase.CreateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("create: first fail than return no rows -> insert success")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().First(s.ctx, mock.Anything, mock.Anything).Return(nil, errs.ErrDB.NoRow)
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().Insert(s.ctx, mock.Anything, mock.Anything).Return(nil)
	s.in.Uuid.(*mock_uuid.MockUuid).EXPECT().GetUUID().Return(chartUuid)
	cond = &bo.MemberChartCreateCond{ProductId: productId, Quantity: 1}
	err = s.memberChartUseCase.CreateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(nil, err)

	chartFirst := &po.MemberChart{
		Id:        chartUuid,
		MemberId:  memberDefaultId,
		ProductId: productId,
		Quantity:  2,
	}
	//
	s.SetupTest()
	s.T().Log("create: first fail than return no rows -> update fail , check update parameter")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().First(s.ctx, mock.Anything, mock.Anything).Return(chartFirst, nil)
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().Update(s.ctx, mock.Anything, &po.MemberChartUpdate{
		Id:       chartUuid,
		MemberId: memberDefaultId,
		Quantity: 3,
	}).Return(errs.CommonUnknownError)
	cond = &bo.MemberChartCreateCond{ProductId: productId, Quantity: 1}
	err = s.memberChartUseCase.CreateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("create: first fail than return no rows -> update success")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().First(s.ctx, mock.Anything, mock.Anything).Return(chartFirst, nil)
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().Update(s.ctx, mock.Anything, mock.Anything).Return(nil)
	cond = &bo.MemberChartCreateCond{ProductId: productId, Quantity: 1}
	err = s.memberChartUseCase.CreateMemberChart(s.ctx, cond)
	s.Assert().ErrorIs(nil, err)
}

func (s *chartSuit) Test_MemberChart_Get() {
	var err error

	//
	s.SetupTest()
	s.T().Log("get session fail")
	s.ctx = context.Background()
	_, err = s.memberChartUseCase.GetMemberChart(s.ctx)
	s.Assert().ErrorIs(errs.MemberTokenError, err)

	//
	s.SetupTest()
	s.T().Log("get: list fail return error, check list parameter")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().GetList(s.ctx, mock.Anything, &po.MemberChartSearch{MemberId: memberDefaultId}).Return(nil, errs.CommonUnknownError)
	_, err = s.memberChartUseCase.GetMemberChart(s.ctx)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("get: list fail return error, check list parameter")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().GetList(s.ctx, mock.Anything, &po.MemberChartSearch{MemberId: memberDefaultId}).Return(nil, errs.CommonUnknownError)
	_, err = s.memberChartUseCase.GetMemberChart(s.ctx)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("get: prodcut common fail return error, check list parameter")
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().GetList(s.ctx, mock.Anything, mock.Anything).Return(nil, nil)
	s.in.ProductCommon.(*mock_common_product.MockProductCommon).EXPECT().GetProduct(s.ctx).Return(nil, errs.CommonUnknownError)
	_, err = s.memberChartUseCase.GetMemberChart(s.ctx)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("get:combine chart")
	product := map[string]*bo.ProductCommon{
		"1": {Name: "product1", Image: "image1", Amount: 10, Inventory: 4, Status: constant.ProductStatusEnum_Closed},
		"2": {Name: "product2", Image: "image2", Amount: 50, Inventory: 1, Status: constant.ProductStatusEnum_Open},
	}
	poChart := []*po.MemberChart{
		{Id: "1", MemberId: memberDefaultId, ProductId: "4", Quantity: 1},
		{Id: "2", MemberId: memberDefaultId, ProductId: "3", Quantity: 2},
		{Id: "3", MemberId: memberDefaultId, ProductId: "2", Quantity: 4},
		{Id: "4", MemberId: memberDefaultId, ProductId: "1", Quantity: 8},
	}
	wantCharts := []*bo.MemberChart{
		{Id: "1", Name: constant.Unknown_Product, Amount: 0, Quantity: 1, Image: "", Inventory: 0, Status: constant.ProductStatusEnum_Closed},
		{Id: "2", Name: constant.Unknown_Product, Amount: 0, Quantity: 2, Image: "", Inventory: 0, Status: constant.ProductStatusEnum_Closed},
		{Id: "3", Name: "product2", Amount: 50, Quantity: 4, Image: "image2", Inventory: 1, Status: constant.ProductStatusEnum_Open},
		{Id: "4", Name: "product1", Amount: 10, Quantity: 8, Image: "image1", Inventory: 4, Status: constant.ProductStatusEnum_Closed},
	}
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().GetList(s.ctx, mock.Anything, mock.Anything).Return(poChart, nil)
	s.in.ProductCommon.(*mock_common_product.MockProductCommon).EXPECT().GetProduct(s.ctx).Return(product, nil)
	charts, err := s.memberChartUseCase.GetMemberChart(s.ctx)
	s.Assert().ErrorIs(nil, err)
	s.Assert().Equal(wantCharts, charts)
}
