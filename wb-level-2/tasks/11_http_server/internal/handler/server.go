package handler

import (
	"git.wildberries.ru/Kovgar.Aleksey/wb-levels/wb-level-2/tasks/11_http_server/internal/service"
	"net/http"
)

type Handler struct {
	svc *service.EventService
	mux *http.ServeMux
}

func NewHandler() *Handler {
	return &Handler{}
}

func (s *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch {
	case r.URL.Path == "/create_event" && r.Method == http.MethodPost:
		s.svc.CreateEvent(w, r)
	case r.URL.Path == "/update_event" && r.Method == http.MethodPost:
		s.svc.UpdateEvent(w, r)
	default:
		http.Error(w, "Unsupported path", http.StatusNotFound)
	}
}
