package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/constant"
	"wxcloudrun-golang/handler"
	"wxcloudrun-golang/model"
)

// IndexHandler 入口函数
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.CallBackMsg{}
	rsp := &model.MsgRsp{Code: constant.Success, ErrorMsg: "ok", MsgType: constant.Text}
	defer func() {
		fmt.Printf("req :%+v rsp :%+v\n", req, rsp)
		msg, _ := json.Marshal(rsp)
		w.Header().Set("content-type", "application/json")
		w.Write(msg)
	}()

	// 解析请求
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		rsp.Code = constant.ErrParseParams
		rsp.ErrorMsg = fmt.Sprintf("json decode failed with %+v", err)
		return
	}
	rsp.ToUserName = req.FromUserName
	rsp.FromUserName = req.ToUserName
	rsp.CreateTime = req.CreateTime

	// 处理请求
	c := req.GetCmd()
	h := handler.GetHandler(c)
	if h == nil {
		rsp.Code = constant.ErrParseParams
		rsp.ErrorMsg = fmt.Sprintf("cmd[%v] not found", c)
		return
	}
	if err := h.Handle(req, rsp); err != nil {
		rsp.Code = constant.ErrHandle
		rsp.ErrorMsg = fmt.Sprintf("handler failed with %+v", err)
		return
	}
}
