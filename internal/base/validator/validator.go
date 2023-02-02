package validation

import (
	"bytes"
	"errors"
	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"reflect"
)

// MyValidate 自定义验证器
type MyValidate struct {
	Validate *validator.Validate
	Trans    ut.Translator
}

// GlobalValidate 全局验证器
var GlobalValidate MyValidate

func init() {
	eng := english.New()
	uni := ut.New(eng, eng)
	trans, ok := uni.GetTranslator("en")
	if !ok {
		panic(ok)
	}
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		jsonTag := fld.Tag.Get("json")
		if len(jsonTag) > 0 {
			return jsonTag
		}
		return fld.Tag.Get("form")
	})
	err := en.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
	GlobalValidate.Validate = validate
	GlobalValidate.Trans = trans
}

// Check 验证器通用验证方法
func (m *MyValidate) Check(value interface{}) error {
	// 首先使用validator进行验证
	err := m.Validate.Struct(value)
	errBuf := bytes.Buffer{}
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		// 几乎不会出现，除非验证器本身异常无法转换，以防万一就判断一下好了
		if !ok {
			return errors.New("validate check exception")
		}
		// 将所有的参数错误进行翻译然后拼装成字符串返回
		for i := 0; i < len(errs); i++ {
			errBuf.WriteString(errs[i].Translate(m.Trans) + " \n")
		}
	}
	// 如果它实现了CanCheck接口，就进行自定义验证
	if v, ok := value.(Checker); ok {
		errs := v.Check()
		for i := 0; i < len(errs); i++ {
			errBuf.WriteString(errs[i].Error() + " \n")
		}
	}
	if errBuf.Len() == 0 {
		return nil
	}
	// 删除掉最后一个空格和换行符
	errStr := errBuf.String()
	return errors.New(errStr[:len(errStr)-2])
}

// Checker 如果需要特殊校验，可以实现验证接口，或者通过自定义tag标签实现
type Checker interface {
	Check() []error
}
