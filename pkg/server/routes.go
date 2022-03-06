package server

func (s *Server) routes() {
	s.router.HandleFunc("/", s.withExecutionLog(s.handleIndex()))
	s.router.HandleFunc("/api/v1/github/users", s.withExecutionLog(s.handleGithubUsersGet()))
}
