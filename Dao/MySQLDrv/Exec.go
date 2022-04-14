package MySQLDrv

import (
	"FusionGO/Utils"
	"database/sql"
	"fmt"
)

func Exec(parma string) bool {

	con, err := sql.Open("mysql", *connString)
	_, err = con.Exec(parma)
	if err != nil {
		fmt.Println("错误发生时间为：" + Utils.TimeStr())
		fmt.Println("错误语句为：" + parma)
		fmt.Println("数据库返回错误为：" + err.Error())
		return false
	}
	defer con.Close()

	return true
}
func Execs(parma1 string, parma2 string) bool {
	db, err := sql.Open("mysql", *connString)
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
