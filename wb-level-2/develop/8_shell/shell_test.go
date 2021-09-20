package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestShell(t *testing.T) {
	t.Run("test cd and pwd", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		dd, err := parseCommand("", "pwd")
		assert.Nil(t, err)
		assert.Equal(t, dir, dd)
		_, err = parseCommand("", "cd test_dir")
		assert.Nil(t, err)
		dd, err = parseCommand("", "pwd")
		assert.Nil(t, err)
		assert.Equal(t, dir+"/test_dir", dd)
	})
	t.Run("test echo", func(t *testing.T) {
		r, err := parseCommand("", "echo 123")
		assert.Nil(t, err)
		assert.Equal(t, "123", r)
	})
	t.Run("test exec", func(t *testing.T) {
		r, err := parseCommand("", "exec echo 123")
		assert.Nil(t, err)
		assert.Equal(t, "123\n", r)
	})
	// не убивает процессы созданные из самой программы
	//t.Run("test kill", func(t *testing.T) {
	//	cmd := exec.Command("./sig_int")
	//	cmd.Start()
	//	pid := cmd.Process.Pid
	//	ps, err := parseCommand("", "ps")
	//	assert.Nil(t, err)
	//	assert.Contains(t, ps, strconv.Itoa(pid))
	//	_, err = parseCommand("", "kill "+strconv.Itoa(pid))
	//	assert.Nil(t, err)
	//	ps, err = parseCommand("", "ps")
	//	assert.Nil(t, err)
	//	assert.NotContains(t, ps, strconv.Itoa(pid))
	//})
	t.Run("test ps and fork()", func(t *testing.T) {
		pid, err := parseCommand("", "fork()")
		assert.Nil(t, err)
		ps, err := parseCommand("", "ps")
		assert.Nil(t, err)
		assert.Contains(t, ps, pid)
	})
	t.Run("test pipeline", func(t *testing.T) {
		dir1, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		ff := strings.Split(dir1, "/")
		dir1 = strings.Join(ff[:len(ff)-1], "/")
		buf := bytes.NewBuffer([]byte{})
		parsePipeline(buf, "echo .. | cd")
		assert.Nil(t, err)
		dir2, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		assert.Equal(t, dir2, dir1)

	})
}
