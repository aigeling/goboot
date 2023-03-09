package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger = zap.L().Sugar()

type GinBinder struct {
	vldt *validator.Validate
}

func NewGinBinder() *GinBinder {
	return &GinBinder{
		vldt: validator.New(),
	}
}

// 参数绑定
//
// 绑定body的json到对象o
//
// 参数：
//
// validate_fields:用于指定要验证的字段名,如果为空则不验证字段
func (v *GinBinder) BindBody(c *gin.Context, o any, validate_fields ...string) (err error) {

	err = c.ShouldBind(o)
	if err != nil {
		logger.Errorf("bind err:%v", err)
		return
	}
	if len(validate_fields) > 0 {
		err = v.vldt.StructPartial(o, validate_fields...)
		logger.Errorf("valid err:%v", err)
	}

	return
}

// 参数绑定
//
// 绑定body的json到对象o
//
// 参数：
//
// except_fields:用于指定不进行验证的字段名,如果为空则验证所有字段
func (v *GinBinder) BindBodyExcept(c *gin.Context, o any, except_fields ...string) (err error) {

	err = c.ShouldBind(o)
	if err != nil {
		return
	}
	if len(except_fields) == 0 {
		err = v.vldt.Struct(o)
	} else {
		err = v.vldt.StructExcept(o, except_fields...)
	}

	return
}

// 参数绑定
//
// 绑定uri的参数到对象o
//
// 参数：
//
// validate_fields:用于指定要验证的字段名,如果为空则验证所有字段
func (v *GinBinder) BindUri(c *gin.Context, o any, validate_fields ...string) (err error) {

	err = c.ShouldBindUri(o)
	if err != nil {
		return
	}

	if len(validate_fields) == 0 {
		err = v.vldt.Struct(o)
	} else {
		err = v.vldt.StructPartial(o, validate_fields...)
	}

	return
}

// 参数绑定
//
// 绑定uri的参数到对象o
//
// 参数：
//
// except_fields:用于指定不进行验证的字段名,如果为空则验证所有字段
func (v *GinBinder) BindUriExcept(c *gin.Context, o any, except_fields ...string) (err error) {

	err = c.ShouldBindUri(o)
	if err != nil {
		return
	}

	if len(except_fields) == 0 {
		err = v.vldt.Struct(o)
	} else {
		err = v.vldt.StructPartial(o, except_fields...)
	}

	return
}
