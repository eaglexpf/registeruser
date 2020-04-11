package sqlUtil

import (
	"errors"
	"reflect"
	"strings"
)

// 生成地址切片
func AddrsEncode(data interface{}, columns []string) ([]interface{}, error) {
	addrs := make([]interface{}, 0)
	data_type := reflect.TypeOf(data)
	if data_type.Kind() != reflect.Ptr {
		return addrs, errors.New("data必须是指针类型")
	}
	data_value := reflect.ValueOf(data) // 获取数据的reflect值
	if !data_value.Elem().CanAddr() {   // Elem返回v持有的接口保管的值的Value封装，使用elem的前提是v必须是ptr或interface类型；CanAddr是否可以获取v持有值的指针。可以获取指针的值被称为可寻址的。
		return addrs, errors.New("data必须是可寻址的")
	}
	data_type = data_type.Elem()   // Elem返回t持有的接口保管的值的Value封装
	data_value = data_value.Elem() // Elem返回v持有的接口保管的值的Value封装
	data_value_type := data_value.Type()
	for i := 0; i < data_value_type.NumField(); i++ {
		data_value_type_field := data_value_type.Field(i)
		data_value_field := data_value.Field(i)
		if data_value_type_field.Anonymous {
			continue
		}
		for data_value_field.Type().Kind() == reflect.Ptr {
			data_value_field = data_value_field.Elem()
		}
		key := strings.Split(data_value_type_field.Tag.Get("sql"), ",")[0]
		if key == "" {
			continue
		}
		for _, col := range columns { // 如果字段tag在查询字段里面，加入值
			if col == key {
				addrs = append(addrs, data_value_field.Addr().Interface()) // Addr返回一个持有指向v持有者的指针的Value封装；interface返回v当前持有的值
				break
			}
		}
	}
	if len(columns) != len(addrs) {
		return nil, errors.New("等待接收的字段必须包含所有的请求字段！")
	}
	return addrs, nil
}
