package model

import (
	"strings"
	"wxcloudrun-golang/constant"
)

// CallBackMsg 请求
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
type CallBackMsg struct {
	ToUserName   string `json:"ToUserName"`   // 小程序/公众号的原始ID
	FromUserName string `json:"FromUserName"` // 该小程序/公众号的用户身份openid
	CreateTime   int    `json:"CreateTime"`   // 消息时间(秒级时间戳)
	MsgType      string `json:"MsgType"`      // 消息类型
	Content      string `json:"Content"`      // 消息内容
	MsgId        int    `json:"MsgId"`        // 唯一消息ID

	tokens []string // 分隔消息内容
}

// GetCmd 获取命令
func (msg *CallBackMsg) GetCmd() string {
	if msg.MsgType != constant.Text || msg.Content == "" {
		return ""
	}
	if len(msg.tokens) == 0 {
		msg.tokens = strings.Split(msg.Content, " ")
	}
	return msg.tokens[0]
}

// MsgRsp 被动回复内容时的结构体
type MsgRsp struct {
	Code     int         `json:"code"`               // 返回码
	ErrorMsg string      `json:"errorMsg,omitempty"` // 错误信息
	Data     interface{} `json:"data"`               // 返回数据

	ToUserName   string `json:"ToUserName"`   // 用户OPENID
	FromUserName string `json:"FromUserName"` // 公众号/小程序原始ID
	CreateTime   int    `json:"CreateTime"`   // 消息时间(秒级时间戳)
	MsgType      string `json:"MsgType"`      // 消息类型
	Content      string `json:"Content"`      // 消息内容
}
