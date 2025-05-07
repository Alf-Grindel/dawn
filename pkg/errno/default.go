package errno

const (
	SuccessCode = 0
	SuccessMsg  = "ok"

	ParamsErrCode    = 40000
	NotLoginErrCode  = 40100
	NoAuthErrCode    = 40101
	NotFoundErrCode  = 40400
	ForbiddenErrCode = 40300
	SystemErrCode    = 50000
	OperationErrCode = 500001
)

var (
	Success = NewErrno(SuccessCode, SuccessMsg)

	ParamsErr    = NewErrno(ParamsErrCode, "请求参数错误")
	NotLoginErr  = NewErrno(NotLoginErrCode, "未登录")
	NoAuthErr    = NewErrno(NoAuthErrCode, "无权限")
	NotFoundErr  = NewErrno(NotFoundErrCode, "请求数据不存在")
	ForbiddenErr = NewErrno(ForbiddenErrCode, "禁止访问")
	SystemErr    = NewErrno(SystemErrCode, "系统内部异常")
	OperationErr = NewErrno(OperationErrCode, "操作失败")
)
