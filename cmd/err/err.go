package err

import (
	"fmt"
	"wxcloudrun-golang/handler"
	"wxcloudrun-golang/model"
)

func Init() {
	handler.RegisterHandler("err", &Handler{})
}

type Handler struct {
}

func (h *Handler) Handle(msg *model.CallBackMsg, rsp *model.MsgRsp) error {
	return fmt.Errorf("err")
}
