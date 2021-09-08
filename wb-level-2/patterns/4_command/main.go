package main

import (
	"fmt"
)

func main() {
	mail := NewMailService()
	user := NewUserService(mail)

	cc := make(chan Command, 5)
	waitCh := make(chan struct{})

	// Executor
	go func() {
		for v := range cc {
			v.Execute()
		}
		waitCh <- struct{}{}
	}()

	cc <- user.SendMessage("fafaf")

	commands := []Command{
		user.SendMessages([]string{"a", "b", "c", "d", "e"}),
		user.SendMessage("fafaf"),
		user.SendMessage("fagaf"),
		user.SendMessage("fahaf"),
		user.SendMessages([]string{"g", "d", "c", "y", "q"}),
		user.SendMessages([]string{"h", "q", "u"}),
		user.SendMessage("fahafa"),
	}

	for _, v := range commands {
		cc <- v
	}

	close(cc)

	<- waitCh
}

type Mail interface {
	Send(m string)
	SendAll(m []string)
}

type MailService struct{}

func NewMailService() *MailService {
	return &MailService{}
}

func (s *MailService) Send(m string) {
	fmt.Println("Exec Send")
	fmt.Println(m)
	// TODO
}

func (s *MailService) SendAll(m []string) {
	fmt.Println("Exec SendAll")
	for _, v := range m {
		fmt.Println(v)
	}
	// TODO
}

type UserService struct {
	mail Mail
}

func NewUserService(mail Mail) *UserService {
	return &UserService{
		mail: mail,
	}
}

type Command interface {
	Execute()
}

type SendMessageCommand struct {
	M string
	s Mail
}

func (s *UserService) SendMessage(m string) Command {
	return &SendMessageCommand{
		M: m,
		s: s.mail,
	}
}

func (c *SendMessageCommand) Execute() {
	c.s.Send(c.M)
}

type SendMessagesCommand struct {
	M []string
	s Mail
}

func (s *UserService) SendMessages(m []string) Command {
	return &SendMessagesCommand{
		M: m,
		s: s.mail,
	}
}

func (c *SendMessagesCommand) Execute() {
	c.s.SendAll(c.M)
}
