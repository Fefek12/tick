package Server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Server struct {
	port      string
	listner   net.Listener
	board     *[3][3]string
	connCount int8
}

func NewServer(port string) *Server {
	initBoardState := [3][3]string{
		{"X", "X", "X"},
		{"y", "y", "y"},
		{"Z", "Z", "Z"},
	}
	p := ":" + port
	listener, err := net.Listen("tcp", p)
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		port:      p,
		listner:   listener,
		board:     &initBoardState,
		connCount: 2,
	}
}

func handleConnection(connection net.Conn, s *Server) {
	go sendBoardToClient(connection, s)
}

// func readClientsDelta() {

// 	// for {
// 	// 	byt := make([]byte, 1024)
// 	// 	size, err := connection.Read(byt)
// 	// 	if err == nil {
// 	// 		data := (strings.Split(string(byt[:size]), " "))
// 	// 		fmt.Println(data)
// 	// 		// 	delta := data[0]
// 	// 		// 	x, _ := strconv.ParseInt(data[1], 10, 64)
// 	// 		// 	y, _ := strconv.ParseInt(data[2], 10, 64)

// 	// 	}
// 	// }
// }

func sendBoardToClient(conn net.Conn, s *Server) {
	for {
		boardByte, err := json.Marshal(s.board)
		if err != nil {
			fmt.Println("Error Mashaling Board", err)
		}
		_, err = conn.Write(boardByte)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func (s *Server) Start() {
	fmt.Printf("Server is Listening at Port %s", s.port)
	for {
		conn, err := s.listner.Accept()
		if err != nil {
			fmt.Println("Error with Connection", conn)
			continue
		}
		//Refactor this later and Pass the connCount to the handle function so it can do the check on its own thread
		if s.connCount < 100 {
			fmt.Println("")
			fmt.Print("Connected", conn.LocalAddr().String()+"\n")
			go handleConnection(conn, s)
			s.connCount++
		} else {
			fmt.Println("Unable to Accept Connection as Connection count is Over 2")
			conn.Write([]byte("404"))
		}
	}
}
