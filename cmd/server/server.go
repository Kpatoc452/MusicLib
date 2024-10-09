package server

import (
	"net/http"
)

type Server struct {
	srv *http.Server
	Rgstr map[string]http.HandlerFunc
}

func NewServer() *Server{
	return &Server{
		srv: nil,
		Rgstr: make(map[string]http.HandlerFunc),
	}
}

func(s *Server) Init(cfg *Config){
	mux := s.Register(s.Rgstr)

	s.srv = &http.Server{
		Addr: cfg.Addr,
		Handler: mux,
	}
}

func(s *Server) Run(){
	s.srv.ListenAndServe()
}

func(s *Server) AddEndpoint(endpoint string, handler http.HandlerFunc){
	s.Rgstr[endpoint] = handler
}

func(s *Server) Register(endpoints map[string]http.HandlerFunc) (*http.ServeMux){
	mux := http.ServeMux{}
	for endpoint, handler := range endpoints{
		mux.HandleFunc(endpoint, handler)
	}
	return &mux
}