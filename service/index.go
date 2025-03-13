package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// IndexHandler 入口函数
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{
		Code:     200,
		ErrorMsg: "succ",
		Data:     "test",
	}
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
