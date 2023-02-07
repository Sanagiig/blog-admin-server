package utils

import (
	"errors"
	"reflect"
)

// ExtracFieds
//
//	@Description: 取出结构体所有字段，包含内嵌结构
//	@param data
//	@param res
func ExtracFieds(data any, res *[]string) {
	if res == nil {
		res = &[]string{}
	}
	var dt reflect.Type = reflect.TypeOf(data)
	var dv reflect.Value

	switch {
	case dt.Kind() == reflect.Pointer:
		dv = reflect.ValueOf(data).Elem()
		dt = reflect.TypeOf(dv.Interface())
	default:
		dv = reflect.ValueOf(data)
	}

	fieldNum := dt.NumField()
	for i := 0; i < fieldNum; i++ {
		if dt.Field(i).Type.Kind() == reflect.Struct {
			if dv.Field(i).CanAddr() {
				ExtracFieds(dv.Field(i).Interface(), res)
			}
		} else {
			*res = append(*res, dt.Field(i).Name)
		}
	}
}

func ExtracNotNullField() {

}

// FieldValTransfer
//
//	@Description: 将 src 的字段值复制到 dest
//	@param dest
//	@param src
//	@return error
func FieldValTransfer(dest any, src any) error {
	dv := reflect.ValueOf(dest)
	dt := reflect.TypeOf(dest)
	sv := reflect.ValueOf(src)
	st := reflect.TypeOf(src)

	if dv.Kind() == reflect.Pointer {
		dv = dv.Elem()
		dt = reflect.TypeOf(dv.Interface())
	}

	if sv.Kind() == reflect.Pointer {
		sv = sv.Elem()
		st = reflect.TypeOf(sv.Interface())
	}

	if dv.Kind() != reflect.Struct || sv.Kind() != reflect.Struct {
		return errors.New("不能对非结构体进行字段转移")
	}

	var destFieldNameList []string
	ExtracFieds(dest, &destFieldNameList)
	for i := 0; i < len(destFieldNameList); i++ {
		fname := destFieldNameList[i]
		df, dfOk := dt.FieldByName(fname)
		sf, sfOk := st.FieldByName(fname)
		if dfOk && sfOk && df.Type == sf.Type {

			if dfv := dv.FieldByName(fname); dfv.CanSet() {
				dfv.Set(sv.FieldByName(fname))
			}
		}
	}

	return nil
}
