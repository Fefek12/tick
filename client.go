package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

const (
	SERVER_HOST = "localhost"
	SERVER_TYPE = "tcp"
)

type client struct {
	state      [3][3]string
	connection net.Conn
	mutx       sync.Mutex
}

func NewClient(port string) (*client, error) {
	addr := fmt.Sprintf("localhost:%s", port)
	conn, err := net.Dial(SERVER_TYPE, addr)
	if err != nil {
		panic(err)
	}

	msg := fmt.Sprintf("Connection was made too %s", conn.RemoteAddr().String())
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
	return &client{
		state:      boardState,
		connection: conn,
		mutx:       sync.Mutex{},
	}, nil
}

func (c *client) Render() {
	c.fetchBoard()
	c.mutx.Lock()
	fmt.Println("Preparing Engine")
	eng := Engine{
		c.state,
	}
	c.mutx.Unlock()
	eng.Render()
}

func (c *client) fetchBoard() {
	fmt.Println("Fetching Board")
	reader := bufio.NewReader(c.connection)
	var board [3][3]string
	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Error Reading Bytes", err)
	}

	err = json.Unmarshal(data, &board)
	if err != nil {
		fmt.Println("Error Unmashaling", err)
	}
	c.mutx.Lock()
	fmt.Println("New Board Recieved", board)
	c.state = board
	fmt.Println("Board has been added to client", c.state)
	c.mutx.Unlock()
}

func (c *client) SendDelta(res string) {
	_, err := c.connection.Write([]byte(res + "\n"))
	if err != nil {
		fmt.Println("Error:", err)
	}
}
