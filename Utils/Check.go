package Utils

import (
	"net/http"
	"reflect"
)

// CheckJsonItemIsNull 判断json是否有空值 是返回true否返回false（仅判断string类型）由于struct自动填充0，flag如果为true则判断int变量是否为空和0
func CheckJsonItemIsNull(parmas interface{}, flag bool) bool {
	var valInfo = reflect.ValueOf(parmas)
	var valType = reflect.TypeOf(parmas)
	num := valInfo.NumField()
	for i := 0; i < num; i++ {
		valT := valType.Field(i).Type.String()
		val := valInfo.Field(i).Interface()
		if valT == "string" && val == "" {
			return true
		}
		if flag == true && valT == "int" && val == 0 {
			return true
		}

	}
	return false
}

// CheckResultIsNull 判断sqlQuery结果是否为空 空为true 否则为false
func CheckResultIsNull(parmas []map[string]interface{}) bool {
	if len(parmas) == 0 {
		return true
	}
	return false
}

// IsGet 判断HTTP请求方式，不是Get返回true
func IsGet(r *http.Request) bool {
	if r.Method != "GET" {
		return true
	}
	return false
}

// IsPost 判断HTTP请求方式，不是Post返回true
func IsPost(r *http.Request) bool {
	if r.Method != "POST" {
		return true
	}
	return false
}
