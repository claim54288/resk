package resk

import (
	"resk/apis/gorpc"
	_ "resk/apis/gorpc"
	_ "resk/apis/web"
	_ "resk/core/accounts"
	_ "resk/core/envelopes"
	"resk/infra"
	"resk/infra/base"
)

//进度：6-12

func init() {
	infra.Register(&base.PropsStarter{})       //配置
	infra.Register(&base.DbxDatabaseStarter{}) //数据库
	infra.Register(&base.ValidatorStarter{})   //验证器
	infra.Register(&base.GoRPCStarter{})       //RPC
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&base.IrisServerStarter{}) //网络
	infra.Register(&infra.WebApiSterter{})    //web接口
}
