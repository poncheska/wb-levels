package main

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net"
	"testing"
	"time"
)

const respVal = "qrqwrFDSGF42fGF2fsdfr2rdfs34DFSfs453gSg\n"

func TestTelnet(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	t.Run("ok test", func(t *testing.T) {
		go func() {
			ln, _ := net.Listen("tcp", ":8081")
			conn, _ := ln.Accept()
			bufio.NewReader(conn).ReadString('\n')
			conn.Write([]byte(respVal))
			conn.Close()
		}()
		time.Sleep(100 * time.Millisecond)

		host, port := "127.0.0.1", "8081"
		timeout := 5 * time.Second
		in := bytes.NewReader([]byte("GET /\n"))
		out := bytes.NewBuffer([]byte{})
		client := NewClient(net.JoinHostPort(host, port), timeout, in, out)
		err := client.Connect()
		if err != nil {
			panic(err)
		}
		defer client.Close()
		err = client.Send()
		if err != nil {
			panic(err)
		}
		for {
			err := (*client).Receive()
			if err != nil {
				break
			}
		}
		assert.Equal(t, respVal, out.String())
	})

	t.Run("invalid host/port", func(t *testing.T) {
		host, port := "127.0.0.1", "99999999999"
		timeout := 5 * time.Second
		in := bytes.NewReader([]byte("GET /\n"))
		out := bytes.NewBuffer([]byte{})
		client := NewClient(net.JoinHostPort(host, port), timeout, in, out)
		assert.NotNil(t, client.Connect())
	})

}
