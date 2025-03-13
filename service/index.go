package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Rsp 返回结构
type Rsp struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// IndexHandler 入口函数
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	res := &Rsp{
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
