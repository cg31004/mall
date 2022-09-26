package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mall/service/internal/errs"
	"mall/service/internal/model/bo"
	"mall/service/internal/utils/ctxs"
	"mall/service/internal/utils/timelogger"
)

type IMemberAuthMiddleware interface {
	Authentication(ctx *gin.Context)
}

type userAuthMiddleware struct {
	in digIn
}

func (u *userAuthMiddleware) Authentication(ctx *gin.Context) {
	defer timelogger.LogTime(ctx)()

	func() {
		defer timelogger.LogTime(ctx)()

		token := ctx.GetHeader("Authorization")
		cond := &bo.MemberToken{
			Token: token,
		}

		session, err := u.in.SessionUseCase.AuthMember(ctx, cond)
		if err == errs.MemberTokenError {
			SetResp(ctx, RespFormat_Standard, http.StatusUnauthorized, "0", errs.MemberTokenError)
			ctx.Abort()
			return
		} else if err != nil {
			SetResp(ctx, RespFormat_Standard, http.StatusInternalServerError, "0", err)
			ctx.Abort()
			return
		}

		ctxs.SetSession(ctx, session)
	}()

	ctx.Next()
}
