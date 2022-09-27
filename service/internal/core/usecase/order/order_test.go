package order

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"simon/mall/service/internal/constant"
	"simon/mall/service/internal/core/common/product/mock_common_product"
	"simon/mall/service/internal/core/common/transaction/mock_common_txn"
	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/model/bo"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/repository/mock_repository"
	"simon/mall/service/internal/thirdparty/mysqlcli"
	"simon/mall/service/internal/utils/ctxs"
	"simon/mall/service/internal/utils/uuid/mock_uuid"
)

type orderSuite struct {
	suite.Suite

	ctx context.Context
	*orderUseCase
}

func TestOrder(t *testing.T) {
	suite.Run(t, &orderSuite{})
}

// 測試初始化
func (s *orderSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	s.orderUseCase = &orderUseCase{}
}

const (
	memberDefaultId      = "111"
	memberDefaultAccount = "simon"
	memberDefaultName    = "simon"
	memberDefaultToken   = "222"
	card1                = "111"
	card2                = "222"
)

func (s *orderSuite) SetupTest() {
	s.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
	ctxs.SetSession(s.ctx.(*gin.Context), &bo.MemberSession{Id: memberDefaultId, Account: memberDefaultAccount, Name: memberDefaultName, Token: memberDefaultToken})
	s.in.DB = mysqlcli.NewMockClient()
	s.in.Uuid = mock_uuid.NewMockUuid(s.T())

	s.in.TxnItemCommon = mock_common_txn.NewMockTxnItemCommon(s.T())
	s.in.ProductCommon = mock_common_product.NewMockProductCommon(s.T())

	s.in.MemberChartRepo = mock_repository.NewMockMemberChartRepo(s.T())
	s.in.TxnRepo = mock_repository.NewMockTxnRepo(s.T())
	s.in.TxnItemRepo = mock_repository.NewMockTxnItemRepo(s.T())
}

func (s *orderSuite) Test_Session_GetOrderList() {
	var err error
	memberId := memberDefaultId

	//
	s.SetupTest()
	s.T().Log("get order list session fail")
	noLogin := context.Background()
	_, err = s.orderUseCase.GetOrderList(noLogin)
	s.Assert().ErrorIs(errs.MemberTokenError, err)

	//
	s.SetupTest()
	s.T().Log("txn get fail")
	s.in.TxnRepo.(*mock_repository.MockTxnRepo).EXPECT().GetList(s.ctx, mock.Anything, &po.TransactionSearch{MemberId: &memberId}).Return(nil, errs.CommonUnknownError)
	_, err = s.orderUseCase.GetOrderList(s.ctx)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	mockTxn := []*po.Transaction{
		{Id: "1", MemberId: memberDefaultId, PaymentNumber: card1, Amount: 180, Status: constant.TransactionStatusEnum_Success},
		{Id: "2", MemberId: memberDefaultId, PaymentNumber: card2, Amount: 200, Status: constant.TransactionStatusEnum_Success},
	}
	txnItemCond := []string{"1", "2"}
	txnItem := map[string][]*bo.TxnItem{
		"1": {{Name: "txn1_name1", Amount: 30, Quantity: 3, Image: "txn1_image1"}, {Name: "txn1_name2", Amount: 2, Quantity: 2, Image: "txn1_image2"}, {Name: "txn1_name3", Amount: 1, Quantity: 6, Image: "txn1_image3"}},
		"2": {{Name: "txn2_name1", Amount: 1, Quantity: 100, Image: "txn2_image1"}, {Name: "txn2_name2", Amount: 4, Quantity: 5, Image: "txn2_image2"}, {Name: "txn2_name3", Amount: 20, Quantity: 4, Image: "txn2_image3"}},
	}

	txn1Item := []*bo.TxnItem{{Name: "txn1_name1", Amount: 30, Quantity: 3, Image: "txn1_image1"}, {Name: "txn1_name2", Amount: 2, Quantity: 2, Image: "txn1_image2"}, {Name: "txn1_name3", Amount: 1, Quantity: 6, Image: "txn1_image3"}}
	txn2Item := []*bo.TxnItem{{Name: "txn2_name1", Amount: 1, Quantity: 100, Image: "txn2_image1"}, {Name: "txn2_name2", Amount: 4, Quantity: 5, Image: "txn2_image2"}, {Name: "txn2_name3", Amount: 20, Quantity: 4, Image: "txn2_image3"}}
	wantTxn := []*bo.Txn{
		{Id: "1", PaymentNumber: card1, Amount: 180, Status: constant.TransactionStatusEnum_Success, Item: txn1Item},
		{Id: "2", PaymentNumber: card2, Amount: 200, Status: constant.TransactionStatusEnum_Success, Item: txn2Item},
	}

	//
	s.SetupTest()
	s.T().Log("txn item get fail")
	s.in.TxnRepo.(*mock_repository.MockTxnRepo).EXPECT().GetList(s.ctx, mock.Anything, &po.TransactionSearch{MemberId: &memberId}).Return(mockTxn, nil)
	s.in.TxnItemCommon.(*mock_common_txn.MockTxnItemCommon).EXPECT().GetTxnItem(s.ctx, &bo.GetTxnItemMapCond{MemberId: memberDefaultId, TxnId: txnItemCond}).Return(nil, errs.CommonUnknownError)
	_, err = s.orderUseCase.GetOrderList(s.ctx)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("get:combine txn")
	s.in.TxnRepo.(*mock_repository.MockTxnRepo).EXPECT().GetList(s.ctx, mock.Anything, &po.TransactionSearch{MemberId: &memberId}).Return(mockTxn, nil)
	s.in.TxnItemCommon.(*mock_common_txn.MockTxnItemCommon).EXPECT().GetTxnItem(s.ctx, &bo.GetTxnItemMapCond{MemberId: memberDefaultId, TxnId: txnItemCond}).Return(txnItem, nil)
	order, err := s.orderUseCase.GetOrderList(s.ctx)
	s.Assert().ErrorIs(nil, err)
	s.Assert().Equal(wantTxn, order)
}

func (s *orderSuite) Test_Session_CreateOrder() {
	var err error
	var cond *bo.CreateOrderCond

	//
	s.SetupTest()
	s.T().Log("get order list session fail")
	noLogin := context.Background()
	cond = &bo.CreateOrderCond{}
	err = s.orderUseCase.CreateOrder(noLogin, cond)
	s.Assert().ErrorIs(errs.MemberTokenError, err)

	//
	s.SetupTest()
	s.T().Log("get: list fail return error, check list parameter")
	cond = &bo.CreateOrderCond{}

	s.in.ProductCommon.(*mock_common_product.MockProductCommon).EXPECT().DeleteProductCache(s.ctx)
	s.in.TxnItemCommon.(*mock_common_txn.MockTxnItemCommon).EXPECT().DeleteTxnItem(s.ctx, &bo.DelTxnItemMapCond{MemberId: memberDefaultId})
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().GetList(s.ctx, mock.Anything, &po.MemberChartSearch{MemberId: memberDefaultId}).Return(nil, errs.CommonUnknownError)

	err = s.orderUseCase.CreateOrder(s.ctx, cond)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("get: prodcut common fail return error, check list parameter")
	cond = &bo.CreateOrderCond{}
	s.in.ProductCommon.(*mock_common_product.MockProductCommon).EXPECT().DeleteProductCache(s.ctx)
	s.in.TxnItemCommon.(*mock_common_txn.MockTxnItemCommon).EXPECT().DeleteTxnItem(s.ctx, &bo.DelTxnItemMapCond{MemberId: memberDefaultId})
	s.in.MemberChartRepo.(*mock_repository.MockMemberChartRepo).EXPECT().GetList(s.ctx, mock.Anything, mock.Anything).Return(nil, nil)
	s.in.ProductCommon.(*mock_common_product.MockProductCommon).EXPECT().GetProduct(s.ctx).Return(nil, errs.CommonUnknownError)

	err = s.orderUseCase.CreateOrder(s.ctx, cond)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)
}
