package main

import (
	"encoding/json"
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
	state [3][3]string
	// startDelta string
}

func NewClient(port string) (*Client, error) {
	addr := fmt.Sprintf("localhost:%s", port)
	conn, err := net.Dial(SERVER_TYPE, addr)
	if err != nil {
		panic(err)
	}
	msg := fmt.Sprintf("Connection was made too %s", conn.RemoteAddr())
	fmt.Println(msg)

	buff := make([]byte, 1024)
	size, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading Buffer", err)
	}
	if string(buff[:size]) == "404" {
		panic("Connection to Server is full")
	}
	var boardState [3][3]string

	err = json.Unmarshal(buff[:size], &boardState)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	return &Client{
		state: boardState,
	}, nil
}
// func (c *Client) play(delta string) error {}
