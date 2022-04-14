package Conf

import (
	"flag"
	"fmt"
)

// 网络地址

const FILESRV string = "10.4.51.87"
const FILESRVPORT string = "5000"
const FILETOTALSRV string = "https://" + FILESRV + ":" + FILESRVPORT
const FILETOTALSRV_HTTP string = "http://" + FILESRV + ":" + FILESRVPORT

// 目录

const RUNTIME_DIR string = "E:/HYCET"
const FILE_TEMP_DIR string = RUNTIME_DIR + "/temp"
const FILE_LONG_DIR string = RUNTIME_DIR + "/long"
const FILE_HIS_DIR string = RUNTIME_DIR + "/history"
const ListenPort int = 8000

// 数据库 Database

var mspassword = flag.String("mspassword", "W991224z", "the database password")
var msport *int = flag.Int("msport", 1433, "the database port")
var msserver = flag.String("msserver", "10.4.51.87\\MYNEWSQL", "the database server")
var msuser = flag.String("msuser", "sa", "the database user")
var msdatabase = flag.String("msdatabase", "plan", "the database name")
var msconnString = flag.String("mssconnString", fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d;encrypt=disable", *msserver, *msuser, *mspassword, *msdatabase, *msport), "mssqlconn")
