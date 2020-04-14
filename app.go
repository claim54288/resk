package resk

import (
	_ "resk/apis/web"
	_ "resk/core/accounts"
	"resk/infra"
	"resk/infra/base"
)

//进度：6-12

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiSterter{})
}
