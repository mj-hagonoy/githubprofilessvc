package githubprofilessvc

import (
	"net/http"

	"github.com/mj-hagonoy/githubprofilessvc/pkg/errors"
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
	go errors.Run()
	if err := http.ListenAndServe(":8080", s.router); err != nil {
		errors.Send(err)
		errors.Stop()
	}
}
