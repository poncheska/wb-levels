package main

import (
	"errors"
	"fmt"
)

func main() {
	cp := NewConnPull(3)
	fmt.Println("put error: ",cp.PutConn("aaa"))
	fmt.Println("put error: ",cp.PutConn("bbb"))
	fmt.Println("put error: ",cp.PutConn("ccc"))
	fmt.Println("put error: ",cp.PutConn("ddd"))
	conn, err := cp.GetConn()
	fmt.Printf("get conn: %v error: %v\n",conn,err)
	conn, err = cp.GetConn()
	fmt.Printf("get conn: %v error: %v\n",conn,err)
	conn, err = cp.GetConn()
	fmt.Printf("get conn: %v error: %v\n",conn,err)
	conn, err = cp.GetConn()
	fmt.Printf("get conn: %v error: %v\n",conn,err)
	fmt.Println("put error: ",cp.PutConn("ddd"))
	conn, err = cp.GetConn()
	fmt.Printf("get conn: %v error: %v\n",conn,err)
}

type ConnPull struct {
	Conns []string
	Limit int
	state ConnPullState
}

func NewConnPull(limit int) *ConnPull {
	if limit < 1 {
		limit = 1
	}
	cp := &ConnPull{Conns: []string{}, Limit: limit}
	cp.ChangeState(NewPutOnlyState(cp))
	return cp
}

func (cp *ConnPull) ChangeState(s ConnPullState) {
	cp.state = s
}

func (cp *ConnPull) GetConn() (string, error) {
	conn, err := cp.state.GetConn()
	return conn, err
}

func (cp *ConnPull) PutConn(conn string) error {
	return cp.state.PutConn(conn)
}

type ConnPullState interface {
	GetConn() (string, error)
	PutConn(conn string) error
}

type NormalState struct {
	cp *ConnPull
}

func NewNormalState(cp *ConnPull) *NormalState {
	return &NormalState{cp: cp}
}

func (s *NormalState) GetConn() (string, error) {
	res := s.cp.Conns[0]
	s.cp.Conns = s.cp.Conns[1:]

	if len(s.cp.Conns) == 0 {
		s.cp.ChangeState(NewPutOnlyState(s.cp))
	}

	return res, nil
}

func (s *NormalState) PutConn(conn string) error {
	s.cp.Conns = append(s.cp.Conns, conn)

	if len(s.cp.Conns) == s.cp.Limit {
		s.cp.ChangeState(NewGetOnlyState(s.cp))
	}

	return nil
}

type GetOnlyState struct {
	cp *ConnPull
}

func NewGetOnlyState(cp *ConnPull) *GetOnlyState {
	return &GetOnlyState{cp: cp}
}

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

func (s *GetOnlyState) PutConn(conn string) error {
	return errors.New("get only state")
}

type PutOnlyState struct {
	cp *ConnPull
}

func NewPutOnlyState(cp *ConnPull) *PutOnlyState {
	return &PutOnlyState{cp: cp}
}

func (s *PutOnlyState) GetConn() (string, error) {
	return "", errors.New("put only state")
}

func (s *PutOnlyState) PutConn(conn string) error {
	s.cp.Conns = append(s.cp.Conns, conn)

	if len(s.cp.Conns) == s.cp.Limit {
		s.cp.ChangeState(NewGetOnlyState(s.cp))
	} else {
		s.cp.ChangeState(NewNormalState(s.cp))
	}

	return nil
}
