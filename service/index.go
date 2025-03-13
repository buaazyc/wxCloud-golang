package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"git.code.oa.com/trpc-go/trpc-go/log"
)

// Rsp 返回结构
type Rsp struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// IndexHandler 入口函数
func IndexHandler(rsp http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	log.DebugContextf(ctx, "req: %+v", req)
	res := &Rsp{
		Code:     200,
		ErrorMsg: "succ",
		Data:     "test",
	}
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(rsp, "内部错误")
		return
	}
	rsp.Header().Set("content-type", "application/json")
	rsp.Write(msg)
}
