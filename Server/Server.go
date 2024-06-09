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

func (s *Server) NewServer(addr string, port int64) *Server {
	return &Server{
		addr: addr,
		port: port,
	}
}

func getDeltaHandler(w http.ResponseWriter, r *http.Response) {

}

func (s *Server) Start() {
	addr := fmt.Sprintf("%s:%d", s.addr, s.port)
	http.ListenAndServe(addr, nil)
	fmt.Printf("Server Listening on %s", addr)
}
