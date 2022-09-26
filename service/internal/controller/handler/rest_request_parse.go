package handler

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/xerrors"

	"mall/service/internal/errs"
	"mall/service/internal/model/dto"
)

type IRequestParse interface {
	Bind(ctx *gin.Context, obj interface{}) error
	ParamInt64(ctx *gin.Context, key string) (int64, error)
}

type requestParseHandler struct {
	in digIn
}

func (r requestParseHandler) Bind(ctx *gin.Context, obj interface{}) error {
	switch ctx.Request.Method {
	case http.MethodGet:
		return r.bindMethodGet(ctx, obj)
	default:
		if err := binding.JSON.Bind(ctx.Request, obj); err != nil {
			r.in.SysLogger.Error(ctx, err)
			return xerrors.Errorf("%s:%w", "requestParseHandler.Bind", err)
		}
	}

	return nil
}

func (r requestParseHandler) bindMethodGet(ctx *gin.Context, obj interface{}) error {
	if err := binding.Query.Bind(ctx.Request, obj); err != nil {
		r.in.SysLogger.Error(ctx, err)
		return xerrors.Errorf("%s:%w", "requestParseHandler.bindGinMethodGet", err)
	}

	var pager *dto.PagerReq
	valPager := reflect.ValueOf(&pager).Elem()

	// 當發現有Pager 時，解析並塞入pager
	val := reflect.Indirect(reflect.ValueOf(obj).Elem())
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Type() == valPager.Type() {
			if field.IsNil() {
				p := &dto.PagerReq{Index: 0, Size: 0}
				field.Set(reflect.ValueOf(&p).Elem())
			}
			break
		}
	}

	return nil
}

func (r requestParseHandler) ParamInt64(ctx *gin.Context, key string) (int64, error) {
	val := ctx.Param(key)
	result, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		// todo error tool
		return 0, xerrors.Errorf("ParamInt64 %w", errs.RequestParamParseFailed)
	}

	return result, err
}
