package helper

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func InitTranslation(locale string) error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New()
		enT := en.New()

		uni := ut.New(enT, zhT, enT)

		Trans, ok = uni.GetTranslator(locale)

		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, Trans)
		default:
			en_translations.RegisterDefaultTranslations(v, Trans)
		}
		return nil
	}
	return nil
}

// func RemoveTopStruct(fileds map[string]string) map[string]string {

// }

// func FmtBindError(err error) (string, error) {

// 	errs, ok := err.(validator.ValidationErrors)

// 	if !ok {
// 		return "", errs
// 	}

// 	fileds := errs.Translate()
// 	rsp := map[string]string{}

// 	for filed, err := range fileds {
// 		rsp[filed[strings.Index(filed, ".")+1:]] = err
// 	}

// 	return rsp

// 	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 	// 	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
// 	// 		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

// 	// 		if name == "-" {
// 	// 			return ""
// 	// 		}
// 	// 		return name
// 	// 	})
// 	// }
// }
