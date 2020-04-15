// sql工具
package sql_util

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func NewSqlUtil(db *sql.DB) *SqlUtil {
	return &SqlUtil{db: db}
}

type SqlUtil struct {
	db *sql.DB
}

func (s *SqlUtil) FetchMap(ctx context.Context, query string, args ...interface{}) (result []map[string]interface{}, err error) {
	fmt.Println(query, args)
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Printf("fetch close error: %v; query: $v; args: %v \n", err, query, args)
		}
	}()
	columnsType, err := rows.ColumnTypes()
	if err != nil {
		return
	}
	addrs := make([]interface{}, len(columnsType))
	for k, v := range columnsType {
		switch v.DatabaseTypeName() {
		case "VARCHAR", "CHAR", "TEST", "ENUM":
			var param string
			addrs[k] = &param
		case "INT", "BIGINT", "TINYINT":
			var param int64
			addrs[k] = &param
		default:
			var param interface{}
			addrs[k] = &param
		}
	}
	result = make([]map[string]interface{}, 0)
	for rows.Next() {
		if err := rows.Scan(addrs...); err != nil {
			fmt.Printf("scan error:%v \n", err)
			continue
		}
		var data map[string]interface{}
		d := &data
		t := reflect.TypeOf(d).Elem()
		v := reflect.ValueOf(d).Elem()
		m := reflect.MakeMap(t)
		for idx, column := range columnsType {
			m.SetMapIndex(reflect.ValueOf(column.Name()), reflect.ValueOf(addrs[idx]).Elem())
		}
		v.Set(m)
		result = append(result, data)
	}
	return
}

func (s *SqlUtil) FetchMapRow(ctx context.Context, query string, args ...interface{}) (result map[string]interface{}, err error) {
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Printf("fetch close error: %v; query: $v; args: %v \n", err, query, args)
		}
	}()
	columnsType, err := rows.ColumnTypes()
	if err != nil {
		return
	}
	addrs := make([]interface{}, len(columnsType))
	for k, v := range columnsType {
		switch v.DatabaseTypeName() {
		case "VARCHAR", "CHAR", "TEST", "ENUM":
			var param string
			addrs[k] = &param
		case "INT", "BIGINT", "TINYINT":
			var param int64
			addrs[k] = &param
		default:
			var param interface{}
			addrs[k] = &param
		}
	}
	for rows.Next() {
		if err := rows.Scan(addrs...); err != nil {
			fmt.Printf("scan error:%v \n", err)
			continue
		}
		d := &result
		t := reflect.TypeOf(d).Elem()
		v := reflect.ValueOf(d).Elem()
		m := reflect.MakeMap(t)
		for idx, column := range columnsType {
			m.SetMapIndex(reflect.ValueOf(column.Name()), reflect.ValueOf(addrs[idx]).Elem())
		}
		v.Set(m)
		break
	}
	if result == nil {
		err = errors.New("未查询到数据")
	}
	return
}

func (s *SqlUtil) Fetch(ctx context.Context, query string, data interface{}, args ...interface{}) (err error) {
	data_type := reflect.TypeOf(data)
	data_value := reflect.ValueOf(data)
	if data_type.Kind() == reflect.Ptr {
		data_type = data_type.Elem()
	}
	if data_value.Kind() == reflect.Ptr {
		data_value = data_value.Elem()
	}
	if !data_value.CanAddr() { // Elem返回v持有的接口保管的值的Value封装，使用elem的前提是v必须是ptr或interface类型；CanAddr是否可以获取v持有值的指针。可以获取指针的值被称为可寻址的。
		err = errors.New("引用数据内部必须是可寻址的")
		return
	}
	if data_type.Kind() == reflect.Slice && data_type.Elem().Kind() == reflect.Ptr {
		err = errors.New("当data是切片类型时，内部数据不能是指针")
		return
	}
	resultMap, err := s.FetchMap(ctx, query, args...)
	if err != nil {
		return
	}
	children := data_type
	if data_type.Kind() == reflect.Slice {
		children = data_type.Elem()
	}
	if children.Kind() == reflect.Ptr {
		children = children.Elem()
	}
	//list := reflect.MakeSlice(data_type, 0, 0)
	for _, v := range resultMap {
		description := reflect.New(children)
		if description.Kind() == reflect.Ptr {
			description = description.Elem()
		}
		for i := 0; i < children.NumField(); i++ {
			value_children_field := children.Field(i)
			if value_children_field.Anonymous {
				continue
			}
			key := strings.Split(value_children_field.Tag.Get("sql"), ",")[0]
			for kk, vv := range v { // 如果字段tag在查询字段里面，加入值
				if key == kk {
					description.Field(i).Set(reflect.ValueOf(vv))
				}
			}

		}
		if data_type.Kind() == reflect.Slice {
			data_value.Set(reflect.Append(data_value, description))
		} else {
			data_value.Set(description)
		}
	}
	//if data_type.Kind() == reflect.Slice {
	//	data_value.Set(list)
	//}
	return
}

func (s *SqlUtil) FetchRow(ctx context.Context, query string, data interface{}, args ...interface{}) (err error) {
	data_type := reflect.TypeOf(data)
	if data_type.Kind() != reflect.Ptr {
		err = errors.New("引用值必须是指针类型")
		return
	}
	data_value := reflect.ValueOf(data) // 获取数据的reflect值
	if !data_value.Elem().CanAddr() {   // Elem返回v持有的接口保管的值的Value封装，使用elem的前提是v必须是ptr或interface类型；CanAddr是否可以获取v持有值的指针。可以获取指针的值被称为可寻址的。
		err = errors.New("引用值必须是可寻址的")
		return
	}

	resultMap, err := s.FetchMapRow(ctx, query, args...)
	if err != nil {
		return
	}

	//if data_value.Kind() == reflect.Ptr {
	//	data_value = reflect.New(data_type.Elem().Elem())
	//}
	data_type = data_type.Elem()   // Elem返回t持有的接口保管的值的Value封装
	data_value = data_value.Elem() // Elem返回v持有的接口保管的值的Value封装
	data_value_type := data_value.Type()
	for col, map_value := range resultMap {
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
			//for col, mapValue := range map_value { // 如果字段tag在查询字段里面，加入值
			if col == key {
				data_value_field.Set(reflect.ValueOf(map_value))
			}
			//}
		}
		//break
	}
	return
}

func (s *SqlUtil) Update(ctx context.Context, query string, args ...interface{}) (affect int64, err error) {
	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return
	}

	affect, err = res.RowsAffected()
	return
}

func (s *SqlUtil) Insert(ctx context.Context, query string, args ...interface{}) (lastID int64, err error) {
	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return
	}

	lastID, err = res.LastInsertId()
	return
}

func (s *SqlUtil) Delete(ctx context.Context, query string, args ...interface{}) (affect int64, err error) {
	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return
	}

	affect, err = res.RowsAffected()
	return
}

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
