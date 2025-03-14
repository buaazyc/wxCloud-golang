package handler

import (
	"sync"
	"wxcloudrun-golang/model"
)

// Handler 回调处理接口
type Handler interface {
	Handle(msg *model.CallBackMsg) (*model.MsgRsp, error) // 处理回调
}

var (
	handlers = make(map[string]Handler)
	lock     sync.Mutex
)

func RegisterHandler(name string, handler Handler) {
	lock.Lock()
	defer lock.Unlock()
	handlers[name] = handler
}

func GetHandler(name string) Handler {
	return handlers[name]
}
