package server

import (
	"net/http"
	"strings"

	"github.com/mj-hagonoy/githubprofilessvc/pkg/errors"
	"github.com/mj-hagonoy/githubprofilessvc/pkg/users"
)

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.JSON(w, map[string]string{
			"message": "OK",
		}, http.StatusOK)
	}
}

func (s *Server) handleGithubUsersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srv := users.NewUsersService()
		result, err := srv.GetUsers(r.Context(), strings.Split(r.URL.Query().Get("usernames"), ",")...)
		if err != nil {
			errors.Send(err)
			s.HttpError(w, err, http.StatusBadRequest)
			return
		}

		s.JSON(w, result, http.StatusOK)
	}
}
