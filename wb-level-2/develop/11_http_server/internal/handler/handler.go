package handler

import (
	"encoding/json"
	"errors"
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/11_http_server/internal/domain"
	"log"
	"net/http"
)

func MarshalResult(ifc interface{}) []byte {
	res := struct {
		Res interface{} `json:"result"`
	}{ifc}
	bs, _ := json.Marshal(res)
	return bs
}

func MarshalError(err error) []byte {
	res := struct {
		Err string `json:"error"`
	}{err.Error()}
	bs, _ := json.Marshal(res)
	return bs
}

type Service interface {
	CreateEvent(e domain.Event) error
	UpdateEvent(e domain.Event) error
	DeleteEvent(eventId string) error
	EventsForDay(userId string) ([]domain.Event, error)
	EventsForWeek(userId string) ([]domain.Event, error)
	EventsForMonth(userId string) ([]domain.Event, error)
}

type Handler struct {
	svc Service
}

func New(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/create_event" && r.Method == http.MethodPost:
		h.createEvent(w, r)
	case r.URL.Path == "/update_event" && r.Method == http.MethodPost:
		h.updateEvent(w, r)
	case r.URL.Path == "/delete_event" && r.Method == http.MethodPost:
		h.deleteEvent(w, r)
	case r.URL.Path == "/events_for_day" && r.Method == http.MethodGet:
		h.eventsForDay(w, r)
	case r.URL.Path == "/events_for_week" && r.Method == http.MethodGet:
		h.eventsForWeek(w, r)
	case r.URL.Path == "/events_for_month" && r.Method == http.MethodGet:
		h.eventsForMonth(w, r)
	default:
		http.Error(w, "unsupported path", http.StatusNotFound)
	}
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	event := domain.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(400)
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	err = h.svc.CreateEvent(event)
	if err != nil {
		w.WriteHeader(503)
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	w.Write(MarshalResult("done"))
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	event := domain.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(400)
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	err = h.svc.UpdateEvent(event)
	if err != nil {
		w.WriteHeader(503)
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	w.Write(MarshalResult("done"))
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventId, ok := r.URL.Query()["event_id"]
	if !ok || len(eventId) < 1 {
		w.WriteHeader(400)
		err := errors.New("request must have parameter event_id")
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	err := h.svc.DeleteEvent(eventId[0])
	if err != nil {
		w.WriteHeader(503)
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	w.Write(MarshalResult("done"))
}

func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.URL.Query()["user_id"]
	if !ok || len(userId) < 1 {
		w.WriteHeader(400)
		err := errors.New("request must have parameter event_id")
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	res, err := h.svc.EventsForDay(userId[0])
	if err != nil {
		w.WriteHeader(503)
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	w.Write(MarshalResult(res))
}

func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.URL.Query()["user_id"]
	if !ok || len(userId) < 1 {
		w.WriteHeader(400)
		err := errors.New("request must have parameter user_id")
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	res, err := h.svc.EventsForWeek(userId[0])
	if err != nil {
		w.WriteHeader(503)
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	w.Write(MarshalResult(res))
}

func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.URL.Query()["user_id"]
	if !ok || len(userId) < 1 {
		w.WriteHeader(400)
		w.Write(MarshalError(errors.New("request must have parameter user_id")))
		return
	}
	res, err := h.svc.EventsForMonth(userId[0])
	if err != nil {
		w.WriteHeader(503)
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	w.Write(MarshalResult(res))
}
