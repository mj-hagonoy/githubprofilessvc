package server

import (
	"encoding/json"
	"fmt"
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

func (s *Server) HttpError(w http.ResponseWriter, err error, httpCode int) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(httpCode)
	_, err = w.Write([]byte(fmt.Sprintf(`{"error": "%v"}`, err)))
	if err != nil {
		errors.Send(err)
	}
}

func (s *Server) Respond(w http.ResponseWriter, data interface{}, httpCode int) {
	bytesData, err := json.Marshal(data)
	if err != nil {
		s.HttpError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(httpCode)
	_, err = w.Write(bytesData)
	if err != nil {
		errors.Send(err)
	}
}
