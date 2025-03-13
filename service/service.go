package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/code"
	"wxcloudrun-golang/handler"
	"wxcloudrun-golang/model"
)

// Rsp 返回结构
type Rsp struct {
	Code     int         `json:"code"`               // 返回码
	ErrorMsg string      `json:"errorMsg,omitempty"` // 错误信息
	Data     interface{} `json:"data"`               // 返回数据
}

// IndexHandler 入口函数
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.CallBackMsg{}
	rsp := &Rsp{Code: code.Success, ErrorMsg: "ok"}
	defer func() {
		log.Printf("req :%+v rsp :%+v\n", req, rsp)
		msg, _ := json.Marshal(rsp)
		w.Header().Set("content-type", "application/json")
		w.Write(msg)
		if rsp.Code != code.Success {
			// todo: 回复错误消息到微信
		}
	}()

	// 解析请求
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		rsp.Code = code.ErrParseParams
		rsp.ErrorMsg = fmt.Sprintf("json decode failed with %+v", err)
		return
	}

	// 处理请求
	c := req.GetCmd()
	if c == "" {
		c = "help"
	}
	h := handler.GetHandler(c)
	if h == nil {
		rsp.Code = code.ErrParseParams
		rsp.ErrorMsg = fmt.Sprintf("cmd[%v] not found", c)
		return
	}
	if err := h.Handle(req); err != nil {
		rsp.Code = code.ErrHandle
		rsp.ErrorMsg = fmt.Sprintf("handler failed with %+v", err)
		return
	}
}
