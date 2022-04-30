package http

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	port   int
	server *http.Server
}

func NewServer(port int) *Server {
	server := &Server{
		port: port,
	}

	return server
}

func (s *Server) Create() {

	router := InitRouter()

	address := fmt.Sprintf("0.0.0.0:%d", s.port)

	server := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,
		Handler:      router,
	}

	s.server = server
}

func (s *Server) Start() {
	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
