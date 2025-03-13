package service

import (
	"encoding/json"
	"net/http"

	"git.code.oa.com/trpc-go/trpc-go/log"
)

type Rsp struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

// IndexHandler 计数器接口
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rsp := &Rsp{
		Code:     0,
		ErrorMsg: "success",
		Data:     "hello world",
	}
	msg, err := json.Marshal(rsp)
	if err != nil {
		log.ErrorContextf(ctx, "json marshal err: %v", err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
