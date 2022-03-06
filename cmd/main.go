package main

import "github.com/mj-hagonoy/githubprofilessvc/pkg/server"

func main() {
	var server *server.Server = server.NewServer()
	server.Run()
}
