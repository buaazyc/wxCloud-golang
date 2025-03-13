package wolf

import (
	"log"
	"wxcloudrun-golang/handler"
	"wxcloudrun-golang/model"
)

func Init() {
	handler.RegisterHandler("狼人杀", &Handler{})
}

type Handler struct {
}

func (h *Handler) Handle(msg *model.CallBackMsg) error {
	log.Printf("wolf Handle msg: %+v", msg)
	return nil
}
