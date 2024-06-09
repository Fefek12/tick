package server

type Server struct {
	addr string
	port int64
}

func (s *Server) newSerber(addr string, port int64) *Server {
	return &Server{
		addr: addr,
		port: port,
	}
}
