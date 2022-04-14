package Services

import (
	"FusionGO/Dao/MSSQLDrv"
	"FusionGO/Utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPartName(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Data map[string]interface{}
	}
	sql1 := fmt.Sprintf("SELECT * FROM check_index")
	result1 := MSSQLDrv.Query(sql1)
	if Utils.CheckResultIsNull(result1) {
		fmt.Fprintln(w, Utils.Message("结果为空,请检查参数", 501))
		return
	}
	type Response struct {
		Code int
		Msg  string
		Data []map[string]interface{}
	}
	response := &Response{
		Data: result1,
		Code: 200,
		Msg:  "操作成功",
	}
	// 2.Json格式化
	result, err := json.Marshal(&response)

	if err != nil {
		fmt.Fprintln(w, Utils.Message("错误", 500))
	}
	// 3.打印
	fmt.Fprintln(w, string(result))
}
