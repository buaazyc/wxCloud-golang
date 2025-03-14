package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/constant"
)

const wxMsgUrl = "http://api.weixin.qq.com/cgi-bin/message/custom/send"

type WxMsgReq struct {
	ToUser  string    `json:"touser"`
	MsgType string    `json:"msgtype"`
	Text    WxMsgText `json:"text"`
}

type WxMsgText struct {
	Content string `json:"content"`
}

type WxMsgRsp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// SendWxMsg 发送微信消息，需要微信号认证后才可以发送
// 参考：https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Explanation_of_interface_privileges.html
func SendWxMsg(toUser, content string) error {
	jsonReq, err := json.Marshal(&WxMsgReq{
		ToUser:  toUser,
		MsgType: constant.Text,
		Text: WxMsgText{
			Content: content,
		},
	})
	if err != nil {
		return fmt.Errorf("json marshal failed with %+v", err)
	}
	rsp, err := http.Post(wxMsgUrl, "application/json", bytes.NewReader(jsonReq))
	if err != nil {
		return fmt.Errorf("http post failed with %+v", err)
	}
	defer rsp.Body.Close()
	wxMsgRsp := &WxMsgRsp{}
	if err := json.NewDecoder(rsp.Body).Decode(wxMsgRsp); err != nil {
		return fmt.Errorf("json decode failed with %+v", err)
	}
	if wxMsgRsp.ErrCode != 0 {
		return fmt.Errorf("wx msg failed with %+v", wxMsgRsp)
	}
	return nil
}
