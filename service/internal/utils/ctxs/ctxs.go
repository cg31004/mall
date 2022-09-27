package ctxs

import (
	"context"

	"github.com/gin-gonic/gin"

	"simon/mall/service/internal/model/bo"
)

const (
	USER_SESSION_KEY = "MemberSessionKey"
	HEADER_CLIENT_IP = "Header_Client_IP"
)

func SetSession(ctx *gin.Context, session *bo.MemberSession) {
	ctx.Set(USER_SESSION_KEY, session)
}

func GetSession(ctx context.Context) (*bo.MemberSession, bool) {
	val := ctx.Value(USER_SESSION_KEY)

	v, ok := val.(*bo.MemberSession)
	if !ok {
		return nil, false
	}

	return v, true
}
