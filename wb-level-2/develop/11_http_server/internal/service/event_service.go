package service

import (
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/11_http_server/internal/domain"
	"time"
)

//EventService ...
type EventService struct {
	s Store
}

//NewEventService ...
func NewEventService() *EventService {
	return &EventService{s: NewEventStorage()}
}

//Store ...
type Store interface {
	CreateEvent(e domain.Event) error
	UpdateEvent(e domain.Event) error
	DeleteEvent(eventID string) error
	GetEventsForPeriod(userID string, p1 time.Time, p2 time.Time) ([]domain.Event, error)
}

//CreateEvent ...
func (s *EventService) CreateEvent(e domain.Event) error {
	return s.s.CreateEvent(e)
}

//UpdateEvent ...
func (s *EventService) UpdateEvent(e domain.Event) error {
	return s.s.UpdateEvent(e)
}

//DeleteEvent ...
func (s *EventService) DeleteEvent(eventID string) error {
	return s.s.DeleteEvent(eventID)
}

//EventsForDay ...
func (s *EventService) EventsForDay(userID string) ([]domain.Event, error) {
	es, err := s.s.GetEventsForPeriod(userID, time.Now(), time.Now().Add(24*time.Hour))
	return es, err
}

//EventsForWeek ...
func (s *EventService) EventsForWeek(userID string) ([]domain.Event, error) {
	es, err := s.s.GetEventsForPeriod(userID, time.Now(), time.Now().Add(7*24*time.Hour))
	return es, err
}

//EventsForMonth ...
func (s *EventService) EventsForMonth(userID string) ([]domain.Event, error) {
	es, err := s.s.GetEventsForPeriod(userID, time.Now(), time.Now().Add(30*24*time.Hour))
	return es, err
}
