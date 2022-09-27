package session

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/model/bo"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/repository/mock_repository"
	"simon/mall/service/internal/thirdparty/mysqlcli"
	"simon/mall/service/internal/utils/ctxs"
	"simon/mall/service/internal/utils/uuid/mock_uuid"
)

type sessionSuit struct {
	suite.Suite

	ctx context.Context
	*sessionUseCase
}

func TestSession(t *testing.T) {
	suite.Run(t, &sessionSuit{})
}

// 測試初始化
func (s *sessionSuit) SetupSuite() {
	gin.SetMode(gin.TestMode)
	s.sessionUseCase = &sessionUseCase{}
}

const (
	memberDefaultId       = "111"
	memberDefaultAccount  = "simon"
	memberDefaultPassword = "simon123"
	memberDefaultName     = "simon"
	memberDefaultToken    = "222"
	salt                  = "92c469c3-1b59-4c70-8607-59d47cb56bd1"
	hashPasswd            = "da7f57067465be72534a9489c4194aa001806cc45b3fa0a1c9edc6fbce49bdf0"
)

func (s *sessionSuit) SetupTest() {
	s.ctx, _ = gin.CreateTestContext(httptest.NewRecorder())
	ctxs.SetSession(s.ctx.(*gin.Context), &bo.MemberSession{Id: memberDefaultId, Account: memberDefaultAccount, Name: memberDefaultName, Token: memberDefaultToken})
	s.in.DB = mysqlcli.NewMockClient()
	s.in.Uuid = mock_uuid.NewMockUuid(s.T())
	s.in.MemberRepo = mock_repository.NewMockMemberRepo(s.T())
	s.in.SessionRepo = mock_repository.NewMockSessionRepo(s.T())
}
func (s *sessionSuit) Test_Session_Login() {
	var err error
	var cond *bo.MemberSessionCond
	loginCtx := context.Background()

	//
	s.SetupTest()
	s.T().Log("login: first by account fail return error, check first by account parameter")
	cond = &bo.MemberSessionCond{Account: memberDefaultAccount, Password: memberDefaultPassword}
	s.in.MemberRepo.(*mock_repository.MockMemberRepo).EXPECT().FirstByAccount(loginCtx, mock.Anything, memberDefaultAccount).Return(nil, errs.MemberNoMatch)
	_, err = s.sessionUseCase.Login(loginCtx, cond)
	s.Assert().ErrorIs(errs.MemberNoMatch, err)

	mockMember := &po.Member{
		Id:       memberDefaultId,
		Account:  memberDefaultAccount,
		Name:     memberDefaultName,
		Password: hashPasswd,
		Salt:     salt,
	}

	//
	s.SetupTest()
	s.T().Log("login: first by account, password fail")
	cond = &bo.MemberSessionCond{Account: memberDefaultAccount, Password: ""}
	s.in.MemberRepo.(*mock_repository.MockMemberRepo).EXPECT().FirstByAccount(loginCtx, mock.Anything, memberDefaultAccount).Return(mockMember, nil)
	_, err = s.sessionUseCase.Login(loginCtx, cond)
	s.Assert().ErrorIs(errs.MemberNoMatch, err)

	//
	s.SetupTest()
	s.T().Log("login: password ok set fail, check parameters ")
	cond = &bo.MemberSessionCond{Account: memberDefaultAccount, Password: memberDefaultPassword}
	s.in.MemberRepo.(*mock_repository.MockMemberRepo).EXPECT().FirstByAccount(loginCtx, mock.Anything, memberDefaultAccount).Return(mockMember, nil)
	s.in.Uuid.(*mock_uuid.MockUuid).EXPECT().GetUUID().Return(memberDefaultToken)
	s.in.SessionRepo.(*mock_repository.MockSessionRepo).EXPECT().SetMemberLogin(loginCtx, &po.MemberSession{Id: memberDefaultId, Account: memberDefaultAccount, Name: memberDefaultName, Token: memberDefaultToken}).Return(errs.CommonUnknownError)
	_, err = s.sessionUseCase.Login(loginCtx, cond)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	mockToken := &bo.MemberToken{Token: memberDefaultToken}
	//
	s.SetupTest()
	s.T().Log("login: password ok set ok")
	cond = &bo.MemberSessionCond{Account: memberDefaultAccount, Password: memberDefaultPassword}
	s.in.MemberRepo.(*mock_repository.MockMemberRepo).EXPECT().FirstByAccount(loginCtx, mock.Anything, memberDefaultAccount).Return(mockMember, nil)
	s.in.Uuid.(*mock_uuid.MockUuid).EXPECT().GetUUID().Return(memberDefaultToken)
	s.in.SessionRepo.(*mock_repository.MockSessionRepo).EXPECT().SetMemberLogin(loginCtx, &po.MemberSession{Id: memberDefaultId, Account: memberDefaultAccount, Name: memberDefaultName, Token: memberDefaultToken}).Return(nil)
	token, err := s.sessionUseCase.Login(loginCtx, cond)
	s.Assert().ErrorIs(nil, err)
	s.Assert().Equal(mockToken, token)

}
func (s *sessionSuit) Test_Session_Logout() {
	var err error

	//
	s.SetupTest()
	s.T().Log("logout session fail")
	s.ctx = context.Background()
	err = s.sessionUseCase.Logout(s.ctx)
	s.Assert().ErrorIs(errs.MemberTokenError, err)

	//
	s.SetupTest()
	s.T().Log("login: check session return error, check  parameter")
	s.in.SessionRepo.(*mock_repository.MockSessionRepo).EXPECT().CheckSessionExist(s.ctx, memberDefaultToken).Return(false, errs.CommonUnknownError)
	err = s.sessionUseCase.Logout(s.ctx)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("login: check session return false no error")
	s.in.SessionRepo.(*mock_repository.MockSessionRepo).EXPECT().CheckSessionExist(s.ctx, memberDefaultToken).Return(false, nil)
	err = s.sessionUseCase.Logout(s.ctx)
	s.Assert().ErrorIs(errs.MemberTokenError, err)

	//
	s.SetupTest()
	s.T().Log("login: remove session return error")
	s.in.SessionRepo.(*mock_repository.MockSessionRepo).EXPECT().CheckSessionExist(s.ctx, memberDefaultToken).Return(true, nil)
	s.in.SessionRepo.(*mock_repository.MockSessionRepo).EXPECT().RemoveUserLogin(s.ctx, memberDefaultToken).Return(errs.CommonUnknownError)
	err = s.sessionUseCase.Logout(s.ctx)
	s.Assert().ErrorIs(errs.CommonUnknownError, err)

	//
	s.SetupTest()
	s.T().Log("login: remove session")
	s.in.SessionRepo.(*mock_repository.MockSessionRepo).EXPECT().CheckSessionExist(s.ctx, memberDefaultToken).Return(true, nil)
	s.in.SessionRepo.(*mock_repository.MockSessionRepo).EXPECT().RemoveUserLogin(s.ctx, memberDefaultToken).Return(nil)
	err = s.sessionUseCase.Logout(s.ctx)
	s.Assert().ErrorIs(nil, err)

}
