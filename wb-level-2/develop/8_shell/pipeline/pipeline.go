package pipeline

import (
	"bytes"
	"git.wildberries.ru/Kovgar.Aleksey/wb-levels/wb-level-2/tasks/8_shell/commands"
	"io"
	"strings"
)

type Pipeline struct {
	Commands []commands.Command
}

func ParsePipeline(
	line string,
	pstdin io.ReadWriter,
	pstdout io.ReadWriter,
	pstderr io.ReadWriter,
) (*Pipeline, error) {
	lines := strings.Split(line, "|")
	cmds := make([]commands.Command, 0, len(lines))
	var stdin io.ReadWriter
	var stdout io.ReadWriter
	for i, v := range lines {
		if i == 0 {
			stdin = pstdin
		} else {
			stdin = stdout
		}
		if i == len(lines)-1 {
			stdout = pstdout
		} else {
			stdout = bytes.NewBuffer([]byte{})
		}
		cmd, err := commands.ParseCommand(v, stdin, stdout, pstderr)
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, cmd)
	}
	return &Pipeline{cmds}, nil
}

func (p *Pipeline) Execute() bool {
	ok := true
	for _, v := range p.Commands {
		if ok {
			ok = v.Execute()
		} else {
			return false
		}
	}
	return true
}