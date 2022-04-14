package Controller

import (
	"FusionGO/Utils"
	"fmt"
	"net/http"
)

func handleIterceptor(Method string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if Method != r.Method {
			fmt.Fprintln(w, Utils.Message("请求方式不正确！", 500))
			return
		}
		h(w, r)
	}
}
