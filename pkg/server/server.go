package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mj-hagonoy/githubprofilessvc/pkg/config"
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

func (s *Server) Run() error {
	return http.ListenAndServe(config.GetConfig().Port, s.router)
}

func (s *Server) HttpError(w http.ResponseWriter, err error, httpCode int) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(httpCode)
	_, err = w.Write([]byte(fmt.Sprintf(`{"error": "%v"}`, err)))
	if err != nil {
		errors.Send(err)
	}
}

func (s *Server) JSON(w http.ResponseWriter, data interface{}, httpCode int) {
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

func (s *Server) NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
