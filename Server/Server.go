package Server

import (
	"fmt"
	"net"
)

type Server struct {
	port string
}

type Delta struct {
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func handleConnection(connection net.Conn) {
	fmt.Println("Accepted Connection from ", connection.RemoteAddr())

	defer connection.Close()
}

func (s *Server) Start() {
	connCount := 0
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
			fmt.Println("Error with Connection")
			continue
		}
		if connCount < 6 {
			fmt.Println(connCount)
			go handleConnection(conn)
			connCount++
		} else {
			fmt.Println("Connection count is Over 2")
		}

	}

}
