package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130"` //gte的意思是greater than or equals 大于等于
	Email     string `validate:"required,email"`
}

func main() {
	user := &User{
		FirstName: "firstName",
		LastName:  "lastName",
		Age:       24,
		Email:     "aaa@qq.com",
	}
	user.Age = 136
	user.Email = "aaqq.com"
	validate := validator.New()
	//创建消息国际化通用翻译器
	cn := zh.New()
	uni := ut.New(cn, cn) //通过通用翻译器来转化，接收fallback和本地化翻译器 两个参数
	transtator, found := uni.GetTranslator("zh")
	if found { //表示是否可以找到翻译器
		err := vtzh.RegisterDefaultTranslations(validate, transtator)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("没有找到翻译器")
	}

	err := validate.Struct(user)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			fmt.Println(err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				fmt.Println(err.Translate(transtator))
			}
		}
	}
}
