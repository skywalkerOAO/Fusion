package MySQLDrv

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var password = flag.String("password", "YourPassword", "the database password")
var port *int = flag.Int("port", 3306, "the database port")
var server = flag.String("server", "127.0.0.1", "the database server")
var user = flag.String("user", "Username", "the database user")
var database = flag.String("database", "Database", "the database name")
var connString = flag.String("connString", fmt.Sprintf("%s:%s@(%s:%d)/%s", *user, *password, *server, *port, *database), "connTotal")

// Conn 测试函数
func Conn() {

	con, err := sql.Open("mysql", *connString) // _写err并没有什么用
	if err != nil {
		log.Fatal("数据库打开失败,错误信息为:", err.Error())
		return
	}
	if con.Ping() != nil {
		fmt.Printf("数据库连接失败:")
		panic(con.Ping())
	}
	defer con.Close()
	fmt.Println("-----MYSQL连接成功-----")
	return
}
