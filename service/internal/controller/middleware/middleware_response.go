package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"simon/mall/service/internal/constant"
	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/thirdparty/errortool"
	"simon/mall/service/internal/utils/ctxs"
	"simon/mall/service/internal/utils/timelogger"
)

const (
	RespFormat_Standard = "RespFormat_Standard"

	Resp_Format      = "Resp_Format"
	Resp_Data        = "Resp_Data"
	Resp_Status      = "Resp_Status"
	Resp_MessageCode = "Resp_MessageCode"

	Resp_Success = "success"
)

type IResponseMiddleware interface {
	Handle(ctx *gin.Context)
}

type responseMiddleware struct {
	in digIn
}

func (m *responseMiddleware) Handle(ctx *gin.Context) {
	if m.in.SysLogger.Level() == "debug" {
		ctx.Set(timelogger.ContextKey, timelogger.NewTimeLogger())
	}

	ctx.Set(constant.App_ChainID, uuid.New().String())
	ctx.Set(ctxs.HEADER_CLIENT_IP, ctx.ClientIP())

	ctx.Next()

	switch ctx.GetString(Resp_Format) {
	case RespFormat_Standard:
		m.standardResponse(ctx)
	default:
	}
}

func (m *responseMiddleware) standardResponse(ctx *gin.Context) {
	resp := m.generateStandardResponse(ctx)

	if m.in.SysLogger.Level() == "debug" {
		processTime, err := timelogger.GetTotalDuration(ctx)
		if err != nil {
			m.in.SysLogger.Error(ctx, err)
		}

		resp.Meta.ProcessTime = processTime.String()

		timeLogs, err := timelogger.GetTimeLogs(ctx)
		if err != nil {
			m.in.SysLogger.Error(ctx, err)
		}

		resp.Meta.TimeLogs = timeLogs
	} else {
		if ctx.GetInt(Resp_Status) >= http.StatusBadRequest {
			resp.Data = nil
		}
	}

	ctx.JSON(
		ctx.GetInt(Resp_Status),
		resp,
	)
}

func (m *responseMiddleware) generateStandardResponse(ctx *gin.Context) response {
	data := ctx.MustGet(Resp_Data)
	code := ctx.GetString(Resp_MessageCode)
	message := ""

	if ctx.GetInt(Resp_Status) >= http.StatusBadRequest {
		if tmpErr, ok := data.(error); ok {
			if err, ok := errortool.Parse(tmpErr); ok {
				code = err.GetCode()
				message = err.GetMessage()
				data = err.GetMessage()
			} else {
				err, _ := errortool.Parse(errs.CommonUnknownError)
				code = err.GetCode()
				message = err.GetMessage()
				data = strings.Split(fmt.Sprintf("%+v", tmpErr), "\n")
			}

			m.in.AppLogger.Error(ctx, tmpErr)
		}
	}
	return response{
		Meta: responseMeta{
			RequestID: m.generateRequestID(ctx),
		},
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (m *responseMiddleware) generateRequestID(ctx *gin.Context) string {
	prefix := ctx.GetString(constant.App_ApiPrefixCode)
	if prefix != "" {
		prefix = prefix + "-"
	}
	return prefix + ctx.GetString(constant.App_ChainID)
}

// SetResp 傳入 Resp 內容
func SetResp(ctx *gin.Context, respFmt string, statusCode int, msgCode string, data interface{}) {
	ctx.Set(Resp_Format, respFmt)
	ctx.Set(Resp_Status, statusCode)
	ctx.Set(Resp_MessageCode, msgCode)
	ctx.Set(Resp_Data, data)
}

// Model

type response struct {
	Meta    responseMeta `json:"meta"`
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

type responseMeta struct {
	RequestID   string   `json:"requestID"`
	ProcessTime string   `json:"processTime,omitempty"`
	TimeLogs    []string `json:"timeLogs,omitempty"`
}
