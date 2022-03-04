package githubprofilessvc

func (s *Server) routes() {
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/api/v1/github/users", s.handleGithubUsersGet())
}
