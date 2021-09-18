package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout")
	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Println("telnet: must have 2 arguments: host and port")
	}
	host := flag.Args()[0]
	port := flag.Args()[1]
	if err := Telnet(*timeout, host, port); err != nil {
		fmt.Println("telnet: " + err.Error())
	}
}

func Telnet(timeout time.Duration, host, port string) error {
	addr := net.JoinHostPort(host, port)
	client := NewClient(addr, timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		return err
	}
	defer client.Close()

	notifyCh := make(chan os.Signal, 1)
	errorCh := make(chan error, 1)
	signal.Notify(notifyCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			err := (*client).Send()
			if err != nil {
				errorCh <- err
				return
			}
		}
	}()

	go func() {
		for {
			err := (*client).Receive()
			if err != nil {
				errorCh <- err
				return
			}
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	var err error
	go func() {
		defer wg.Done()
		for {
			select {
			case <-notifyCh:
				return
			case err = <-errorCh:
				if err != nil {
					return
				}
			}
		}
	}()

	wg.Wait()
	return err
}

type Client struct {
	addr       string
	timeout    time.Duration
	conn       net.Conn
	connReader *bufio.Reader
	in         *bufio.Reader
	out        io.Writer
}

func NewClient(addr string, timeout time.Duration, in io.Reader, out io.Writer) *Client {
	return &Client{
		addr:    addr,
		timeout: timeout,
		in:      bufio.NewReader(in),
		out:     out,
	}
}

func (c *Client) Connect() error {
	conn, err := net.DialTimeout("tcp", c.addr, c.timeout)
	if err != nil {
		return err
	}
	c.conn = conn
	c.connReader = bufio.NewReader(c.conn)
	log.Println("Connected to " + c.addr)
	return nil
}

func (c *Client) Receive() error {
	line, err := c.connReader.ReadBytes('\n')
	if err != nil {
		return err
	}
	if _, err := c.out.Write(line); err != nil {
		return err
	}
	return nil
}

func (c *Client) Send() error {
	line, err := c.in.ReadBytes('\n')
	if err != nil {
		return err
	}
	_, err = c.conn.Write(line)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() error {
	err := c.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
