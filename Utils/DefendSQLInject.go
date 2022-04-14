package Utils

import "regexp"

func FilteredSQLInject(to_match_str interface{}) bool {
	// 过滤 ‘
	// ORACLE 注解 --  /**/
	// 关键字过滤 update ,delete
	// 正则的字符串, 不能用 " " 因为" "里面的内容会转义
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|truncate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute|SELECT|UPDATE|INSERT|DELETE|AND|OR|INTO|CHAR|SUBSTR|ASCII|DECLARE|EXEC|COUNT|MASTER|DROP|EXECUTE|TRUNCATE)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		panic(err.Error())
		return true
	}
	return re.MatchString(to_match_str.(string))
}
