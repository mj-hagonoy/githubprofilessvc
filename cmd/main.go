package main

import "github.com/mj-hagonoy/githubprofilessvc/pkg/githubprofilessvc"

func main() {
	var server *githubprofilessvc.Server = githubprofilessvc.NewServer()
	server.Run()
}
