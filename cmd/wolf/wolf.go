package wolf

import (
	"wxcloudrun-golang/handler"
	"wxcloudrun-golang/model"
)

func Init() {
	handler.RegisterHandler("狼人杀", &Handler{})
}

type Handler struct {
}

func (h *Handler) Handle(msg *model.CallBackMsg, rsp *model.MsgRsp) error {
	sendMsg := "欢迎来玩狼人杀！"
	rsp.Content = sendMsg
	return nil
}
