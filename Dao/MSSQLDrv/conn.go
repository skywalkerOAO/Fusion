package MSSQLDrv

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

/*正式环境配置*/
//var mspassword = flag.String("password", "W991224z", "the database password")
//var msport *int = flag.Int("port", 1433, "the database port")
//var msserver = flag.String("server", "GWM", "the database server")
//var msuser = flag.String("user", "szb", "the database user")
//var msdatabase = flag.String("database", "LMS", "the database name")
//var msconnString = fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d;encrypt=disable", *msserver, *msuser, *mspassword, *msdatabase, *msport)

/*测试环境配置*/
var mspassword = flag.String("mspassword", "YourPassword", "the database password")
var msport *int = flag.Int("msport", 1433, "the database port")
var msserver = flag.String("msserver", "127.0.0.1\\YourDatabase-instance", "the database server")
var msuser = flag.String("msuser", "sa", "the database user")
var msdatabase = flag.String("msdatabase", "YourDatabase", "the database name")
var msconnString = flag.String("mssconnString", fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d;encrypt=disable", *msserver, *msuser, *mspassword, *msdatabase, *msport), "mssqlconn")

func Conn() {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d;encrypt=disable", *msserver, *msuser, *mspassword, *msdatabase, *msport)

	con, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("连接数据库驱动错误,错误信息为:", err.Error())
	}
	if con.Ping() != nil {
		fmt.Printf("数据库连接失败:%s\n", con.Ping())
	} else {
		fmt.Println("Connect DataBase Success!")
	}

	defer con.Close()
}
