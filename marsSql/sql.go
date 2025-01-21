package marsSql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/marsli9945/mars-go/marsLog"
	"reflect"
	"strconv"
	"strings"
)

func ExecuteContext(ctx context.Context, db *sql.DB, query string, args ...any) error {
	_, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("execute failed: %w", err)
	}
	return nil
}

func SelectContext(ctx context.Context, db *sql.DB, fieldTag string, results any, sentence string, args ...any) error {
	query, err := db.QueryContext(ctx, sentence, args...)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	defer func(query *sql.Rows) {
		err = query.Close()
		if err != nil {
			marsLog.Logger().ErrorF("Close query error: %v", err)
			return
		}
	}(query)
	cols, err := query.Columns()
	if err != nil {
		return err
	}
	// 一行数据，使用any是为了避开数据类型的问题
	var rows = make([]any, len(cols))
	// 存实际的值，是byte数组，长度以列的数量为准
	var values = make([]any, len(cols))
	for i := 0; i < len(cols); i++ {
		rows[i] = &values[i]
	}
	var resultMap []map[string]any
	var resultItem map[string]any
	for query.Next() {
		if err = query.Scan(rows...); err != nil {
			return err
		}
		resultItem = map[string]any{}
		for i, v := range values {
			resultItem[cols[i]] = v
		}
		resultMap = append(resultMap, resultItem)
	}
	if resultMap != nil {
		return mapToAllSlice(fieldTag, resultMap, results)
	}
	return nil
}

func mapToAllSlice(fieldTag string, data []map[string]any, results any) error {
	// 获取目标结构体的类型和值
	resultsType := reflect.TypeOf(results)
	resultsValue := reflect.ValueOf(results)

	// 确保目标参数是切片类型
	if resultsType.Kind() != reflect.Ptr {
		return fmt.Errorf("results argument must be a pointer to a slice, but was a %s", resultsType.Kind())
	}

	sliceVal := resultsType.Elem()
	if sliceVal.Kind() == reflect.Interface {
		sliceVal = sliceVal.Elem()
	}

	if sliceVal.Kind() != reflect.Slice {
		return fmt.Errorf("results argument must be a pointer to a slice, but was a pointer to %s", sliceVal.Kind())
	}

	switch sliceVal.Elem().Kind() {
	case reflect.Map:
		mapToMapSlice(data, sliceVal, resultsValue)
	case reflect.Struct:
		mapToStructSlice(fieldTag, data, sliceVal, resultsValue)
	default:
		return fmt.Errorf("unsupported slice element type: %s", sliceVal.Kind())
	}
	return nil
}

func firstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func mapToStructSlice(fieldTag string, data []map[string]any, sliceVal reflect.Type, resultsValue reflect.Value) {
	// 遍历 data 中的每个 map
	for _, item := range data {
		// 创建一个新的目标结构体
		newElem := reflect.New(sliceVal.Elem()).Elem()
		for i := 0; i < sliceVal.Elem().NumField(); i++ {
			elemType := sliceVal.Elem().Field(i)
			field := newElem.FieldByName(elemType.Name)
			if field.IsValid() && field.CanSet() {
				fieldType := field.Type()
				tag := elemType.Tag.Get(fieldTag)
				if tag == "" {
					tag = firstLower(elemType.Name)
				}
				if itemData, ok := item[tag]; ok {
					if itemData == nil {
						continue
					}
					value := reflect.ValueOf(itemData)
					// 确保字段类型与值的类型相匹配
					if value.Type() == reflect.TypeOf([]uint8{}) {
						switch fieldType.Kind() {
						case reflect.Int, reflect.Int64, reflect.Int32:
							intValue, err := strconv.ParseInt(string(itemData.([]uint8)), 10, 64)
							if err != nil {
								continue
							}
							field.Set(reflect.ValueOf(intValue).Convert(fieldType))
						case reflect.Float32, reflect.Float64:
							floatValue, err := strconv.ParseFloat(string(itemData.([]uint8)), 10)
							if err != nil {
								continue
							}
							field.Set(reflect.ValueOf(floatValue).Convert(fieldType))
						case reflect.String:
							field.Set(reflect.ValueOf(string(itemData.([]uint8))).Convert(fieldType))
						default:
							marsLog.Logger().ErrorF("unhandled default case")
						}
					} else if value.Type().ConvertibleTo(fieldType) {
						field.Set(value.Convert(fieldType))
					}
				}

			}
		}
		// 将新的结构体添加到目标切片中
		resultsValue.Elem().Set(reflect.Append(resultsValue.Elem(), newElem))
	}
}

func mapToMapSlice(data []map[string]any, sliceVal reflect.Type, resultsValue reflect.Value) {
	// 创建用于存放结果的切片
	resultsSlice := reflect.MakeSlice(sliceVal, 0, len(data))

	// 遍历 data 中的每个 map
	for _, item := range data {
		// 创建一个新的 map
		newMap := reflect.MakeMap(sliceVal.Elem())

		// 遍历 map 中的字段，并将值设置到新的 map 中
		for fieldName, fieldValue := range item {
			fieldType := reflect.TypeOf(fieldValue)
			value := reflect.ValueOf(fieldValue)

			// 确保字段类型与值的类型相匹配
			if fieldValue != nil && value.Type().ConvertibleTo(fieldType) {
				// 将字段名称和值添加到新的 map 中
				newMap.SetMapIndex(reflect.ValueOf(fieldName), value.Convert(fieldType))
			}
		}

		// 将新的 map 添加到结果切片中
		resultsSlice = reflect.Append(resultsSlice, newMap)
	}

	// 将结果切片赋值给 results
	resultsValue.Elem().Set(resultsSlice)
}
