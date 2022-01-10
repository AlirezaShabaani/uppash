package httpHandler

import (
	"espad/back/uppash/internal/core/port"
	"net/http"
)

type HttpHndlr struct {
	services port.TusSrv
}

func New(services port.TusSrv) HttpHndlr  {
	return HttpHndlr{services: services}
}

func (h *HttpHndlr)Upload(w http.ResponseWriter,r *http.Request)   {
}
