package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mitchellh/go-ps"
	"io"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

var (
	green  = "\033[0;32m"
	purple = "\033[0;35m"
	nc     = "\033[0m"
)

func main() {
	startShell()
}

func startShell() {
	sc := bufio.NewScanner(os.Stdin)
	printHeader()
	for sc.Scan() {
		txt := sc.Text()
		if txt == "exit" {
			return
		}
		parsePipeline(os.Stdout, txt)
		//parseFork(os.Stdin, os.Stdout, txt)
		printHeader()
	}
}

//func parseFork(in io.Reader, out io.Writer, line string) {
//	ff := strings.Split(line, "&")
//}
//
func parsePipeline(out io.Writer, line string) {
	ss := strings.Split(line, "|")
	var res string
	var err error
	for _, v := range ss {
		res, err = parseCommand(res, v)
		if err != nil {
			fmt.Fprintln(out, err)
		}
	}
	if res != "" {
		fmt.Fprintln(out, res)
	}
}

func parseCommand(in string, line string) (string, error) {
	ss := strings.Fields(line)
	if len(ss) == 0 {
		return "", errors.New("empty command")
	}
	if in != "" {
		ss = append(ss, in)
	}
	switch ss[0] {
	case "cd":
		if len(ss) != 2 {
			return "", errors.New("cd: must be 1 parameter")
		}
		err := os.Chdir(ss[1])
		if err != nil {
			return "", err
		}
	case "pwd":
		if len(ss) != 1 {
			return "", errors.New("pwd: unused parameters")
		}
		res, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return res, nil
	case "echo":
		if len(ss) != 2 {
			return "", errors.New("echo: must be 1 parameter")
		}
		return ss[1], nil
	case "kill":
		if len(ss) < 2 {
			return "", errors.New("kill: not enough parameters")
		}
		for i := 1; i < len(ss); i++ {
			pid, err := strconv.Atoi(ss[i])
			if err != nil {
				return "", err
			}
			err = syscall.Kill(pid, syscall.SIGINT)
			if err != nil {
				return "", err
			}
		}
		return "", nil
	case "ps":
		processes, err := ps.Processes()
		if err != nil {
			return "", err
		}
		var builder strings.Builder
		builder.WriteString("\tPID\tCMD\n")
		for _, proc := range processes {
			builder.WriteString(
				fmt.Sprintf("\t%v\t%v\n", proc.Pid(), proc.Executable()),
			)
		}
		return builder.String(), nil
	case "fork()":
		id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		return strconv.Itoa(int(id)), nil
	case "exec":
		if len(ss) < 2 {
			return "", errors.New("exec: not enough parameters")
		}
		cmd := exec.Command(ss[1], ss[2:]...)
		stdout, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return string(stdout), err
	case "netcat":
		if len(ss) == 2 {
			res, err := netcat(ss[1], false)
			if err != nil {
				return "", err
			}
			return res, err
		} else if len(ss) == 3 {
			if ss[1] == "-u" {
				res, err := netcat(ss[2], true)
				if err != nil {
					return "", err
				}
				return res, err
			}
		}
	}
	return "", nil
}

func printHeader() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%vgo-shell%v<%v>%v: ", green, purple, wd, nc)
}

func netcat(addr string, isUdp bool) (string, error) {
	network := "tcp"
	if isUdp {
		network = "udp"
	}
	con, err := net.Dial(network, addr)
	if err != nil {
		return "", err
	}
	defer con.Close()
	err = stdInToConn(con)
	return "", err
}

func stdInToConn(conn net.Conn) error {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		if txt == "exit" {
			return nil
		}
		_, err := conn.Write([]byte(txt))
		if err != nil {
			return err
		}
	}
	return nil
}
