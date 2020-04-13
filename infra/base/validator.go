package base

import (
	"gopkg.in/go-playground/validator.v9"
	"resk/infra"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	log "github.com/sirupsen/logrus"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

var validate *validator.Validate
var translator ut.Translator

func Validate() *validator.Validate {
	Check(validate)
	return validate
}

func Translate() ut.Translator {
	Check(translator)
	return translator
}

type ValidatorStarter struct {
	infra.BaseStarter
}

func (v *ValidatorStarter) Init(ctx infra.StarterContext) {
	validate = validator.New()
	//创建消息国际化通用翻译器
	cn := zh.New()
	uni := ut.New(cn, cn) //通过通用翻译器来转化，接收fallback和本地化翻译器 两个参数
	var found bool
	translator, found = uni.GetTranslator("zh")
	if found { //表示是否可以找到翻译器
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			log.Error(err)
		}
	} else {
		log.Error("Not found translator:zh")
	}
}

func ValidateStruct(s interface{}) (err error) {
	//验证
	err = Validate().Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Error(err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				log.Error(err.Translate(translator))
			}
		}
		return err
	}
	return nil
}
