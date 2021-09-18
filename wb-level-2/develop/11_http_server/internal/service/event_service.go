package service

import (
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/11_http_server/internal/domain"
	"time"
)

type EventService struct {
	s Store
}

func NewEventService() *EventService {
	return &EventService{s: NewEventStorage()}
}

type Store interface {
	CreateEvent(e domain.Event) error
	UpdateEvent(e domain.Event) error
	DeleteEvent(eventId string) error
	GetEventsForPeriod(userId string, p1 time.Time, p2 time.Time) ([]domain.Event, error)
}

func (s *EventService) CreateEvent(e domain.Event) error {
	return s.s.CreateEvent(e)
}

func (s *EventService) UpdateEvent(e domain.Event) error {
	return s.s.UpdateEvent(e)
}

func (s *EventService) DeleteEvent(eventId string) error {
	return s.s.DeleteEvent(eventId)
}

func (s *EventService) EventsForDay(userId string) ([]domain.Event, error) {
	es, err := s.s.GetEventsForPeriod(userId, time.Now(), time.Now().Add(24*time.Hour))
	return es, err
}

func (s *EventService) EventsForWeek(userId string) ([]domain.Event, error) {
	es, err := s.s.GetEventsForPeriod(userId, time.Now(), time.Now().Add(7*24*time.Hour))
	return es, err
}

func (s *EventService) EventsForMonth(userId string) ([]domain.Event, error) {
	es, err := s.s.GetEventsForPeriod(userId, time.Now(), time.Now().Add(30*24*time.Hour))
	return es, err
}
