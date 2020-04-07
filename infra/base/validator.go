package base

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
	"resk/infra"
)

var validate *validator.Validate
var translator ut.Translator

func Validate() *validator.Validate {
	return validate
}

func Translate() ut.Translator {
	return translator
}

type ValidatorStarter struct {
	infra.BaseStarter
}

func (v *ValidatorStarter) Init(ctx infra.StarterContext) {
	validate := validator.New()
	//创建消息国际化通用翻译器
	cn := zh.New()
	uni := ut.New(cn, cn) //通过通用翻译器来转化，接收fallback和本地化翻译器 两个参数
	var found bool
	translator, found = uni.GetTranslator("zh")
	if found { //表示是否可以找到翻译器
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			logrus.Error(err)
		}
	} else {
		logrus.Error("Not found translator:zh")
	}
}
