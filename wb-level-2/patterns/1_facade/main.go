package main

func main() {

}

// Service реализует паттерн фасад
type Service struct {
	s Store
	c Cache
	t TokenService
}

//Store ...
type Store interface {
	SaveValue(val interface{}) error
	GetValue(id int64) (interface{}, error)
}

//Cache ...
type Cache interface {
	SaveValue(val interface{}) error
	GetValue(id int64) (interface{}, error)
}

//TokenService ...
type TokenService interface {
	CheckToken(token string) error
}

//SaveValue ...
func (s *Service) SaveValue(token string, v interface{}) error {
	if err := s.t.CheckToken(token); err != nil {
		return err
	}

	if err := s.s.SaveValue(v); err != nil {
		return err
	}

	s.c.SaveValue(v)

	return nil
}

//GetValue ...
func (s *Service) GetValue(token string, id int64) (interface{}, error) {
	if err := s.t.CheckToken(token); err != nil {
		return nil, err
	}

	if v, err := s.c.GetValue(id); err == nil {
		return v, err
	}

	v, err := s.s.GetValue(id)
	if err != nil {
		return nil, err
	}

	s.c.SaveValue(v)

	return v, nil
}
