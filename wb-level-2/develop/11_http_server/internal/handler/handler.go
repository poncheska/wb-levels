package handler

import (
	"encoding/json"
	"errors"
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/11_http_server/internal/domain"
	"log"
	"net/http"
)

//MarshalResult ...
func MarshalResult(ifc interface{}) []byte {
	res := struct {
		Res interface{} `json:"result"`
	}{ifc}
	bs, _ := json.Marshal(res)
	return bs
}

//MarshalError ...
func MarshalError(err error) []byte {
	res := struct {
		Err string `json:"error"`
	}{err.Error()}
	bs, _ := json.Marshal(res)
	return bs
}

//Service ...
type Service interface {
	CreateEvent(e domain.Event) error
	UpdateEvent(e domain.Event) error
	DeleteEvent(eventID string) error
	EventsForDay(userID string) ([]domain.Event, error)
	EventsForWeek(userID string) ([]domain.Event, error)
	EventsForMonth(userID string) ([]domain.Event, error)
}

//Handler ...
type Handler struct {
	svc Service
}

//New ...
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
	eventID, ok := r.URL.Query()["event_id"]
	if !ok || len(eventID) < 1 {
		w.WriteHeader(400)
		err := errors.New("request must have parameter event_id")
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	err := h.svc.DeleteEvent(eventID[0])
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
	userID, ok := r.URL.Query()["user_id"]
	if !ok || len(userID) < 1 {
		w.WriteHeader(400)
		err := errors.New("request must have parameter event_id")
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	res, err := h.svc.EventsForDay(userID[0])
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
	userID, ok := r.URL.Query()["user_id"]
	if !ok || len(userID) < 1 {
		w.WriteHeader(400)
		err := errors.New("request must have parameter user_id")
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	res, err := h.svc.EventsForWeek(userID[0])
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
	userID, ok := r.URL.Query()["user_id"]
	if !ok || len(userID) < 1 {
		w.WriteHeader(400)
		w.Write(MarshalError(errors.New("request must have parameter user_id")))
		return
	}
	res, err := h.svc.EventsForMonth(userID[0])
	if err != nil {
		w.WriteHeader(503)
		w.Write(MarshalError(err))
		log.Println(err)
		return
	}
	w.WriteHeader(200)
	w.Write(MarshalResult(res))
}
