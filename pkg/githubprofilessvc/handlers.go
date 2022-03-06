package githubprofilessvc

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "OK"}`))
	}
}

func (s *Server) handleGithubUsersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srv := GithubUsersService{}
		result, err := srv.GetUsers(r.Context(), strings.Split(r.URL.Query().Get("usernames"), ",")...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		bytesData, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytesData)
	}
}
