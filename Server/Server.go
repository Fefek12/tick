package Server

import (
	"encoding/json"
	"fmt"
	"net"
)

type Server struct {
	port string
}

type Delta struct {
	state [3][3]string `json:"boardState"`
	done  bool         `json:"done"`
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func handleConnection(connection net.Conn) {
	initBoardState := [3][3]string{
		{"S", "S", "S"},
		{"y", "y", "y"},
		{"z", "z", "z"},
	}
	boardByte, err := json.Marshal(initBoardState)
	if err != nil {
		fmt.Println("Error Mashaling Board", err)
		return
	}
	_, err = connection.Write(boardByte)
	if err != nil {
		fmt.Println("Error Sending Message", err)
		return
	}
}

func (s *Server) Start() {
	connCount := 0
	m := make([]string, 2)
	port := ":" + s.port
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer listener.Close()
	fmt.Printf("Server is Listening at Port %s", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error with Connection", conn)
			continue
		}
		//Refactor this later and Pass the connCount to the handle function so it can do the check on its own thread
		if connCount < 2 {
			go handleConnection(conn)
			m = append(m, conn.LocalAddr().String())
			connCount++
		} else {
			fmt.Println("Unable to Accept Connection as Connection count is Over 2")
			conn.Write([]byte("404"))
		}
	}
}
