package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Req 请求
// 消息体样例：
//
//	{
//		"ToUserName": "gh_919b00572d95", // 小程序/公众号的原始ID，资源复用配置多个时可以区别消息是给谁的
//		"FromUserName": "oVneZ57wJnV-ObtCiGv26PRrOz2g", // 该小程序/公众号的用户身份openid
//		"CreateTime": 1651049934, // 消息时间
//		"MsgType": "text", // 消息类型
//		"Content": "回复文本", // 消息内容
//		"MsgId": 23637352235060880 // 唯一消息ID，可能发送多个重复消息，需要注意用此ID去重
//	  }
type Req struct {
	ToUserName   string `json:"ToUserName"`   // 小程序/公众号的原始ID
	FromUserName string `json:"FromUserName"` // 该小程序/公众号的用户身份openid
	CreateTime   int    `json:"CreateTime"`   // 消息时间(秒级时间戳)
	MsgType      string `json:"MsgType"`      // 消息类型
	Content      string `json:"Content"`      // 消息内容
	MsgId        int    `json:"MsgId"`        // 唯一消息ID
}

// Rsp 返回结构
type Rsp struct {
	Code     int         `json:"code"`               // 返回码
	ErrorMsg string      `json:"errorMsg,omitempty"` // 错误信息
	Data     interface{} `json:"data"`               // 返回数据
}

// IndexHandler 入口函数
func IndexHandler(rsp http.ResponseWriter, r *http.Request) {
	req := &Req{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		fmt.Printf("json decode failed with %+v", err)
		return
	}
	fmt.Printf("req :%+v\n", req)
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
