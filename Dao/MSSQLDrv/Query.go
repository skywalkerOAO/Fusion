package MSSQLDrv

import (
	"FusionGO/Utils"
	"database/sql"
	"fmt"
)

// Query 返回查询所有行
func Query(params string, args ...interface{}) []map[string]interface{} {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d;encrypt=disable", *msserver, *msuser, *mspassword, *msdatabase, *msport)
	con, _ := sql.Open("mssql", connString)
	stmt, _ := con.Prepare(params)
	defer stmt.Close()
	row, err := stmt.Query(args...)
	if err != nil {
		fmt.Println("错误发生时间为：" + Utils.TimeStr())
		fmt.Println("错误语句为：" + params)
		fmt.Println("数据库返回错误为：" + err.Error())
		return nil
	}
	tableData := make([]map[string]interface{}, 0)
	columns, err := row.Columns()
	if err != nil {
		return tableData
	}
	count := len(columns)
	if count == 0 {
		fmt.Println(tableData)
		return tableData
	}
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for row.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		row.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData
}
