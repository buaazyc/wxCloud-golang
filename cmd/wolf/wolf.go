package wolf

import (
	"wxcloudrun-golang/constant"
	"wxcloudrun-golang/handler"
	"wxcloudrun-golang/model"
)

func Init() {
	handler.RegisterHandler("狼人杀", &Handler{})
}

type Handler struct {
}

func (h *Handler) Handle(msg *model.CallBackMsg) (*model.MsgRsp, error) {
	sendMsg := "hello world"
	return &model.MsgRsp{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   msg.CreateTime,
		MsgType:      constant.Text,
		Content:      sendMsg,
	}, nil
}
