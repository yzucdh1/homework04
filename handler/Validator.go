package handler

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GetValidMessage(err error, obj any) string {
	// obj为结构体指针
	getObj := reflect.TypeOf(obj)
	// 断言为具体的类型，err是一个接口
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg") //错误信息不需要全部返回，当找到第一个错误的信息时，就可以结束
			}
		}
	}
	return err.Error()
}
