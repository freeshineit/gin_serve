package helper

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func removeMapKey(fields map[string]string) []string {
	// "LoginForm.user": "user长度不能超过10个字符" ->"user长度不能超过10个字符"
	// 提取user,去掉Key --> "user长度不能超过10个字符"
	rsp := []string{}
	for _, err := range fields {
		rsp = append(rsp, err)
	}
	return rsp
}

func InitTrans(locale string) (err error) {
	// 修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json的tag的自定义方法
		// "LoginForm.User": "User长度不能超过10个字符"
		// 将大写的User替换为json中定义的tag标签 -- "LoginForm.user": "user长度不能超过10个字符"
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			en_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func ParseBindingError(err error) string {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return "Parse Binding Error"
	}
	return strings.Join(removeMapKey(errs.Translate(trans)), "\n")
}
