package main

import (
	"errors"
	"fmt"
)

func main() {
	cp := NewConnPool(3)
	fmt.Println("put error: ", cp.PutConn("aaa"))
	fmt.Println("put error: ", cp.PutConn("bbb"))
	fmt.Println("put error: ", cp.PutConn("ccc"))
	fmt.Println("put error: ", cp.PutConn("ddd"))
	conn, err := cp.GetConn()
	fmt.Printf("get conn: %v error: %v\n", conn, err)
	conn, err = cp.GetConn()
	fmt.Printf("get conn: %v error: %v\n", conn, err)
	conn, err = cp.GetConn()
	fmt.Printf("get conn: %v error: %v\n", conn, err)
	conn, err = cp.GetConn()
	fmt.Printf("get conn: %v error: %v\n", conn, err)
	fmt.Println("put error: ", cp.PutConn("ddd"))
	conn, err = cp.GetConn()
	fmt.Printf("get conn: %v error: %v\n", conn, err)
}

//ConnPool ...
type ConnPool struct {
	Conns []string
	Limit int
	state ConnPoolState
}

//NewConnPool ...
func NewConnPool(limit int) *ConnPool {
	if limit < 1 {
		limit = 1
	}
	cp := &ConnPool{Conns: []string{}, Limit: limit}
	cp.ChangeState(NewPutOnlyState(cp))
	return cp
}

//ChangeState ...
func (cp *ConnPool) ChangeState(s ConnPoolState) {
	cp.state = s
}

//GetConn ...
func (cp *ConnPool) GetConn() (string, error) {
	conn, err := cp.state.GetConn()
	return conn, err
}

//PutConn ...
func (cp *ConnPool) PutConn(conn string) error {
	return cp.state.PutConn(conn)
}

//ConnPoolState ...
type ConnPoolState interface {
	GetConn() (string, error)
	PutConn(conn string) error
}

//NormalState ...
type NormalState struct {
	cp *ConnPool
}

//NewNormalState ...
func NewNormalState(cp *ConnPool) *NormalState {
	return &NormalState{cp: cp}
}

//GetConn ...
func (s *NormalState) GetConn() (string, error) {
	res := s.cp.Conns[0]
	s.cp.Conns = s.cp.Conns[1:]

	if len(s.cp.Conns) == 0 {
		s.cp.ChangeState(NewPutOnlyState(s.cp))
	}

	return res, nil
}

//PutConn ...
func (s *NormalState) PutConn(conn string) error {
	s.cp.Conns = append(s.cp.Conns, conn)

	if len(s.cp.Conns) == s.cp.Limit {
		s.cp.ChangeState(NewGetOnlyState(s.cp))
	}

	return nil
}

//GetOnlyState ...
type GetOnlyState struct {
	cp *ConnPool
}

//NewGetOnlyState ...
func NewGetOnlyState(cp *ConnPool) *GetOnlyState {
	return &GetOnlyState{cp: cp}
}

//GetConn ...
func (s *GetOnlyState) GetConn() (string, error) {
	res := s.cp.Conns[0]
	s.cp.Conns = s.cp.Conns[1:]
	if len(s.cp.Conns) == 0 {
		s.cp.ChangeState(NewPutOnlyState(s.cp))
	} else {
		s.cp.ChangeState(NewNormalState(s.cp))
	}
	return res, nil
}

//PutConn ...
func (s *GetOnlyState) PutConn(conn string) error {
	return errors.New("get only state")
}

//PutOnlyState ...
type PutOnlyState struct {
	cp *ConnPool
}

//NewPutOnlyState ...
func NewPutOnlyState(cp *ConnPool) *PutOnlyState {
	return &PutOnlyState{cp: cp}
}

//GetConn ...
func (s *PutOnlyState) GetConn() (string, error) {
	return "", errors.New("put only state")
}

//PutConn ...
func (s *PutOnlyState) PutConn(conn string) error {
	s.cp.Conns = append(s.cp.Conns, conn)

	if len(s.cp.Conns) == s.cp.Limit {
		s.cp.ChangeState(NewGetOnlyState(s.cp))
	} else {
		s.cp.ChangeState(NewNormalState(s.cp))
	}

	return nil
}
