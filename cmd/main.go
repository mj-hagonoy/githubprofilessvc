package main

import (
	"flag"
	"os"

	"github.com/mj-hagonoy/githubprofilessvc/pkg/config"
	"github.com/mj-hagonoy/githubprofilessvc/pkg/server"
)

func main() {
	configFile := flag.String("config", "config.yaml", "configuration file")
	flag.Parse()
	must(config.ParseConfig(*configFile))

	var server *server.Server = server.NewServer()
	server.Run()
}

func must(err error) {
	if err != nil {
		os.Exit(1)
	}
}
