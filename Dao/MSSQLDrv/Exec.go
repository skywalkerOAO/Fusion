package MSSQLDrv

import (
	"FusionGO/Utils"
	"database/sql"
	"fmt"
)

// Exec
// Expaple:
// blabla
// blabla

func Exec(params string, args ...interface{}) bool {

	con, err := sql.Open("mssql", *msconnString)
	stmt, _ := con.Prepare(params)
	defer stmt.Close()
	res, err := stmt.Exec(args...)

	if err != nil {
		fmt.Println("错误发生时间为：" + Utils.TimeStr())
		fmt.Println("错误语句为：" + params)
		fmt.Println("数据库返回错误为：" + err.Error())
		return false
	}
	if res != nil {
		RowsAffected, err := res.RowsAffected()
		if err == nil {
			fmt.Println("操作：" + params)
			rows := fmt.Sprintln(RowsAffected, "行受影响")
			fmt.Println(rows)
		}
	}
	defer con.Close()

	return true
}
func Execs(parma1 string, parma2 string) bool {
	db, err := sql.Open("mysql", *msconnString)
	tx, _ := db.Begin()
	var flag1, flag2 int64 = 0, 0
	result1, _ := tx.Exec("")
	result2, _ := tx.Exec("")
	if result1 != nil {
		flag1, _ = result1.RowsAffected()
	}
	if result2 != nil {
		flag2, _ = result2.RowsAffected()
	}
	if flag1 == 1 && flag2 == 1 {
		//提交事务
		tx.Commit()
		return true
	} else {
		//回滚
		fmt.Println("错误发生时间为：" + Utils.TimeStr())
		fmt.Println("错误语句为：" + parma1)
		fmt.Println("错误语句为：" + parma2)
		fmt.Println("数据库返回错误为：" + err.Error())
		tx.Rollback()
		return false
	}
	return true
}
