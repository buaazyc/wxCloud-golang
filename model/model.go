package model

import "strings"

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
}

// GetCmd 获取命令
func (msg *CallBackMsg) GetCmd() string {
	if msg.MsgType != "text" || msg.Content == "" {
		return ""
	}
	tokens := strings.Split(msg.Content, " ")
	return tokens[0]
}
