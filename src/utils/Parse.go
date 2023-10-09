package utils

import (
	"reflect"
	"strconv"
	"strings"
)

// SetStructFieldByJsonName 将hgetall出来的map[string]string转换成对应的结构体
func SetStructFieldByJsonName(ptr interface{}, fields map[string]string) {

	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {

		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("json")

		//去掉逗号后面内容 如 `json:"voucher_usage,omitempty"`
		name = strings.Split(name, ",")[0]

		if value, ok := fields[name]; ok {

			//给结构体赋值
			//保证赋值时数据类型一致
			//fmt.Println("类型1：", reflect.ValueOf(value).Type(), "类型2：", v.FieldByName(fieldInfo.Name).Type())
			if reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			} else if v.FieldByName(fieldInfo.Name).Type().String() == "uint64" {
				iv, _ := strconv.Atoi(value)
				nv := uint64(iv)
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(nv))
			} else if v.FieldByName(fieldInfo.Name).Type().String() == "float64" {
				fv, _ := strconv.ParseFloat(value, 64)
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(fv))
			} else if v.FieldByName(fieldInfo.Name).Type().String() == "int32" {
				iv, _ := strconv.Atoi(value)
				nv := int32(iv)
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(nv))
			} else if v.FieldByName(fieldInfo.Name).Type().String() == "uint32" {
				iv, _ := strconv.Atoi(value)
				nv := uint32(iv)
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(nv))
			} else if v.FieldByName(fieldInfo.Name).Type().String() == "int" {
				iv, _ := strconv.Atoi(value)
				nv := iv
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(nv))
			}

		}
	}

}
