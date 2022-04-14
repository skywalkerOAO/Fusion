package Services

import (
	"FusionGO/Utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func Submit(w http.ResponseWriter, r *http.Request) {
	type getJson struct {
		Name    string
		Batch   string
		Comenum int
		Number  string
	}

	var parmas getJson
	// 检查前端参数是否合规
	// JSON解码并判断格式是否正确
	if err := json.NewDecoder(r.Body).Decode(&parmas); err != nil {
		fmt.Fprintln(w, Utils.Message("参数类型错误", 500))
		return
	}
	// 判断参数是否完全
	if Utils.CheckJsonItemIsNull(parmas, true) {
		fmt.Fprintln(w, Utils.Message("参数不全", 504))
		return
	}
}
