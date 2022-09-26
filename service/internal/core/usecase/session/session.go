package session

import (
	"context"

	"github.com/jinzhu/copier"
	"golang.org/x/xerrors"

	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/model/bo"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/utils/hash"
	"simon/mall/service/internal/utils/timelogger"
	"simon/mall/service/internal/utils/uuid"
)

type ISessionUseCase interface {
	Login(ctx context.Context, cond *bo.MemberSessionCond) (*bo.MemberToken, error)
	Logout(ctx context.Context, token *bo.MemberToken) error
	AuthMember(ctx context.Context, token *bo.MemberToken) (*bo.MemberSession, error)
}

func newSessionUseCase(in digIn) ISessionUseCase {
	return &sessionUseCase{in: in}
}

type sessionUseCase struct {
	in digIn
}

// Login :wrap doLogin for log login
func (uc *sessionUseCase) Login(ctx context.Context, cond *bo.MemberSessionCond) (*bo.MemberToken, error) {
	defer timelogger.LogTime(ctx)()

	db := uc.in.DB.Session()

	// get member by account
	member, err := uc.in.MemberRepo.FirstByAccount(ctx, db, cond.Account)
	if err != nil {
		return nil, xerrors.Errorf("sessionUseCase.Login -> MemberRepo.FirstByAccount: %w", errs.MemberNoMatch)
	}

	// attempt password
	hashPasswd, err := hash.GetPasswordHash(cond.Password, member.Salt)
	if err != nil {
		return nil, xerrors.Errorf("sessionUseCase.checkUserPassword -> hash.GetPasswordHash: %w", err)
	}
	if hashPasswd != member.Password {
		return nil, errs.MemberNoMatch
	}

	token := uuid.GetUUID()
	// set login session
	session := &po.MemberSession{
		Id:      member.Id,
		Account: member.Account,
		Name:    member.Name,
		Token:   token,
	}
	if err = uc.in.SessionRepo.SetMemberLogin(ctx, session); err != nil {
		return nil, xerrors.Errorf("sessionUseCase.Login -> UserAuthDao.SetMemberLogin: %w", err)
	}

	return &bo.MemberToken{Token: token}, nil
}

func (uc *sessionUseCase) Logout(ctx context.Context, token *bo.MemberToken) error {
	exist, err := uc.in.SessionRepo.CheckSessionExist(ctx, token.Token)
	if err != nil {
		return xerrors.Errorf("sessionUseCase.Logout -> SessionRepo.CheckSessionExist: %w", err)
	}

	if !exist {
		return errs.MemberTokenError
	}

	if err = uc.in.SessionRepo.RemoveUserLogin(ctx, token.Token); err != nil {
		return xerrors.Errorf("sessionUseCase.Logout -> SessionRepo.RemoveUserLogin: %w", err)
	}

	return nil
}

func (uc *sessionUseCase) AuthMember(ctx context.Context, token *bo.MemberToken) (*bo.MemberSession, error) {
	defer timelogger.LogTime(ctx)()

	// check if token been login
	exist, err := uc.in.SessionRepo.CheckSessionExist(ctx, token.Token)
	if err != nil {
		return nil, xerrors.Errorf("sessionUseCase.AuthMember -> SessionRepo.CheckSessionExist: %w", err)
	}

	if !exist {
		return nil, errs.MemberTokenError
	}

	session, err := uc.in.SessionRepo.GetSessionByToken(ctx, token.Token)
	if err != nil {
		return nil, xerrors.Errorf("sessionUseCase.AuthMember -> SessionRepo.GetSessionByToken: %w", err)
	}

	var boSession bo.MemberSession
	if err = copier.Copy(&boSession, session); err != nil {
		return nil, xerrors.Errorf("sessionUseCase.AuthMember -> copier.Copy: %w", err)
	}

	return &boSession, nil
}
