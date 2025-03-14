package help

import (
	"wxcloudrun-golang/handler"
	"wxcloudrun-golang/model"
)

func Init() {
	handler.RegisterHandler("help", &Handler{})
}

type Handler struct {
}

func (h *Handler) Handle(msg *model.CallBackMsg, rsp *model.MsgRsp) error {
	sendMsg := "hello world!"
	rsp.Content = sendMsg
	return nil
}
