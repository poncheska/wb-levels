package commands

import (
	"io"
	"io/ioutil"
	"os"
)

type Command interface {
	Execute() bool
}

type CommandInfo struct {
	Name   string
	StdIn  io.ReadWriter
	StdOut io.ReadWriter
	StdErr io.ReadWriter
	Props  []Property
}

type Property struct {
	Name  string
	Value string
}

func ParseCommand(
	line string,
	stdin io.ReadWriter,
	stdout io.ReadWriter,
	stderr io.ReadWriter,
) (Command, error) {
	return nil, nil
}

type CDCommand struct {
	CommandInfo
}

func (c *CDCommand) Execute() bool {
	dir, err := ioutil.ReadAll(c.StdIn)
	if err != nil {
		c.StdErr.Write([]byte(err.Error()))
		return false
	}
	err = os.Chdir(string(dir))
	if err != nil {
		c.StdErr.Write([]byte(err.Error()))
		return false
	}
	return true
}
