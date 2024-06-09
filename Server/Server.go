package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	addr string
	port int64
}

type Delta struct {
}

func (s *Server) newServer(addr string, port int64) *Server {
	return &Server{
		addr: addr,
		port: port,
	}
}

func getDeltaHandler(w http.ResponseWriter, r *http.Response) {

}

func (s *Server) Start() {
	http.ListenAndServe(s.addr, nil)
	fmt.Printf("Server Listening on %d", s.port)
}
