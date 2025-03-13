package wolf

import (
	"fmt"
	"wxcloudrun-golang/client"
	"wxcloudrun-golang/handler"
	"wxcloudrun-golang/model"
)

func Init() {
	handler.RegisterHandler("狼人杀", &Handler{})
}

type Handler struct {
}

func (h *Handler) Handle(msg *model.CallBackMsg) error {
	if err := client.SendWxMsg(msg.FromUserName, "hello world"); err != nil {
		return fmt.Errorf("send wx msg failed: %v", err)
	}
	return nil
}
