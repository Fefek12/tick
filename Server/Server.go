package Server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Server struct {
	port        string
	listner     net.Listener
	board       *[3][3]string
	connections []net.Conn
	wg          sync.WaitGroup
	mutx        sync.Mutex
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
		port:        p,
		listner:     listener,
		board:       &initBoardState,
		connections: make([]net.Conn, 0),
		wg:          sync.WaitGroup{},
	}
}

func handleConnection(connection net.Conn, s *Server) {
	defer connection.Close()
	s.mutx.Lock()
	s.connections = append(s.connections, connection)
	s.mutx.Unlock()
	ch := make(chan bool)
	go sendBoardToClient(connection, s)
	go readClientsDelta(connection, s, ch)
	<-ch
	connection.Close()
}

func (s *Server) broadCastBroad() {
	s.mutx.Lock()
	defer s.mutx.Unlock()
	boardByte, err := json.Marshal(s.board)
	if err != nil {
		fmt.Println("Error Mashaling Board", err)
	}
	boardByte = append(boardByte, '\n')
	for _, conn := range s.connections {
		_, err = conn.Write(boardByte)
		if err != nil {
			fmt.Println("Error sending Board to Client")
		}
	}
}

func readClientsDelta(conn net.Conn, s *Server, send chan bool) {
	defer func() {
		send <- true
	}()
	for {
		byt := make([]byte, 1024)
		size, err := conn.Read(byt)
		if err == nil {
			data := (strings.Split(string(byt[:size]), " "))
			delta := data[0]
			x, _ := strconv.ParseInt(data[1], 10, 64)
			y, _ := strconv.ParseInt(data[2], 10, 64)
			// if err1 != nil || err2 != nil || x < 0 || x >= 3 || y < 0 || y >= 3 {
			// 	fmt.Println("Invalid coordinates from client:", data)
			// 	continue
			// }
			s.mutx.Lock()
			s.board[x][y] = delta
			s.mutx.Unlock()
			fmt.Println(s.board)
			s.broadCastBroad()
		}
	}
}

func sendBoardToClient(conn net.Conn, s *Server) {
	for i := 0; i <= 1; i++ {
		boardByte, err := json.Marshal(s.board)
		if err != nil {
			fmt.Println("Error Mashaling Board", err)
			return
		}
		boardByte = append(boardByte, '\n')
		time.Sleep(time.Second)
		fmt.Println("Sending ", s.board)
		_, err = conn.Write(boardByte)
		if err != nil {
			fmt.Println("Error sending Board to Client")
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
		if len(s.connections) < 2 {
			fmt.Println("")
			fmt.Println("Connected", conn.LocalAddr().String())
			go handleConnection(conn, s)
		} else {
			fmt.Println("Unable to Accept Connection as Connection count is Over 2")
			conn.Write([]byte("404"))
			conn.Close()
		}
	}
}
