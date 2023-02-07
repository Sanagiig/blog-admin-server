package utils

import (
	"errors"
	"go-blog/model"
	"reflect"
)

type transData struct {
	Type string
}

// 转换字典中类型的名称到数据 typeName 字段
func TransStrucField2Another(dicList []model.Dictionary, data any, typeName string) error {
	dv := reflect.ValueOf(data)
	if dv.Kind() != reflect.Slice {
		return errors.New("数据格式错误")
	}

	for i := 0; i < dv.Len(); i++ {
		for j := range dicList {
			val := dv.Index(i)
			if val.Kind() == reflect.Pointer {
				val = val.Elem()
			}

			if val.Kind() != reflect.Struct {
				continue
			}

			fv := val.FieldByName(typeName)
			if fv.String() == dicList[j].ID {
				fv.SetString(dicList[j].Name)
			}
		}
	}

	return nil
}
