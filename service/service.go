package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/constant"
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
	rsp := &Rsp{Code: constant.Success, ErrorMsg: "ok"}
	w.Header().Set("content-type", "application/json")
	defer func() {
		fmt.Printf("req :%+v rsp :%+v\n", req, rsp)
		if rsp.Code != constant.Success {
			msg, _ := json.Marshal(rsp)
			w.Write(msg)
		}
	}()

	// 解析请求
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		rsp.Code = constant.ErrParseParams
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
		rsp.Code = constant.ErrParseParams
		rsp.ErrorMsg = fmt.Sprintf("cmd[%v] not found", c)
		return
	}
	msgRsp, err := h.Handle(req)
	if err != nil {
		rsp.Code = constant.ErrHandle
		rsp.ErrorMsg = fmt.Sprintf("handler failed with %+v", err)
		return
	}
	msg, _ := json.Marshal(msgRsp)
	w.Write(msg)
}
