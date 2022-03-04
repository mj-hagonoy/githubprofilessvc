package githubprofilessvc

import "net/http"

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "OK"}`))
	}
}

func (s *Server) handleGithubUsersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "OK"}`))
	}
}
