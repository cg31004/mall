package repository

import (
	"context"
	"time"

	"golang.org/x/xerrors"

	"simon/mall/service/internal/constant"
	"simon/mall/service/internal/model/po"
)

//go:generate mockery --name ISessionRepo --structname MockSessionRepo --output mock_repository --outpkg mock_repository --filename mock_session.go --with-expecter

//TODO 可以抽換成 redis
type ISessionRepo interface {
	SetMemberLogin(ctx context.Context, session *po.MemberSession) error

	CheckTokenExist(ctx context.Context, memberId string) (bool, error)
	CheckSessionExist(ctx context.Context, token string) (bool, error)

	GetTokenById(ctx context.Context, memberId string) (string, error)
	GetSessionByToken(ctx context.Context, token string) (*po.MemberSession, error)

	RemoveUserLogin(ctx context.Context, token string) error
}

func newSessionRepoByRedis(in repositoryIn) ISessionRepo {
	return &sessionRepoByRedis{in: in}
}

type sessionRepoByRedis struct {
	in repositoryIn
}

func (dao *sessionRepoByRedis) SetMemberLogin(ctx context.Context, session *po.MemberSession) error {
	isNeedKick, err := dao.CheckTokenExist(ctx, session.Id)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	// if login kick login
	if isNeedKick {
		token, err := dao.GetTokenById(ctx, session.Id)
		if err != nil {
			return xerrors.Errorf("%w", err)
		}

		err = dao.RemoveUserLogin(ctx, token)
		if err != nil {
			return xerrors.Errorf("%w", err)
		}
	}

	// uid -> token
	dao.in.LocalCache.SaveWithExpiration(dao.keyByMemberId(session.Id), session.Token, time.Duration(dao.in.AppConf.GetLocalCacheConfig().SessionExpirationSec)*time.Second)
	// token -> session(name id account ...
	dao.in.LocalCache.SaveWithExpiration(dao.keyByToken(session.Token), session, time.Duration(dao.in.AppConf.GetLocalCacheConfig().SessionExpirationSec)*time.Second)
	return nil
}

func (dao *sessionRepoByRedis) CheckTokenExist(ctx context.Context, userId string) (bool, error) {
	_, ok := dao.in.LocalCache.Get(dao.keyByMemberId(userId))
	if !ok {
		return false, nil
	}
	return true, nil
}

func (dao *sessionRepoByRedis) CheckSessionExist(ctx context.Context, token string) (bool, error) {
	_, ok := dao.in.LocalCache.Get(dao.keyByToken(token))
	if !ok {
		return false, nil
	}
	return true, nil
}

func (dao *sessionRepoByRedis) GetTokenById(ctx context.Context, memberId string) (string, error) {
	val, ok := dao.in.LocalCache.Get(dao.keyByMemberId(memberId))
	if !ok {
		return "", xerrors.New("not match data")
	}

	token, ok := val.(string)
	if !ok {
		return "", xerrors.New("no match data")
	}

	return token, nil
}

func (dao *sessionRepoByRedis) GetSessionByToken(ctx context.Context, token string) (*po.MemberSession, error) {
	val, ok := dao.in.LocalCache.Get(dao.keyByToken(token))
	if !ok {
		return nil, xerrors.New("not match data")
	}

	session, ok := val.(po.MemberSession)
	if !ok {
		return nil, xerrors.New("no match data")
	}

	return &session, nil

}

func (dao *sessionRepoByRedis) RemoveUserLogin(ctx context.Context, token string) error {
	session, err := dao.GetSessionByToken(ctx, token)
	if err != nil {
		return xerrors.Errorf("%w", err)
	}

	dao.in.LocalCache.Delete(dao.keyByToken(session.Token))
	dao.in.LocalCache.Delete(dao.keyByMemberId(session.Id))
	return nil
}

func (dao *sessionRepoByRedis) keyByToken(token string) string {
	return constant.CacheSessionByToken + token
}

func (dao *sessionRepoByRedis) keyByMemberId(memberId string) string {
	return constant.CacheSessionByMemberId + memberId
}
