package service

import (
	"errors"
	"git.wildberries.ru/kovgar.aleksey/wb-levels/wb-level-2/develop/11_http_server/internal/domain"
	"sync"
	"time"
)

//EventStorage ...
type EventStorage struct {
	events map[string]domain.Event
	mx     *sync.RWMutex
}

//NewEventStorage ...
func NewEventStorage() *EventStorage {
	return &EventStorage{
		events: make(map[string]domain.Event),
		mx:     &sync.RWMutex{},
	}
}

//CreateEvent ...
func (s *EventStorage) CreateEvent(e domain.Event) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if _, ok := s.events[e.ID]; !ok {
		s.events[e.ID] = e
		return nil
	}
	return errors.New("event with this id already exist")
}

//UpdateEvent ...
func (s *EventStorage) UpdateEvent(e domain.Event) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if _, ok := s.events[e.ID]; ok {
		s.events[e.ID] = e
		return nil
	}
	return errors.New("event with this id not exist")
}

//DeleteEvent ...
func (s *EventStorage) DeleteEvent(eventID string) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if _, ok := s.events[eventID]; ok {
		delete(s.events, eventID)
		return nil
	}
	return errors.New("event with this id not exist")
}

//GetEventsForPeriod ...
func (s *EventStorage) GetEventsForPeriod(userID string, p1 time.Time, p2 time.Time) ([]domain.Event, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	res := make([]domain.Event, 0)
	for _, v := range s.events {
		if v.Date.Before(p2) && v.Date.After(p1) && v.CreatorID == userID {
			res = append(res, v)
		}
	}
	return res, nil
}
