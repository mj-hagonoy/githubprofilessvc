package main

import (
	"flag"

	"github.com/mj-hagonoy/githubprofilessvc/pkg/config"
	"github.com/mj-hagonoy/githubprofilessvc/pkg/server"
)

func main() {
	configFile := flag.String("config", "config.yaml", "configuration file")
	flag.Parse()
	if err := config.ParseConfig(*configFile); err != nil {
		panic(err)
	}
	var server *server.Server = server.NewServer()
	server.Run()
}
