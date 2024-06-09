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
	fmt.Println("Accepted COnnection from ", connection.RemoteAddr())

	defer connection.Close()
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer listener.Close()
	fmt.Printf("Server is Listening at Port %s", s.port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error with Connection")
			continue
		}
		go handleConnection(conn)
	}

}
