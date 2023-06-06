package helper

import (
	"reflect"
	"strings"
	"mango-user-center/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	trans    translator.Translator
)

func init() {
	uni := translator.New(en.New())
	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return "" + strings.ToLower(fld.Name) + ""
	})
	err := en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic("翻译器初始化失败 : " + err.Error())
	}
}

// 如果经过auth，直接拿，否则从token中解析
func GetUid(c *gin.Context) int {
	return c.GetInt("uid")
}

// 参数验证，务必在false时return。 若非强校验，使用c.ShouldBind(&req)
func Bind(c *gin.Context, to interface{}) bool {
	c.ShouldBind(to)
	err := validate.Struct(to)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			s := e.Translate(trans)
			response.Echo(c, nil, response.Msg(s))
			return false
		}
	}

	return true
}

// 尽量少使用trim
func BindAndTrim(c *gin.Context, to interface{}) bool {
	if !Bind(c, to) {
		return false
	}
	return true
}

func trimString(obj interface{}) {
	elem := reflect.Indirect(reflect.ValueOf(obj))
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		switch field.Kind() {
		case reflect.String:
			if field.String() != "" {
				field.SetString(strings.TrimSpace(field.String()))
			}
		case reflect.Struct:
			trimString(field.Addr().Interface())
		}
	}
}
