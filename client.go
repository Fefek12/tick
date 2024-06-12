package main

import (
	"errors"
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_TYPE = "tcp"
)

type ClientConn interface {
	connect() (bool, error)
}

type Client struct {
	addr  string
	state [3][3]string `json:board`
	conn  net.Conn
}

func NewClient(port string) (*Client, error) {
	addr := fmt.Sprintf("localhost:%s", port)
	conn, err := net.Dial(SERVER_TYPE, addr)
	if err != nil {
		return nil, errors.New("error connecting to server")
	}
	msg := fmt.Sprintf("Connection was made too %s", conn.RemoteAddr())
	fmt.Println(msg)
	return &Client{}, nil

}
