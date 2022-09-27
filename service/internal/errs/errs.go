package errs

import "simon/mall/service/internal/thirdparty/errortool"

var (
	ConvertDB         = errortool.ConvertDB
	Parse             = errortool.Parse
	ConciseParseParse = errortool.ConciseParse
	Equal             = errortool.Equal
)

var (
	ErrDB = errortool.ErrDB
)

var (
	commGroup                = errortool.Codes.Group("01")
	CommonUnknownError       = commGroup.Error("001", "未知错误")
	CommonNoData             = commGroup.Error("002", "查无资料")
	CommonRawSQLNotFound     = commGroup.Error("003", "找不到执行档")
	CommonServiceUnavailable = commGroup.Error("004", "系统维护中")
	CommonConfigureInvalid   = commGroup.Error("005", "设置参数错误")
	CommonParseError         = commGroup.Error("006", "解析失败")
	CommonEsInsertZero       = commGroup.Error("007", "zero insert count")
)

var (
	requestGroup                  = errortool.Codes.Group("02")
	RequestParamInvalid           = requestGroup.Error("001", "请求参数错误")
	RequestParamParseFailed       = requestGroup.Error("002", "请求参数解析失败")
	RequestPageError              = requestGroup.Error("003", "请求的页数错误")
	RequestParseError             = requestGroup.Error("004", "解析失败")
	RequestParseTimeZoneError     = requestGroup.Error("005", "时区解析错误")
	RequestFrequentOperationError = requestGroup.Error("006", "频繁操作，请稍后再尝试")
)

var (
	fileServerGroup         = errortool.Codes.Group("03")
	FileServerUploadFailed  = fileServerGroup.Error("001", "图片上传失败")
	FileServerResponseNotOK = fileServerGroup.Error("002", "图片库异常")
)

var (
	dataConvertGroup = errortool.Codes.Group("04")
	DataConvertError = dataConvertGroup.Error("001", "格式转换错误")
)

var (
	memberGroup      = errortool.Codes.Group("05")
	MemberNoMatch    = memberGroup.Error("001", "登入失败，帐号或密码错误")
	MemberTokenError = memberGroup.Error("002", "请重新登入")
)

var (
	orderGroup          = errortool.Codes.Group("06")
	OrderProductNoMatch = orderGroup.Error("001", "商品已下架或沒有庫存，請更新訂單")
)
