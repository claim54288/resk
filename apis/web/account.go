package web

import (
	"github.com/kataras/iris"
	"resk/infra"
	"resk/infra/base"
	"resk/services"
)

//定义web api的时候，对每一个子业务，定义统一的前缀
//资金账户的根路径定义为:/account
//版本号：/v1

const (
	ResCodeBizeTransferedFailure = base.ResCode(6010)
)

func init() {
	infra.RegisterApi(&AccountApi{})
}

type AccountApi struct{}

func (*AccountApi) Init() {
	groupRouter := base.Iris().Party("/v1/account")
	groupRouter.Post("/create", createHandler)
	groupRouter.Post("/transfer", transferHandler)
	groupRouter.Get("/envelope/get", getEnvelopeAccountHandler)
	groupRouter.Get("/get", getAccountHandler)
}

//账户创建的接口：/v1/accmount/createHandler
//POST body json
func createHandler(ctx iris.Context) {
	//获取请求参数
	account := services.AccountCreatedDTO{}
	err := ctx.ReadJSON(&account)
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		ctx.JSON(&r)
		return
	}
	//执行创建账户的代码
	service := services.GetAccountService()
	dto, err := service.CreateAccount(account)
	if err != nil {
		r.Code = base.ResCodeInnerServerError
		r.Message = err.Error()
	}
	r.Data = dto
	ctx.JSON(&r)
}

//转账的接口：Post /v1/accmount/transfer
func transferHandler(ctx iris.Context) {
	//获取请求参数
	account := services.AccountTransferDTO{}
	err := ctx.ReadJSON(&account)
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		ctx.JSON(&r)
		return
	}
	//执行转账
	service := services.GetAccountService()
	status, err := service.Transfer(account)
	if err != nil {
		r.Code = base.ResCodeInnerServerError
		r.Message = err.Error()
	}
	r.Data = status
	if status != services.TransferedStatusSuccess {
		r.Code = ResCodeBizeTransferedFailure
		r.Message = err.Error()
	}
	ctx.JSON(&r)
}

//查询红包账户的web接口: /v1/account/envelope/get
func getEnvelopeAccountHandler(ctx iris.Context) {
	userId := ctx.URLParam("userId")
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if userId == "" {
		r.Code = base.ResCodeRequestParamsError
		r.Message = "用户ID不能为空"
		ctx.JSON(&r)
		return
	}
	service := services.GetAccountService()
	account := service.GetEnvelopeAccountByUserId(userId)
	r.Data = account
	ctx.JSON(&r)
}

//查询账户信息的web接口：/v1/account/get
func getAccountHandler(ctx iris.Context) {
	accountNo := ctx.URLParam("accountNo")
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if accountNo == "" {
		r.Code = base.ResCodeRequestParamsError
		r.Message = "账户编号不能为空"
		ctx.JSON(&r)
		return
	}
	service := services.GetAccountService()
	account := service.GetAccount(accountNo)
	r.Data = account
	ctx.JSON(&r)
}
