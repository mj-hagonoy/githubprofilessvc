package githubprofilessvc

import (
	"fmt"
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	s := &Server{
		router: http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) Run() {
	err := http.ListenAndServe(":8080", s.router)
	fmt.Println(err)
}
