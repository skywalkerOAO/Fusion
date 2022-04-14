package Utils

import "encoding/json"

// Message 入参为MSG,CODE 返回值为json格式
func Message(Msg string, Code int) string {
	type Warn struct {
		Msg  string
		Code int
	}
	warn := &Warn{
		Msg:  Msg,
		Code: Code,
	}
	res, _ := json.Marshal(&warn)
	return string(res)
}
