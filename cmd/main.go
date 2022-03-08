package main

import (
	"flag"
	"log"
	"os"

	"github.com/mj-hagonoy/githubprofilessvc/pkg/config"
	"github.com/mj-hagonoy/githubprofilessvc/pkg/errors"
	"github.com/mj-hagonoy/githubprofilessvc/pkg/server"
)

func main() {
	errors.Run()
	configFile := flag.String("config", "config.yaml", "configuration file")
	flag.Parse()
	must(config.ParseConfig(*configFile))

	var server *server.Server = server.NewServer()
	must(server.Run())
}

func must(err error) {
	if err != nil {
		log.Println(err)
		errors.Stop()
		os.Exit(1)
	}
}
