package base

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
	"resk/infra"
	"sync"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
	logrus.Info("初始化配置.")
	GetSystemAccount()
}

type SystemAccount struct {
	AccountNo   string
	AccountName string
	UserId      string
	Username    string
}

var systemAccount *SystemAccount
var systemAccountOnce sync.Once

func GetSystemAccount() *SystemAccount {
	systemAccountOnce.Do(func() {
		systemAccount = new(SystemAccount)
		err := kvs.Unmarshal(Props(), systemAccount, "system.account")
		if err != nil {
			panic(err)
		}
	})
	return systemAccount
}

func GetEnvelopeActivityLink() string {
	return Props().GetDefault("envelope.link", "/v1/envelope/link")
}

func GetEnvelopeDomain() string {
	return Props().GetDefault("envelope.domain", "http://localhost")
}
